package payprogo

import (
	"encoding/json"
	"fmt"
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

// UnmarshalJSON will handle the custom unmarshalling of the payment data
func (e *Payment) UnmarshalJSON(b []byte) error {
	var intermediate struct {
		Payment _payment `json:"return"`
	}

	err := json.Unmarshal(b, &intermediate)
	if err != nil {
		return err
	}

	if intermediate.Payment.URL == "" {
		return fmt.Errorf("The response from PayPro does not contain a payment URL")
	}

	if intermediate.Payment.Hash == "" {
		return fmt.Errorf("The response from PayPro does not contain a payment Hash")
	}

	*e = Payment(intermediate.Payment)

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

// CreateProductPayment issues a 'create_product_payment' request
func (p *PayPro) CreateProductPayment(productid int, amount int, customerEmail, returnURL, description string) (*Payment, error) {

	c := p.NewCommand("create_product_payment")
	c.Set("product_id", strconv.Itoa(productid))
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

// GetPaymentStatus can be used to check the status of the payment. If you have
// a single(non-recurring) payment you can use the sequenceId '1'
func (p *PayPro) GetPaymentStatus(sequenceId int, hash string) (interface{}, error) {
	c := p.NewCommand("get_status")
	c.Set("payment_hash", hash)
	c.Set("sequence_number", sequenceId)

	var payment interface{}
	err := c.Execute(&payment)

	if err != nil {
		return nil, err
	}

	return &payment, err
}
