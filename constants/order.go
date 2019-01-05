package constants

// Status strings
const (
	OrderStatusUnpaidStr     = "unpaid"
	OrderStatusProcessingStr = "processing"
	OrderStatusPaidStr       = "paid"
	OrderStatusClaimedStr    = "claimed"
)

// Statuses
const (
	OrderStatusUnpaid int64 = iota
	OrderStatusProcessing
	OrderStatusPaid
	OrderStatusClaimed
)

var OrderStatusStrToInt = map[string]int64{
	OrderStatusUnpaidStr:     OrderStatusUnpaid,
	OrderStatusProcessingStr: OrderStatusProcessing,
	OrderStatusPaidStr:       OrderStatusPaid,
	OrderStatusClaimedStr:    OrderStatusClaimed,
}
