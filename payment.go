package payprogo

import (
	"encoding/json"
	"strconv"
)

// Payment holds the payment_url and payment_hash returned from the
// PayPro api
type Payment struct {
	URL  string `json:"payment_url"`
	Hash string `json:"payment_hash"`
}

// _payment avoids recursion while unmarshalling
type _payment Payment

func (e *Payment) UnmarshalJSON(b []byte) error {
	var intermediate struct {
		_payment `json:"return"`
	}

	err := json.Unmarshal(b, &intermediate)
	if err != nil {
		return err
	}

	*e = Payment(intermediate._payment)

	return nil
}

// CreateSimplePayment issues a simple 'create_payment' request
func (p *PayPro) CreateSimplePayment(amount int, customerEmail, returnURL, description string) (*Payment, error) {

	c := p.NewCommand("create_payment")
	c.Set("amount", strconv.Itoa(amount))
	c.Set("consumer_email", customerEmail)
	c.Set("description", description)
	c.Set("return_url", returnURL)

	var payment Payment
	err := c.Execute(&payment)

	if err != nil {
		return nil, err
	}

	return &payment, err
}
