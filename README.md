# [KageAIO Backend](https://github.com/KageSolutions/Kage-Solutions-MC-Go)

## Requirements

[![GO](https://img.shields.io/badge/Go-007D9C?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

## Set variables
```go
var (
    apiHost = "api.example.com"
    clientKey = "YOUR_CLIENT_API_KEY"
)
```

## Init

```go
func init() {
    go tokens.LaunchClient(apiHost, clientKey)
}
```

## Request New Token

```go
tokens.TokenRequest(tokens.RequestToken{
    Message: "get-token",
    Task:    "sdfdfdfds",
    Site:    "finishline.com",
    Key:     clientKey,
})
```

## Check if token is ready

```go
for {
    token, success := tokens.CheckToken("sdfdfdfds")
    if success {
        log.Println(token)
        break
    } else {
        log.Println("token not received")
    }
    time.Sleep(5000 * time.Millisecond)
}
```
