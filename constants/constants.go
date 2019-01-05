package constants

const (
	CurrencyStrBTC = "BTC"
	CurrencyStrEUR = "EUR"
)

const (
	CurrencyBTC int64 = iota
	CurrencyEUR
)

var CurrenciesStrToInt = map[string]int64{
	CurrencyStrBTC: CurrencyBTC,
	CurrencyStrEUR: CurrencyEUR,
}
