package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCreateTopic(t *testing.T) {

	samplePayload := []byte(`{"topic_name": "Sulav","queue": ""}`)

	//API endpoint setup
	req, err := http.NewRequest("POST", "/newTopic", bytes.NewBuffer(samplePayload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTopic)

	//Making the API call
	handler.ServeHTTP(rr, req)

	if req.Method != "POST" {
		t.Errorf("Expected 'POST' request, got '%s'", req.Method)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//FIXME: String matching comparison doesnt work
	expectedPayload := `{"status":"Success","msg":"Topic Sulav Added.","err":""}`
	if rr.Body.String() != expectedPayload {
		t.Errorf("newTopic was incorrect, got: %s (%s), want: %s (%s).",
			expectedPayload, reflect.TypeOf(rr.Body.String()), rr.Body.String(), reflect.TypeOf(expectedPayload))
	}

}

func TestSendMessage(t *testing.T) {
	// TODO: Write the tests
}

func TestGetMessage(t *testing.T) {
	// TODO: Etire the tests
}
