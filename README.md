# KBank-API
KBank-API Project using the Kasikorn Bank's bill payment API including Inquiry & Payment, to act as a middleman to enable payments between clients and stores. Use 
Inquiry for check bill details before making a payment and Payment for confirm payment.

### Table of contents

- [ How To Run ](#how-to-run)
- [ ER Diagram ](#er-diagram)
- [ API Service ](#api-service)
    - [ Routes ](#routes)
    - [ KBank API (bill-payment-service) ](#kbank-api-bill-payment-service)
    - [ Backend API (database) ](#backend-api-database)
- [ Conference ](#conference)

<a name="how-to-run"></a>
## How To Run 

1. Clone the repository:
```git
git clone https://github.com/Teerapat-Kuiyawattananon/kbank-api.git
```

2. Install PostgreSQL with Docker:
```bash
docker compose up -d
```

3. Run Program:
```bash
go run main.go 
```
or 
```
air 
```

4. Call Swagger API:
```
curl localhost:8080/swagger/doc.json
```

5. Open SwaggerUI in browser by using URL:
```
http://localhost:8080/swagger/index.html
```

<a name="er-diagram"></a>
## ER Diagram
Database ER diagram for this project
- [ER_Diagram](https://dbdocs.io/nattapon.mee/Bill-Payment?view=relationships)

<a name="api-service"></a>
## API Service

<a name="routes"></a>
### Routes
#### Bill-Payment
- POST `http://localhost:8080/api/billpayment/lookup`
- POST `http://localhost:8080/api/billpayment/payment`
#### Billers
- GET  `http://localhost:8080/api/billers`
- POST `http://localhost:8080/api/billers`
#### Customers
- GET `http://localhost:8080/api/customers`
- GET `http://localhost:8080/api/customers/:id`
- POST `http://localhost:8080/api/customers`
#### Bills
- GET `http://localhost:8080/api/bills`
- GET `http://localhost:8080/api/bills/search?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`
- PUT `http://localhost:8080/api/bills/:id`
- POST `http://localhost:8080/api/bills`

<a name="kbank-api"></a>
### KBank API (bill-payment-service)
#### Inquiry Lookup
- METHOD POST `http://localhost:8080/api/billpayment/lookup`
- BODY
```json
{
  "functionName": "BillLookup",
  "transactionId": "98099310720200530183382100",
  "transactionDateTime": "2023-08-15T21:39:46+07:00",
  "billerType": "",
  "billerId": "98499",
  "terminalNo": "xxxxxxx8888",
  "channelCode": "MOB",
  "tranAmount": "120.00",
  "senderBankCode": "Kbank",
  "reference1": "300000025751",
  "reference2": "1733084",
  "language": "EN",
  "apiKey": "SecureKey123"
}
```
#### Response
- status : `200 OK`
- BODY
```json
{
    "functionName": "BillLookupResponse",
    "transactionId": "98099310720200530183382100",
    "transactionDateTime": "2023-08-17T21:17:04+07:00",
    "billerTransactionId": "98099310720200530183382100",
    "responseCode": "0000",
    "responseDescription": "Success",
    "billerType": "",
    "billerId": "98499",
    "terminalNo": "xxxxxxx8888",
    "promptPayTransactionId": "",
    "typeofReceiver": "",
    "reference1": "300000025751",
    "reference2": "1733084",
    "reference3": "",
    "tranAmount": "",
    "info1": "",
    "info2": "",
    "info3": "",
    "additional": {
        "toBillerAccountName": "",
        "toBillerServiceName": "",
        "payerFee": "",
        "settlementDate": "",
        "receiverTaxID": "",
        "dueDate": "",
        "rtpReference": "",
        "rsAppId": ""
    }
}
```
#### Payment Confirm
- METHOD POST `http://localhost:8080/api/billpayment/payment`
- BODY
```json
{
  "functionName": "BillPayment",
  "transactionId": "98099310720200530183382100",
  "transactionDateTime": "2023-08-06T10:33:18.450+07:00",
  "billerType": "",
  "billerId": "98499",
  "terminalNo": "xxxxxxx8888",
  "channelCode": "MOB",
  "tranAmount": "120.00",
  "senderBankCode": "Kbank",
  "isRetry": "0",
  "reference1": "300000025751",
  "reference2": "1733084",
  "language": "EN",
  "apiKey": "SecureKey123"
}
```
#### Response
- status : `200 OK`
- BODY
```json
{
    "functionName": "BillPaymentResponse",
    "transactionId": "98099310720200530183382100",
    "responseDateTime": "2023-08-17T21:09:08+07:00",
    "billerTransactionId": "98099310720200530183382100",
    "responseCode": "0000",
    "responseDescription": "Success",
    "terminalNo": "xxxxxxx8888",
    "promptPayTransactionId": "",
    "additional": {
        "settlementDate": "",
        "rsAppId": ""
    }
}
```
*** After you payment confirm that bill will update status to already_paid

<a name="backend-api-database"></a>
### Backend API (Database)
#### Biller_account
- METHOD GET `http://localhost:8080/api/billers`
#### Response
```json
[
    {
        "id": 98499,
        "name": "ThaiTee",
        "service_name": "Abank"
    }
]
```
- METHOD POST `http://localhost:8080/api/billers`
- BODY
```json
{
    "name" : "Yum",
    "service_name" : "SBB"
}
```
##### Response
```json
{
    "id": 2,
    "name": "Yum",
    "service_name": "SBB"
}
```
#### Customer
- METHOD GET `http://localhost:8080/api/customers`
#### Response
```json
[
    {
        "id": 300000025751,
        "firstName": "Teerapat",
        "lastName": "Ku",
        "titleName": "",
        "mobileNumber": "0888888888",
        "createdAt": "2023-08-16T15:06:05.652791Z"
    },
    {
        "id": 1,
        "firstName": "Testers",
        "lastName": "Tu",
        "titleName": "",
        "mobileNumber": "0999999999",
        "createdAt": "2023-08-16T22:23:28.643741Z"
    }
]
```
- METHOD GET `http://localhost:8080/api/customers/:id`
#### Response
```json
{
    "id": 1,
    "firstName": "Testers",
    "lastName": "Tu",
    "titleName": "",
    "mobileNumber": "0999999999",
    "createdAt": "2023-08-16T22:23:28.643741Z"
}
```
- METHOD POST `http://localhost:8080/api/customers`
- BODY
```json
{
    "firstName": "Teera",
    "lastName" : "Pl",
    "titleName" : "",
    "mobileNumber" : "0999999999"
}
```
#### Response
```json
{
    "id": 2,
    "firstName": "Teera",
    "lastName": "Pl",
    "titleName": "",
    "mobileNumber": "0999999999",
    "createdAt": "2023-08-17T21:53:56+07:00"
}
```
#### Bill
- METHOD GET `http://localhost:8080/api/bills`
#### Response
```json
[
    {
        "id": 1,
        "billerId": 98499,
        "reference1": 300000025751,
        "reference2": 1733084,
        "trasactionId": "98099310720200530183382100",
        "channelCode": "MOB",
        "senderBankCode": "Kbank",
        "status": "already_paid",
        "tranAmount": 120,
        "createdAt": "2023-08-16T15:06:05Z",
        "updatedAt": "2023-08-17T15:39:57Z"
    },
    {
        "id": 2,
        "billerId": 98499,
        "reference1": 1,
        "reference2": 17003,
        "trasactionId": "98099310720200530183382100",
        "channelCode": "MOB",
        "senderBankCode": "Kbank",
        "status": "already_paid",
        "tranAmount": 210,
        "createdAt": "2023-08-17T13:12:19Z",
        "updatedAt": "2023-08-17T13:15:19Z"
    }
]
```
- METHOD POST `http://localhost:8080/api/bills`
- BODY
```json
{
    "billerId": 98499,
    "reference1": 1,
    "reference2": 170049,
    "status": "waiting",
    "tranAmount": 150
}
```
#### Response
```json
{
    "id": 1,
    "billerId": 98499,
    "reference1": 1,
    "reference2": 170049,
    "trasactionId": "",
    "channelCode": "",
    "senderBankCode": "",
    "status": "waiting",
    "tranAmount": 150,
    "createdAt": "2023-08-18T16:35:04+07:00",
    "updatedAt": "2023-08-18T16:35:04+07:00"
}
```
- METHOD PUT `http://localhost:8080/api/bills/:id`
- BODY
```json
{
    "billerId": 98499,
    "reference1": 1,
    "reference2": 170049,
    "status": "waiting",
    "tranAmount": 220
}
```
#### Response
```json
{
    "id": 1,
    "billerId": 98499,
    "reference1": 1,
    "reference2": 170049,
    "trasactionId": "",
    "channelCode": "",
    "senderBankCode": "",
    "status": "waiting",
    "tranAmount": 220,
    "createdAt": "2023-08-18T16:35:04+07:00",
    "updatedAt": "2023-08-18T16:36:23+07:00"
}
```
- METHOD GET `http://localhost:8080/api/bills/search?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`
#### Response
````json
[
    {
        "id": 1,
        "billerId": 98499,
        "reference1": 300000025751,
        "reference2": 1733084,
        "trasactionId": "98099310720200530183382100",
        "channelCode": "MOB",
        "senderBankCode": "Kbank",
        "status": "already_paid",
        "tranAmount": 120,
        "createdAt": "2023-08-16T15:06:05Z",
        "updatedAt": "2023-08-17T15:39:57Z"
    },
    {
        "id": 2,
        "billerId": 98499,
        "reference1": 1,
        "reference2": 17003,
        "trasactionId": "98099310720200530183382100",
        "channelCode": "MOB",
        "senderBankCode": "Kbank",
        "status": "already_paid",
        "tranAmount": 210,
        "createdAt": "2023-08-17T13:12:19Z",
        "updatedAt": "2023-08-17T13:15:19Z"
    }
]
````

## Conference
- [KBank API (Bill-Payment)](https://apiportal.kasikornbank.com/product/public/All/Bill%20Payment/Introduction/%E0%B9%80%E0%B8%A3%E0%B8%B4%E0%B9%88%E0%B8%99%E0%B8%95%E0%B9%89%E0%B8%99%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%87%E0%B8%B2%E0%B8%99)