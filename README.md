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

Link to OpenApi specification
```bash
http://localhost:8080/swagger/index.html/
```

## Request/response examples

The users, transactions, masterBalance tables are initially empty. The service table looks like this:<br><br>
![serviceInitialTable](https://user-images.githubusercontent.com/45081619/201553901-f49fd332-1453-412d-aa99-747e4866d2d0.png)<br> *Pic 1. Table service.* <br><br>

### The method of billing to the balance has the following url: http://localhost:8080/billing <br><br>
#### Request:<br><br>
![billingRequest](https://user-images.githubusercontent.com/45081619/201554334-0d65e4d9-9dfc-4f66-89c5-42888cef068d.png)
<br> *Pic 2. Billing request.* <br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201555690-41b8103e-813d-4d5a-87af-62dc5279781c.png)
<br> *Pic 3. Billing succesful response.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201653801-4e0d56a0-4dbc-45e6-8d0e-de6238b72de4.png)
<br> *Pic 3.1. Billing succesful response with creatig new user.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201632684-5acca833-e7e8-4c81-ab97-d98b013eaea7.png)
<br> *Pic 4. Billing response with error.* <br><br>

### The method of reserving money from the main balance in a master account has the following url: http://localhost:8080/moneyFreeze
#### Request:<br><br>
![moneyFreezeRequest](https://user-images.githubusercontent.com/45081619/201554458-4eead1d4-c7cf-42fa-9c6c-c80385956e0f.png)
<br>*Pic 5. Reserving request.* <br><br>
<br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201633193-d3b74f34-8f6d-4535-a2fc-3e5c3116540e.png)
*Pic 5. Reserving successfully response.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201633756-7c67adcd-c7e8-4498-a70e-d7022f6fda16.png)
*Pic 6. Reserving badly response.* <br><br>

### Accept money method has the following url: http://localhost:8080/moneyAccept <br><br>
#### Request:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201554827-16a809b2-d250-42b9-800e-c55c36cf1633.png)
<br> *Pic 7. Recognition request.* <br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201634248-e115141e-0ced-4851-9928-1eb7f851855f.png)
*Pic 8. Reserving successfully response.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201634414-b044016e-5261-4777-a6b7-f97c3973fd26.png)
*Pic 9. Reserving badly response.* <br><br>

### Reject money method has the following url: http://localhost:8080/moneyReject <br><br>
#### Request:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201636135-8d140755-a57b-4fd2-9dd1-e898a419f32a.png)
<br> *Pic 10. Reject request.* <br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201636917-bc62ba5d-e381-4cec-86f8-a4c2ef1fb5e8.png)
*Pic 11. Reserving successfully response.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201636774-c413f104-0108-4971-895f-fff3da5b8746.png)
*Pic 12. Reserving badly response.* <br><br>

### User balance receipt method has the following url: http://localhost:8080/users/:id <br><br>
#### Request:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201554944-2cc88558-0835-438c-b429-ca1c3e3c6ad0.png)
<br> *Pic 13. Get balance request.* <br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201554981-a82dcb3c-a1eb-4b2a-bae2-e18cb9e01c8f.png)
<br> *Pic 14. Get balance successfully response.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201634564-35a5ee8b-e79f-4e41-a1ea-8e0537b57e11.png)
<br> *Pic 15. Reserving badly response.* <br><br>

### Method to get monthly report has the following url: http://localhost:8080/report/:month/:year <br><br>
#### Request:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201555091-db058c0b-80a6-4d86-9106-5f50ef9b99e4.png)
<br> *Pic 16. Get report request.* <br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201555107-d1eb8c58-2a9e-4e3c-91db-f4b9c52b2b76.png)
<br> *Pic 17. Get report successfully response.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201634889-abf329cf-2ab0-462f-9644-82a9e95fccfe.png)
<br> *Pic 18. Get report badly response 1.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201635037-47da9163-45c1-4481-bb27-87495f19fb8d.png)
<br> *Pic 19. Get report badly response 2.* <br><br>
#### Example report
![изображение](https://user-images.githubusercontent.com/45081619/201555375-f430e321-7a01-4339-8c75-47c815053747.png)<br>
*Pic 20. Report example.* <br><br>

### Method to get user transactions has the following url: http://localhost:8080/transactions/:userid/:pageNum/sortSum/sortDate <br><br>
#### Request:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201555193-dbc1a997-4461-4c4c-96ab-11525d4cb407.png)
<br> *Pic 21. Get page request.* <br><br>
#### Response:<br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201635380-df8dca32-d960-4d4b-95ce-b362376a8339.png)
<br> *Pic 22. Get page successfully response.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201635380-df8dca32-d960-4d4b-95ce-b362376a8339.png)
<br> *Pic 23. Get page badly response.* <br><br>
![изображение](https://user-images.githubusercontent.com/45081619/201635794-8f743a21-6a7c-4e80-b9d5-6f601ee7a1a9.png)
<br> *Pic 24. Get page badly response.* <br><br>
