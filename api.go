// Package payprogo provides a simple handler for binding with the
// "https://www.paypro.nl/" payment provider.
// A simple payment command can be executed as follows:
//   func main() {
//      p := payprogo.New("[[api-key]]")
//	var paymentInfo payprogo.PaymentResponse
//      err := p.NewCommand("create_product_payment").Set("product_id", "24611").Set("consumer_email", "mijnklant@mailadres.nl").Execute(&paymentInfo)
//	if err == nil {
//          log.Printf("%v", paymentInfo)
//      }
//   }
package payprogo

// New returns a *PayPro struct which can be used to execute commands
func New(key string) *PayPro {
	return &PayPro{
		key,
		"https://www.paypro.nl/post_api/",
		false,
	}
}

// PayPro Container holding authentication information and used for issuing commands
type PayPro struct {
	key   string
	url   string
	debug bool
}

// Debug sets the api in debug mode. This will also dump the http.Request and
// http.Response variables
func (p *PayPro) Debug(d bool) {
	p.debug = d
}

// NewCommand returns a new command on which you can Set parameters and Execute
func (p *PayPro) NewCommand(c string) *Command {
	r := &Command{
		p.url,
		c,
		p.key,
		p.debug,
		make(map[string]interface{}),
	}

	if p.debug {
		r.Set("test_mode", "true")
	}

	return r
}
