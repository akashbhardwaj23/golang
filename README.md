# Golang Report Generation Service

## Task

1. Exposes a gRPC API with the following endpoint: [✅]
 - GenerateReport(UserID string) -> (ReportID string, Error) [✅]
2. Periodically triggers a cron job (every 10 seconds is fine for this test) that: [✅]
 - Calls GenerateReport for a few predefined user IDs. [✅]
3. Stores the generated reports in memory (a simple map is fine). [✅]
4. Logs all operations with timestamps.[✅]
5. Add a health check gRPC endpoint: HealthCheck() -> (Status string) [✅]



## Test The Service

The service has both client and server at the same place so you can start the service by just Using the Service Itself as a Client or PostMan as a Client

Ensure that you have <strong>Golang Installed on your System </strong>

- <strong>Clonning the Repo</strong>

```bash
    git clone https://github.com/akashbhardwaj23/golang.git
```

- <strong>Start The Service </strong>

```bash
    go run ./src/main.go
```

<h2 style="text-align:center;border-bottom:0">Or</h2>

## Postman

Use Postman As A Client To Send Request to the grpc Server

- Start Postman

- Select <strong>New</strong> From Request Tab

- Select Grpc from the Tab

![protoselect](/public//postmanselect.png)

- Enter Url In The Box
```bash
    localhost:50051
```
- Select Method Box and Select Import .proto file from the path

![proto](/public/postmanproto.png)

- Select The method you want to Call
    - GenerateReport
    - HealthCheck

- Invoke the method with the correct Parameter
    - GenerateReport - userId
    - HealthCheck - No parameter


## That's All

email - [akashbhardwaj415@gmail.com](mailto:akashbhardwaj415@gmail.com)


