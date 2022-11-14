package user

import (
	"avitoIntershipBackend/internal/config"
	"avitoIntershipBackend/internal/masterBalance"
	"avitoIntershipBackend/internal/service"
	"avitoIntershipBackend/internal/transaction"
	"avitoIntershipBackend/pkg/client/postgresql"
	"avitoIntershipBackend/pkg/logging"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type BusinessLogic interface {
	Billing(ctx context.Context, user *User) error
	ReserveMoney(ctx context.Context, balance *masterBalance.MasterBalance) error
	AcceptMoney(ctx context.Context, balance *masterBalance.MasterBalance) error
	RejectMoney(ctx context.Context, balance *masterBalance.MasterBalance) error
	Report(ctx context.Context, month string, year string) (string, error)
	GetBalance(ctx context.Context, id string) (User, error)
	GetUserTransactions(ctx context.Context, id string, pageNum string, sortSum string, sortDate string) ([]string, error)
}

type bisLogic struct {
	repositoryUser          Repository
	repositoryMasterBalance masterBalance.Repository
	repositoryTransaction   transaction.Repository
	repositoryService       service.Repository
	logger                  *logging.Logger
	conn                    postgresql.Client
	mutex                   *sync.Mutex
}

func (b bisLogic) Billing(ctx context.Context, user *User) error {
	var userInDB User
	var err error

	if (strings.Contains(user.Balance, "-") == true) || (strings.Contains(user.Balance, "+") == true) {
		err = fmt.Errorf("incorrect balanace parametr in request")
		b.logger.Debugf("can't add money with billing due to error: %v", err)
		return err
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.logger.Debug("read val 1")
	userInDB, err = b.repositoryUser.FindOne(ctx, fmt.Sprintf("%d", user.ID))
	if err != nil {
		newUser := User{Balance: user.Balance}
		b.repositoryUser.Create(ctx, &newUser)
		user.ID = newUser.ID
		return nil
	}

	tmpVal1, convertErr := strconv.ParseFloat(user.Balance, 64)
	if convertErr != nil {
		b.logger.Debugf("can't convert input balance data to float error: %v", convertErr)
		return convertErr
	}

	tmpVal2, _ := strconv.ParseFloat(userInDB.Balance, 64)

	b.logger.Debug("count balance 2")
	userInDB.Balance = fmt.Sprintf("%f", tmpVal1+tmpVal2)
	user.Balance = userInDB.Balance

	tx, errTx := b.conn.BeginTx(context.Background(), pgx.TxOptions{})
	if errTx != nil {
		b.logger.Debugf("can't strat transaction")
		return errTx
	}

	b.logger.Debug("update value in db 3")
	err = b.repositoryUser.Update(ctx, userInDB)

	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	newTransaction := transaction.Transaction{
		FromId:      "0",
		ToId:        fmt.Sprintf("%d", userInDB.ID),
		ForService:  "billing",
		OrderId:     "-",
		MoneyAmount: user.Balance,
		Status:      "complete",
	}

	err = b.repositoryTransaction.Create(ctx, &newTransaction)
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't write action in transactions list")
		return err
	}
	tx.Commit(context.Background())

	return nil
}

func (b bisLogic) ReserveMoney(ctx context.Context, balance *masterBalance.MasterBalance) error {

	var user User
	var err error

	if (strings.Contains(balance.MoneyAmount, "-") == true) || (strings.Contains(balance.MoneyAmount, "+") == true) {
		err = fmt.Errorf("incorrect balanace parametr in request")
		b.logger.Debugf("can't reserve money for service due to error: %v", err)
		return err
	}

	b.mutex.Lock()
	user, err = b.repositoryUser.FindOne(ctx, balance.FromId)
	if err != nil {
		b.logger.Debugf("can't get user from db")
		return err
	}

	tmpVal1, convertErr := strconv.ParseFloat(balance.MoneyAmount, 64)
	if convertErr != nil {
		b.logger.Debugf("can't convert input balance data to float error: %v", convertErr)
		return convertErr
	}

	tmpVal2, _ := strconv.ParseFloat(user.Balance, 64)

	balanceAfterFreeze := tmpVal2 - tmpVal1

	if balanceAfterFreeze < 0 {
		b.logger.Trace("user don't have enough money to pay for service")
		return fmt.Errorf("lack of money for service is %f", balanceAfterFreeze)
	}

	user.Balance = fmt.Sprintf("%f", balanceAfterFreeze)
	tx, errTx := b.conn.BeginTx(context.Background(), pgx.TxOptions{})
	if errTx != nil {
		b.logger.Debugf("can't strat transaction")
		return errTx
	}

	err = b.repositoryUser.Update(ctx, user)
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't update user in db")
		return err
	}

	err = b.repositoryMasterBalance.Create(ctx, balance)
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't save masterbal in db")
		return err
	}

	serviceTmp, err := b.repositoryService.FindOne(ctx, balance.ServiceId)
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't get service from db")
		return err
	}

	newTransaction := transaction.Transaction{
		FromId:      balance.FromId,
		ToId:        "0",
		ForService:  serviceTmp.Name,
		OrderId:     balance.OrderId,
		MoneyAmount: balance.MoneyAmount,
		Status:      "freeze",
	}

	err = b.repositoryTransaction.Create(ctx, &newTransaction)
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't write action in transactions list")
		return err
	}
	tx.Commit(context.Background())
	b.mutex.Unlock()

	return nil
}

