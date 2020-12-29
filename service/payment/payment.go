package payment // all temp
import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sinabakh/go-zarinpal-checkout"
	"log"
	"net/http"
)

func RedirectToPay(c *gin.Context, role, email string) (bool, error) {
	price := PaymentPrice(role)
	if price == 0 {
		return false, errors.New("role value is incorrect")
	}

	zp, err := zarinpal.NewZarinpal(MerchandID, true)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "error on payment")
		return false, err
	}

	paymentUrl, authority, statusCode, err := zp.NewPaymentRequest(
		price,
		CallbackUrl,
		Descriptions,
		email,
		"",
	)
	if err != nil {
		log.Println(err)
		if statusCode == -3 {
			return false, errors.New("price value is not valid")
		}

		return false, err
	}
	fmt.Println(authority)
	c.Redirect(302, paymentUrl)

	varified, refId, statusCode, err := zp.PaymentVerification(price, authority)
	log.Println("varified: ", varified)
	log.Println(err)
	if err != nil || !varified {
		c.JSON(http.StatusInternalServerError, err.Error())
		return false, errors.New("error in payment")
	}

	c.JSON(http.StatusOK, refId)
	return true, nil

}

func PaymentPrice(role string) int {
	switch role {
	case "bronze":
		return BronzePrice
	case "silver":
		return SilverPrice
	case "gold":
		return GoldPrice
	default:
		return 0
	}
}
