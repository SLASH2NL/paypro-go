# Go binding for PayPro 

Bindings for the PayPro payment provider in go
 
### Installation

    go get github.com/SLASH2NL/payprogo
    
### Usage

    package main
    import (
           "log"
           "github.com/SLASH2NL/payprogo"
    )

    func main() {
           p := payprogo.New("[[api-key]]")
           p.Debug(true) // to see the http-dump and set PayPro in test_mode

       // simple test payment
           c := p.NewCommand("create_product_payment")
           c.Set("consumer_email", "mijnklant@mailadres.nl")
           c.Set("psp_code", "0021").Set("product_id", "24611")

        var payment payprogo.PaymentResponse
           err := c.Execute(&payment) 
           err := c.Execute(&payment)

           if err != nil {
                   log.Fatal(err)
           }

           log.Printf("%v", payment)
    }
