package main

import (
	"github.com/unrolled/render"
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"io/ioutil"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

func TestCreateMatch(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(http.HandlerFunc(createMatchHandler(formatter, repo)))
	defer server.Close()

	body := []byte("{\"test\":\"this is valid json, but doesn't conform to server expectations.\"}")

	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body1))
	if err != nil {
		t.Errorf("Error in creating POST request for createMatchHandler: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		t.Errorf("Error in POST to createMatchHandler: %v", err)
	}
	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err !=nil {
		t.Error()
	}

	fmt.
}