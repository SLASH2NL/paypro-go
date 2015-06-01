package payprogo

import "strconv"

type Payment struct {
	Url  string `json:"payment_url"`
	Hash string `json:"payment_hash"`
}

type PaymentResponse struct {
	Errors  string  `json:"errors"`
	Payment Payment `json:"return"`
}

func (p *payPro) CreateSimplePayment(amount int, customerEmail, returnUrl, description string) (*Payment, error) {

	c := p.NewCommand("create_payment")
	c.Set("amount", strconv.Itoa(amount))
	c.Set("consumer_email", customerEmail)
	c.Set("description", description)
	c.Set("return_url", returnUrl)

	var payment PaymentResponse
	err := c.Execute(&payment)

	if err != nil {
		return nil, err
	}

	return &payment.Payment, err
}
