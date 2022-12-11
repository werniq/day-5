package card

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

// all calls to stripe
type Card struct {
	Secret  string
	Key     string
	Curreny string
}

type Transaction struct {
	TransactionStatusId int
	Amount				int
	Curreny 			string 
	LastFour 			string
	BankReturnCode 		string 
}

func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

// function to charfge a credit card
func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string,  error) {
	// Access to my secret key 
	stripe.Key = c.Secret
	// Create a payment intent
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	// params add metadata("key", "value")

	// get payment intent
	pi, err := paymentintent.New(params)
	if err != nil {
		var message string
		if stripeErr, ok := err.(*stripe.Error); ok {
			message = cardErrorMessage(stripeErr.Code)
		}
		return nil, message, err
	}
	return pi, "", nil
}

func cardErrorMessage(code stripe.ErrorCode) string {
	var message string
	switch code {
	case stripe.ErrorCodeCardDeclined:
		message = "Your card was declined"
	case stripe.ErrorCodeExpiredCard:
		message = "Your card was expired"
	case stripe.ErrorCodeIncorrectCVC:
		message = "Incorrect CVC provided"
	case stripe.ErrorCodeIncorrectZip:
		message = "Incorrect zip/postal code"
	case stripe.ErrorCodeAmountTooLarge:
		message = "The amount is too large to charge your card"
	case stripe.ErrorCodeAmountTooSmall:
		message = "The amount is too large to charge your card"
	case stripe.ErrorCodeBalanceInsufficient:
		message = "Not enough money on balance"
	case stripe.ErrorCodePostalCodeInvalid:
		message = "Your postal code is invalid"
	default:
		message = "Your card was declined"
	}
	
	return message 
}