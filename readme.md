# go-oremote

toolkit for using oremote.org APIs

```go
instance := gooremote.New("SECRET KEY")

paymentID, err := instance.NewPayment(10000, "https://next.com", "https://webhook.com", "payment description")
if err != nil {
    log.Println(err.Error())
}

log.Println("payment gateway url: ", instance.GetGatewayURL(paymentID))

paid, err := instance.Verify(paymentID)
if err != nil {
    log.Println(err.Error())
}
if !paid {
    log.Printf("payment %s, is not paid yet.\n", paymentID)
}

log.Printf("payment %s, is paid.\n", paymentID)
```

## HTTPS API guide

note that API SECRET KEY should be passed in header as `x-oremote-api-access-token` in `POST` requests.
### create payment
use `POST` method
```
https://api.oremote.org/pay/v1/new
```
response if there is no error:
```json
{
    "ok": true,
    "result": {
        "id": "string-id"
    }
}
```

response if there is an error:
```json
{
    "ok": false,
    "error_code": xxx,
    "error_description": "error description"
}
```

### payment 
`GET` this url, it'll redirect you to payment page
```
https://api.oremote.org/pay/v1/pay/id
```

### verify payment (check if it's paid or not)
use `POST` method
```
https://api.oremote.org/pay/v1/verify/id
```
replace id with corresponding received in payment creation response

response if there is no error and payment is paid:
```json
{
    "ok": true,
    "result": {
        "paid": true
    }
}
```

response if there is no error and payment is not paid:
```json
{
    "ok": true,
    "result": {
        "paid": false
    }
}
```

response if there is an error:
```json
{
    "ok": false,
    "error_code": xxx,
    "error_description": "error description"
}
```