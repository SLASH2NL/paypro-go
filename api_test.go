package payprogo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStartPayment(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("command") != "create_payment" {
			http.Error(w, "Expected command create_payment got "+r.FormValue("command"), http.StatusBadRequest)
		}

		p := &Payment{
			URL:  r.FormValue("command"),
			Hash: r.FormValue("apikey"),
		}

		enc := json.NewEncoder(w)
		enc.Encode(map[string]interface{}{
			"return": p,
		})
	}))

	defer ts.Close()

	p := &PayPro{
		"test",
		ts.URL,
		false,
	}
	payment, err := p.CreateSimplePayment(2000, "test@example.com", "http://example.com/return-payment", "example description")
	if err != nil {
		t.Fatal(err)
	}

	if payment.URL != "create_payment" {
		t.Error("Expected echo of command")
	}

	if payment.Hash != "test" {
		t.Error("Expected echo of apikey")
	}
}

func TestValidatePayment(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("command") != "create_payment" {
			http.Error(w, "Expected command create_payment got "+r.FormValue("command"), http.StatusBadRequest)
		}

		p := &Payment{
			URL:  r.FormValue("command"),
			Hash: r.FormValue("apikey"),
		}

		enc := json.NewEncoder(w)
		enc.Encode(map[string]interface{}{
			"return": p,
		})
	}))

	defer ts.Close()

	p := &PayPro{
		"test",
		ts.URL,
		false,
	}
	payment, err := p.CreateSimplePayment(2000, "test@example.com", "http://example.com/return-payment", "example description")
	if err != nil {
		t.Fatal(err)
	}

	if payment.URL != "create_payment" {
		t.Error("Expected echo of command")
	}

	if payment.Hash != "test" {
		t.Error("Expected echo of apikey")
	}
}
