package payprogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
)

type command struct {
	url        string
	command    string
	key        string
	verbose    bool
	parameters map[string]interface{}
}

func (c *command) Set(key string, value interface{}) *command {
	c.parameters[key] = value
	return c
}

func (c *command) paramsAsJSON() []byte {
	b, _ := json.Marshal(c.parameters)
	return b
}

// When using this method remember to close the response body
func (c *command) RawExecute() (*http.Response, error) {
	var postBytes bytes.Buffer
	w := multipart.NewWriter(&postBytes)

	fw, _ := w.CreateFormField("apikey")
	fw.Write([]byte(c.key))

	fw, _ = w.CreateFormField("command")
	fw.Write([]byte(c.command))

	fw, _ = w.CreateFormField("params")
	fw.Write(c.paramsAsJSON())

	w.Close()

	req, err := http.NewRequest("POST", c.url, &postBytes)
	if err != nil {
		return nil, err
	}

	// set the content-type
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	req.Body.Close()

	if c.verbose {
		b, err := httputil.DumpRequest(req, true)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("Could not dump request got error %s", err))
		} else {
			fmt.Print(string(b))
		}

		b, err = httputil.DumpResponse(res, true)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Could not dump response got error %s", err))
		} else {
			fmt.Print(string(b))
		}
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		return nil, err
	}

	return res, nil
}

func (c *command) Execute(x interface{}) error {
	resp, err := c.RawExecute()
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(x)

	if err != nil { // check if we received an api-error
		var apiErr ApiError
		if err = decoder.Decode(&apiErr); err == nil {
			return errors.New("API:" + apiErr.Message)
		}
	}

	return err
}
