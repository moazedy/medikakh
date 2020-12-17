package payment // all temp

func RedirectToPay() bool {
	return true
}

func CheckPayment(role string) bool {
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
