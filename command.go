package payprogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
)

// Command is used to execute commands against the PayPro API
type Command struct {
	url        string
	command    string
	key        string
	verbose    bool
	parameters map[string]interface{}
}

// Set sets a parameter like 'consumer_email'
func (c *Command) Set(key string, value interface{}) *Command {
	c.parameters[key] = value
	return c
}

func (c *Command) paramsAsJSON() []byte {
	b, _ := json.Marshal(c.parameters)
	return b
}

// RawExecute can be used to get the raw unparsed http.Response.
// Remember to close the response body when using this function.
func (c *Command) RawExecute() (*http.Response, error) {
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
		dumpRequestAndResponse(req, res)
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status 200 got: %s", res.Status)
	}

	return res, nil
}

func dumpRequestAndResponse(req *http.Request, res *http.Response) {
	b, err := httputil.DumpRequest(req, true)

	if err == nil {
		fmt.Print(string(b))
	}

	b, err = httputil.DumpResponse(res, true)
	if err == nil {
		fmt.Print(string(b))
	}
}

// Execute will transform the input the same way json.Unmarshal does.
// It takes a pointer to a variable and will read the response json.
// If the api returns an error this is returned as a normal go error
func (c *Command) Execute(x interface{}) error {
	resp, err := c.RawExecute()
	if err != nil {
		return err
	}

	// Copy the response body in buffer to re-read the contents
	buff := &bytes.Buffer{}
	io.Copy(buff, resp.Body)
	resp.Body.Close()

	// check for error first, this will remove the need for error handling
	// in custom response types(for now only Payment but this could be used
	// for Sales etc..)
	var apiErr apiError

	if err = json.Unmarshal(buff.Bytes(), &apiErr); err == nil && apiErr.Errors == true {
		return errors.New("Error from API:" + apiErr.Message)
	}

	err = json.Unmarshal(buff.Bytes(), x)

	return err
}
