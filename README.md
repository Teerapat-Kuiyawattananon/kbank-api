# KBank-API
KBank-API Projects using the Kasikorn Bank's bill payment API.

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

4. Call Swagger api:
```
curl localhost:8080/swagger/doc.json
```

5. Open SwaggerUI in browser by using URL:
```
http://localhost:1150/swagger/index.html
```