func (b bisLogic) AcceptMoney(ctx context.Context, balance *masterBalance.MasterBalance) error {
	var err error

	if (strings.Contains(balance.MoneyAmount, "-") == true) || (strings.Contains(balance.MoneyAmount, "+") == true) {
		err = fmt.Errorf("incorrect balanace parametr in request")
		b.logger.Debugf("can't accept money for service due to error: %v", err)
		return err
	}

	b.mutex.Lock()
	err = b.repositoryMasterBalance.FindOneByParam(ctx, balance)

	if err != nil {
		b.logger.Debugf("can't find request in db")
		return err
	}

	tx, errTx := b.conn.BeginTx(context.Background(), pgx.TxOptions{})
	if errTx != nil {
		b.logger.Debugf("can't strat transaction")
		return errTx
	}

	err = b.repositoryMasterBalance.Delete(ctx, fmt.Sprintf("%d", balance.ID))
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't delete masterbal in db")
		return err
	}

	serviceTmp, err := b.repositoryService.FindOne(ctx, balance.ServiceId)
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't get service from db")
		return err
	}

	newTransaction := transaction.Transaction{
		FromId:      balance.FromId,
		ToId:        "0",
		ForService:  serviceTmp.Name,
		OrderId:     balance.OrderId,
		MoneyAmount: balance.MoneyAmount,
		Status:      "complete",
	}

	err = b.repositoryTransaction.Create(ctx, &newTransaction)
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't write action in transactions list")
		return err
	}
	tx.Commit(context.Background())
	b.mutex.Unlock()

	return nil
}

func (b bisLogic) RejectMoney(ctx context.Context, balance *masterBalance.MasterBalance) error {

	var user User
	var err error

	if (strings.Contains(balance.MoneyAmount, "-") == true) || (strings.Contains(balance.MoneyAmount, "+") == true) {
		err = fmt.Errorf("incorrect balanace parametr in request")
		b.logger.Debugf("can't reject money request for service due to error: %v", err)
		return err
	}

	b.mutex.Lock()
	err = b.repositoryMasterBalance.FindOneByParam(ctx, balance)
	if err != nil {
		b.logger.Debugf("can't get masterBal from db")
		return err
	}

	user, err = b.repositoryUser.FindOne(ctx, balance.FromId)
	if err != nil {
		b.logger.Debugf("can't get user from db")
		return err
	}

	tmpVal1, _ := strconv.ParseFloat(balance.MoneyAmount, 64)
	tmpVal2, _ := strconv.ParseFloat(user.Balance, 64)

	balanceAfterReject := tmpVal2 + tmpVal1

	user.Balance = fmt.Sprintf("%f", balanceAfterReject)

	tx, errTx := b.conn.BeginTx(context.Background(), pgx.TxOptions{})
	if errTx != nil {
		b.logger.Debugf("can't strat transaction")
		return errTx
	}

	err = b.repositoryUser.Update(ctx, user)
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't update user in db")
		return err
	}

	err = b.repositoryMasterBalance.Delete(ctx, fmt.Sprintf("%d", balance.ID))
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't delete masterbal from db")
		return err
	}

	serviceTmp, err := b.repositoryService.FindOne(ctx, balance.ServiceId)
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't get service from db")
		return err
	}

	newTransaction := transaction.Transaction{
		FromId:      balance.FromId,
		ToId:        "0",
		ForService:  serviceTmp.Name,
		OrderId:     balance.OrderId,
		MoneyAmount: balance.MoneyAmount,
		Status:      "canceled",
	}

	err = b.repositoryTransaction.Create(ctx, &newTransaction)
	if err != nil {
		tx.Rollback(context.Background())
		b.logger.Debugf("can't write action in transactions list")
		return err
	}
	tx.Commit(context.Background())
	b.mutex.Unlock()

	return nil
}

func (b bisLogic) GetBalance(ctx context.Context, id string) (User, error) {
	var user User
	var err error

	user, err = b.repositoryUser.FindOne(ctx, id)
	if err != nil {
		err = fmt.Errorf("no users with such id %s", id)
		return User{}, err
	}

	return user, nil
}

