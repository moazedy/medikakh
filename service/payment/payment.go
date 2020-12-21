package payment // all temp

func RedirectToPay(role string) bool {
	return checkPayment(role)
}

func checkPayment(role string) bool {
	switch role {
	case "bronze":
		return true
	case "silver":
		return true
	case "gold":
		return true
	default:
		return false
	}
}
