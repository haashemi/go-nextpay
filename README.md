# Go [NextPay](https://nextpay.org/) API Wrapper

[![Go Report Card](https://goreportcard.com/badge/github.com/haashemi/go-nextpay)](https://goreportcard.com/report/github.com/haashemi/go-nextpay) [![Go Reference](https://pkg.go.dev/badge/github.com/haashemi/go-nextpay.svg)](https://pkg.go.dev/github.com/haashemi/go-nextpay)


### Installation
```
go get github.com/haashemi/go-nextpay
```

------

## Usage

### 0- First step
```go
package main

import (
 ...
	"github.com/haashemi/go-nextpay"
)

func main() {
	nx := nextpay.New("My-NextPay-APIKey")
 ...
}
```

### 1- Create Token 
```go 
token, err := nx.CreateToken("MYORDER999", MyCallbackUri, 20000, nil)
// Handle err...
fmt.Println(token.PaymentURL())

////////
// OR //
////////

token, err := nx.CreateToken("MYORDER999", MyCallbackUri, 200_000, &nextpay.CreateTokenOptions{
	Currency:      nextpay.CurrencyIRR,
	CustomerPhone: "09123456789",
	AutoVerify:    true,
	PayerName:     "علی هاشمی",
	CustomFields: &nextpay.Map{
		"Product": "MyProduct",
		"UserID":  "23908758",
	},
	...
})
// Handle err...
fmt.Println(token.PaymentURL())
```

### 2- Verify
```go 
// TransactionID => token.TransactionID from step one
// 20_000 => IRT amount from step one

vr, err := nx.Verify(TransactionID, 20_000)
// Handle err...
fmt.Println(vr)
```

### 2.1- Refund
```go
// TransactionID => token.TransactionID from step one
// 20_000 => IRT amount from step one

rf, err := nx.Refund(token.TransactionID, 20_000)
// Handle err...
fmt.Println(rf)
```

------
