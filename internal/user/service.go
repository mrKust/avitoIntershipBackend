package user

import (
	"avitoIntershipBackend/internal/masterBalance"
	"avitoIntershipBackend/internal/service"
	"avitoIntershipBackend/internal/transaction"
	"avitoIntershipBackend/pkg/logging"
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
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
}

func (b bisLogic) Billing(ctx context.Context, user *User) error {
	var userInDB User
	var err error
	userInDB, err = b.repositoryUser.FindOne(ctx, fmt.Sprintf("%d", user.ID))
	if err != nil {
		if err.Error() != "rows not found" {
			return err
		}
		newUser := User{Balance: user.Balance}
		b.repositoryUser.Create(ctx, &newUser)
		return nil
	}

	if (strings.Contains(user.Balance, "-") == true) || (strings.Contains(user.Balance, "+") == true) {
		err = fmt.Errorf("incorrect \"balanace\" parametr in request")
		b.logger.Debugf("can't add money with billing due to error: %v", err)
		return err
	}

	tmpVal1, convertErr := strconv.ParseFloat(user.Balance, 64)
	if convertErr != nil {
		b.logger.Debugf("can't convert input balance data to float error: %v", convertErr)
		return convertErr
	}

	tmpVal2, _ := strconv.ParseFloat(userInDB.Balance, 64)

	userInDB.Balance = fmt.Sprintf("%f", tmpVal1+tmpVal2)
	user.Balance = userInDB.Balance
	err = b.repositoryUser.Update(ctx, userInDB)
	if err != nil {
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
		b.logger.Debugf("can't write action in transactions list")
		return err
	}

	return nil
}

func (b bisLogic) ReserveMoney(ctx context.Context, balance *masterBalance.MasterBalance) error {

	var user User
	var err error

	if (strings.Contains(balance.MoneyAmount, "-") == true) || (strings.Contains(balance.MoneyAmount, "+") == true) {
		err = fmt.Errorf("incorrect \"balanace\" parametr in request")
		b.logger.Debugf("can't reserve money for service due to error: %v", err)
		return err
	}

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
	err = b.repositoryUser.Update(ctx, user)
	if err != nil {
		b.logger.Debugf("can't update user in db")
		return err
	}

	err = b.repositoryMasterBalance.Create(ctx, balance)
	if err != nil {
		b.logger.Debugf("can't save masterbal in db")
		return err
	}

	serviceTmp, err := b.repositoryService.FindOne(ctx, balance.ServiceId)
	if err != nil {
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
		b.logger.Debugf("can't write action in transactions list")
		return err
	}

	return nil
}

func (b bisLogic) AcceptMoney(ctx context.Context, balance *masterBalance.MasterBalance) error {
	var err error

	if (strings.Contains(balance.MoneyAmount, "-") == true) || (strings.Contains(balance.MoneyAmount, "+") == true) {
		err = fmt.Errorf("incorrect \"balanace\" parametr in request")
		b.logger.Debugf("can't accept money for service due to error: %v", err)
		return err
	}

	err = b.repositoryMasterBalance.FindOneByParam(ctx, balance)

	err = b.repositoryMasterBalance.Delete(ctx, fmt.Sprintf("%d", balance.ID))
	if err != nil {
		b.logger.Debugf("can't delete masterbal in db")
		return err
	}

	serviceTmp, err := b.repositoryService.FindOne(ctx, balance.ServiceId)
	if err != nil {
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
		b.logger.Debugf("can't write action in transactions list")
		return err
	}

	return nil
}

func (b bisLogic) RejectMoney(ctx context.Context, balance *masterBalance.MasterBalance) error {

	var user User
	var err error

	if (strings.Contains(balance.MoneyAmount, "-") == true) || (strings.Contains(balance.MoneyAmount, "+") == true) {
		err = fmt.Errorf("incorrect \"balanace\" parametr in request")
		b.logger.Debugf("can't reject money request for service due to error: %v", err)
		return err
	}

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
	err = b.repositoryUser.Update(ctx, user)
	if err != nil {
		b.logger.Debugf("can't update user in db")
		return err
	}

	err = b.repositoryMasterBalance.Delete(ctx, fmt.Sprintf("%d", balance.ID))
	if err != nil {
		b.logger.Debugf("can't delete masterbal from db")
		return err
	}

	serviceTmp, err := b.repositoryService.FindOne(ctx, balance.ServiceId)
	if err != nil {
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
		b.logger.Debugf("can't write action in transactions list")
		return err
	}

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

	transactionsList, _ = b.repositoryTransaction.FindAllForPeriod(ctx, month, year)

	resultMap := make(map[string]float64)

	for _, transactionTmp := range transactionsList {
		if strings.Contains(transactionTmp.ForService, "billing") {
			continue
		}
		if val, ok := resultMap[transactionTmp.ForService]; ok {
			moneyOne := val
			moneyTwo, err := strconv.ParseFloat(transactionTmp.MoneyAmount, 64)
			if err != nil {

			}
			resultMap[transactionTmp.ForService] = moneyOne + moneyTwo
		} else {
			resultMap[transactionTmp.ForService], err = strconv.ParseFloat(transactionTmp.MoneyAmount, 64)
			if err != nil {

			}
		}
	}

	reportName := strings.ReplaceAll("./reports/report"+time.Now().Format(time.RFC822)+".csv", ":", "-")
	file, errName := os.Create(reportName)
	if err != nil {
		fmt.Println(errName)
	}

	w := csv.NewWriter(file)
	w.Comma = ';'
	for k, v := range resultMap {
		row := []string{k, fmt.Sprintf("%f", v)}
		w.Write(row)
	}
	w.Flush()

	return "", nil
}

func (b bisLogic) GetUserTransactions(ctx context.Context, id string, pageNum string, sortSum string, sortDate string) ([]string, error) {

	var err error
	transactionsList := make([]transaction.Transaction, 0)

	transactionsList, _ = b.repositoryTransaction.FindPageForUser(ctx, id, pageNum, sortSum, sortDate)
	fmt.Println(transactionsList)

	return nil, err
}

func NewService(repositoryUser Repository, repositoryMasterBalance masterBalance.Repository,
	repositoryTransaction transaction.Repository, repositoryService service.Repository, logger *logging.Logger) BusinessLogic {
	return &bisLogic{
		repositoryUser:          repositoryUser,
		repositoryMasterBalance: repositoryMasterBalance,
		repositoryTransaction:   repositoryTransaction,
		repositoryService:       repositoryService,
		logger:                  logger,
	}
}
