# avitoIntershipBackend

## Description
Микросервис для работы с балансом пользователей

## Installation
First step clone the git repository
```bash
git clone https://github.com/mrKust/avitoIntershipBackend.git
```

After that go to the folder avitoIntershipBackend
```bash
cd avitoIntershipBackend
```

After that run the following command
```bash
docker-compose up 
```

After the containers are deployed, the microservice is ready.

## Request/response examples

The users, transactions, masterBalance tables are initially empty. The service table looks like this:<br><br>
![serviceInitialTable](https://user-images.githubusercontent.com/45081619/201553901-f49fd332-1453-412d-aa99-747e4866d2d0.png)<br> *Pic 1. Table service.* <br><br>

### The method of billing to the balance has the following url: http://localhost:8080/billing <br><br>
#### Request:<br><br>
![billingRequest](https://user-images.githubusercontent.com/45081619/201554334-0d65e4d9-9dfc-4f66-89c5-42888cef068d.png)
<br> *Pic 2. Billing request.* <br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201555690-41b8103e-813d-4d5a-87af-62dc5279781c.png)<br> *Pic 3. Billing response.* <br><br>

### The method of reserving money from the main balance in a master account has the following url: http://localhost:8080/moneyFreeze
#### Request:<br><br>
![moneyFreezeRequest](https://user-images.githubusercontent.com/45081619/201554458-4eead1d4-c7cf-42fa-9c6c-c80385956e0f.png)
<br>*Pic 4. Reserving request.* <br><br>
<br><br>
#### Response:<br><br>
![moneyFreezeResponse](https://user-images.githubusercontent.com/45081619/201554581-851c594f-2eac-416a-915a-b6f1997ef832.png)
*Pic 5. Reserving successfully response.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201554510-e031b614-54e1-42f3-a12b-0484279cf20d.png)
*Pic 6. Reserving badly response.* <br><br>

### Accept money method has the following url: http://localhost:8080/moneyAccept <br><br>
#### Request:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201554827-16a809b2-d250-42b9-800e-c55c36cf1633.png)
<br> *Pic 7. Recognition request.* <br><br>

### User balance receipt method has the following url: http://localhost:8080/users/:id <br><br>
#### Request:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201554944-2cc88558-0835-438c-b429-ca1c3e3c6ad0.png)
<br> *Pic 8. Get balance request.* <br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201554981-a82dcb3c-a1eb-4b2a-bae2-e18cb9e01c8f.png)
<br> *Pic 9. Get balance successfully response.* <br><br>
![image](https://user-images.githubusercontent.com/79422421/197415575-118cc5be-680b-49b6-8204-43d6d9166dab.png)<br> *Pic 10. Reserving badly response.* <br><br>

### Method to get monthly report has the following url: http://localhost:8080/report/:month/:year <br><br>
#### Request:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201555091-db058c0b-80a6-4d86-9106-5f50ef9b99e4.png)<br> *Pic 11. Get report request.* <br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201555107-d1eb8c58-2a9e-4e3c-91db-f4b9c52b2b76.png)
<br> *Pic 12. Get report successfully response.* <br><br>
![image](https://user-images.githubusercontent.com/79422421/197415695-74dfdb5e-d0eb-4c9d-8dbb-ee7999ab8c30.png)<br> *Pic 13. Get report badly response 1.* <br><br>
#### Example report
![изображение](https://user-images.githubusercontent.com/45081619/201555375-f430e321-7a01-4339-8c75-47c815053747.png)<br>
*Pic 14. Report example.* <br><br>

### Method to get user transactions has the following url: http://localhost:8080/transactions/:userid/:pageNum/sortSum/sortDate <br><br>
#### Request:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201555193-dbc1a997-4461-4c4c-96ab-11525d4cb407.png)<br> *Pic 15. Get page request.* <br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201555341-53a6f594-b8a0-40b1-815c-774aab9faa61.png)
<br> *Pic 16. Get page successfully response.* <br><br>
