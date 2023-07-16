package util

const (
	USD = "USD"
	EUR = "EUR"
	INR = "INR"
	CAD = "CAD"
)

//This function returns true id the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, CAD, EUR, INR:
		return true
	}
	return false
}
