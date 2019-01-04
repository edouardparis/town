package resources

import (
	"github.com/LizardsTown/opennode"
)

type Charge struct {
	PayReq     string `json:"payreq"`
	Amount     int64  `json:"amount"`
	AmountFiat int64  `json:"amount_fiat"`
	Currency   string `json:"currency"`
	Status     string `json:"status"`
	OrderID    string `json:"order_id"`
}

func NewCharge(charge *opennode.Charge) *Charge {
	return &Charge{
		PayReq:     charge.LightningInvoice.PayReq,
		Amount:     charge.Amount,
		AmountFiat: charge.FiatValue,
		Currency:   charge.Currency,
		Status:     charge.Status,
		OrderID:    charge.OrderID,
	}
}
