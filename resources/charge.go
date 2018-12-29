package resources

import (
	"git.iiens.net/edouardparis/town/opennode"
)

type Charge struct {
	PayReq     string `json:"payreq"`
	Amount     int64  `json:"amount"`
	AmountFiat int64  `json:"amount_fiat"`
	Currency   string `json:"currency"`
	Status     string `json:"status"`
	OrderUUID  string `json:"order_uuid"`
}

func NewCharge(charge *opennode.Charge) *Charge {
	return &Charge{
		PayReq:     charge.LightningInvoice.PayReq,
		Amount:     charge.Amount,
		AmountFiat: charge.FiatValue,
		Currency:   charge.Currency,
		Status:     charge.Status,
	}
}
