package nextpay

type CreateTokenResp struct {
	Code          int    `json:"code"`
	TransactionID string `json:"trans_id"`
}

func (d *CreateTokenResp) PaymentURL() string {
	return BaseURL + "/payment/" + d.TransactionID
}

type VerifyResp struct {
	Code          int    `json:"code"`
	Amount        int    `json:"amount"`
	OrderID       string `json:"order_id"`
	CardHolder    string `json:"card_holder"`
	CustomerPhone string `json:"customer_phone"`
	ShaparakRefID string `json:"Shaparak_Ref_Id"`
	CustomField   Map    `json:"custom"`
}

type CheckoutResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CallbackData struct {
	TransactionID string `json:"trans_id"`
	OrderID       string `json:"order_id"`
	Amount        int    `json:"amount"`
}

type Currency string

const (
	CurrencyIRT Currency = "IRT"
	CurrencyIRR Currency = "IRR"
)

type CreateTokenOptions struct {
	Currency      Currency
	CustomerPhone string
	CustomFields  *Map
	PayerName     string
	PayerDesc     string
	AutoVerify    bool
	AllowedCart   string
}