func (b bisLogic) Report(ctx context.Context, month string, year string) (string, error) {

	var err error
	transactionsList := make([]transaction.Transaction, 0)

	transactionsList, err = b.repositoryTransaction.FindAllForPeriod(ctx, month, year)
	if err != nil {
		err = fmt.Errorf("can't get transactions")
		return "", err
	}

	if len(transactionsList) == 0 {
		b.logger.Debugf("transactions for choosen period didn't found")
		err = fmt.Errorf("no transactions for this period")
		return "", err
	}

	resultMap := make(map[string]float64)

	for _, transactionTmp := range transactionsList {
		if strings.Contains(transactionTmp.ForService, "billing") {
			continue
		}
		if val, ok := resultMap[transactionTmp.ForService]; ok {
			moneyOne := val
			moneyTwo, err := strconv.ParseFloat(transactionTmp.MoneyAmount, 64)
			if err != nil {
				b.logger.Debugf("can't convert money to float due to error: %v", err)
				return "", err
			}
			resultMap[transactionTmp.ForService] = moneyOne + moneyTwo
		} else {
			resultMap[transactionTmp.ForService], err = strconv.ParseFloat(transactionTmp.MoneyAmount, 64)
			if err != nil {
				b.logger.Debugf("can't convert money to float due to error: %v", err)
				return "", err
			}
		}
	}

	reportName := strings.ReplaceAll("./reports/report"+time.Now().Format(time.RFC822)+".csv", ":", "-")
	file, errName := os.Create(reportName)
	if err != nil {
		b.logger.Debugf("can't create file due to error: %v", err)
		fmt.Println(errName)
	}

	w := csv.NewWriter(file)
	w.Comma = ';'
	for k, v := range resultMap {
		row := []string{k, fmt.Sprintf("%f", v)}
		err = w.Write(row)
		if err != nil {
			b.logger.Debugf("can't write data to file due to error %v", err)
			os.Remove(reportName)
			return "", err
		}
	}
	w.Flush()
	err = file.Close()
	if err != nil {
		b.logger.Debugf("can't close report's file")
		return "", err
	}
	resultString := fmt.Sprintf("%s:%s%s", config.GetConfig().Listen.BindIP, config.GetConfig().Listen.Port, reportName[1:])

	return resultString, nil
}

func (b bisLogic) GetUserTransactions(ctx context.Context, id string, pageNum string, sortSum string, sortDate string) ([]string, error) {

	var err error
	transactionsList := make([]transaction.Transaction, 0)

	if !strings.Contains(sortSum, "desc") && !strings.Contains(sortSum, "asc") {
		b.logger.Debugf("incorect parametr sortSum")
		err = fmt.Errorf("incorrect input parametrs")
		return nil, err
	}

	if !strings.Contains(sortDate, "desc") && !strings.Contains(sortDate, "asc") {
		b.logger.Debugf("incorect parametr sortDate")
		err = fmt.Errorf("incorrect input parametrs")
		return nil, err
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		b.logger.Debugf("incorect parametr id")
		err = fmt.Errorf("incorrect input parametrs")
		return nil, err
	}
	pageNumInt, err := strconv.ParseInt(pageNum, 10, 64)
	if err != nil {
		b.logger.Debugf("incorect parametr pageNum")
		err = fmt.Errorf("incorrect input parametrs")
		return nil, err
	}

	if (idInt < 0) || (pageNumInt <= 0) {
		b.logger.Debugf("id and pageNum must be positive (>0)")
		err = fmt.Errorf("parametrs equal 0")
		return nil, err
	}

	transactionsList, err = b.repositoryTransaction.FindPageForUser(ctx, id, pageNum, sortSum, sortDate)
	if err != nil {
		b.logger.Debugf("can't get page of transactions of user due to error %v", err)
		return nil, err
	}
	if len(transactionsList) == 0 {
		return nil, fmt.Errorf("no transactions for user")
	}

	var resultString []string
	resultString = make([]string, 0)
	for _, transactionTmp := range transactionsList {
		date := fmt.Sprint(transactionTmp.Date)
		resultString = append(resultString, fmt.Sprintf("user pay %s for %s at %s", transactionTmp.MoneyAmount, transactionTmp.ForService, date))
	}

	return resultString, err
}

func NewService(repositoryUser Repository, repositoryMasterBalance masterBalance.Repository,
	repositoryTransaction transaction.Repository, repositoryService service.Repository,
	logger *logging.Logger, conn postgresql.Client, mutex *sync.Mutex) BusinessLogic {
	return &bisLogic{
		repositoryUser:          repositoryUser,
		repositoryMasterBalance: repositoryMasterBalance,
		repositoryTransaction:   repositoryTransaction,
		repositoryService:       repositoryService,
		logger:                  logger,
		conn:                    conn,
		mutex:                   mutex,
	}
}
