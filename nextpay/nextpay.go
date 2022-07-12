package nextpay

import (
	"encoding/json"
	"strconv"
)

const BaseURL = "https://nextpay.org/nx/gateway"

type NextPay struct {
	apiKey string
}

func New(apiKey string) *NextPay {
	return &NextPay{apiKey}
}

func (next *NextPay) CreateToken(orderID, callbackUri string, amount int, options *CreateTokenOptions) (*CreateTokenResp, error) {
	data := Map{
		"api_key":      next.apiKey,
		"order_id":     orderID,
		"amount":       strconv.Itoa(amount),
		"callback_uri": callbackUri,
	}

	if options != nil {
		if options.Currency != "" {
			data["currency"] = string(options.Currency)
		}

		if options.CustomerPhone != "" {
			data["customer_phone"] = string(options.CustomerPhone)
		}

		if options.CustomFields != nil {
			fields, err := json.Marshal(options.CustomFields)
			if err != nil {
				return nil, err
			}

			data["custom_json_fields"] = string(fields)
		}

		if options.PayerName != "" {
			data["payer_name"] = options.PayerName
		}

		if options.PayerDesc != "" {
			data["payer_desc"] = options.PayerName
		}

		if options.AutoVerify {
			data["auto_verify"] = "yes"
		}

		if options.AllowedCart != "" {
			data["allowed_card"] = options.AllowedCart
		}
	}

	resp, err := post[CreateTokenResp]("/token", data)
	if err != nil {
		return nil, err
	}
	return resp, findErrorByCode(resp.Code, -1)
}

func (next *NextPay) Verify(transactionID string, amount int) (*VerifyResp, error) {
	resp, err := post[VerifyResp]("/verify", Map{
		"api_key":  next.apiKey,
		"trans_id": transactionID,
		"amount":   strconv.Itoa(amount),
	})
	if err != nil {
		return nil, err
	}

	return resp, findErrorByCode(resp.Code, 0)
}

func (next *NextPay) Refund(transactionID string, amount int) (*VerifyResp, error) {
	resp, err := post[VerifyResp]("/verify", Map{
		"api_key":        next.apiKey,
		"trans_id":       transactionID,
		"amount":         strconv.Itoa(amount),
		"refund_request": "yes_money_back",
	})
	if err != nil {
		return nil, err
	}

	return resp, findErrorByCode(resp.Code, -90)
}

// ToDo
// func (next *NextPay) Checkout(wid, amount int, auth string, sheba, name string)
