package constants

// Status strings
const (
	LangENStr = "en"
	LangFRStr = "fr"
)

// Statuses
const (
	LangEN int64 = iota
	LangFR
)

var LangIntToStr = map[int64]string{
	LangEN: LangENStr,
	LangFR: LangFRStr,
}
