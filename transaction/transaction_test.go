package transaction

import (
	"encoding/json"
	"reflect"
	"testing"
)

type HeadersWrapper struct {
	Headers Headers `json:"headers,omitempty"`
}

func TestSettingFailToString(t *testing.T) {
	trans := Transaction{}
	trans.Fail = "Fail"

	if trans.Fail != "Fail" {
		t.Errorf("Transaction.Fail member must be able to be set as string")
	}
}

func TestHeadersJsonUnmarshal(t *testing.T) {
	data := []byte(`{
	    "headers":{
		"User-Agent":"Dredd/1.4.0 (Darwin 15.4.0; x64)",
		"Content-Length":11,
		"Set-Cookie":["Test=Yo"]
	    }
	}`)

	hdrs := &HeadersWrapper{}
	err := json.Unmarshal(data, hdrs)
	if err != nil {
		t.Errorf("Error %v", err)
	}

	if !reflect.DeepEqual(hdrs.Headers["Set-Cookie"], []string{"Test=Yo"}) {
		t.Errorf("Set-Cookie should be slice with length of 1")
	}

	if !reflect.DeepEqual(hdrs.Headers["User-Agent"], []string{"Dredd/1.4.0 (Darwin 15.4.0; x64)"}) {
		t.Errorf("User-Agent should be converted to a slice")
	}

	if !reflect.DeepEqual(hdrs.Headers["Content-Length"], []string{"11"}) {
		t.Errorf("Content-Length should be converted to a string and be a slice")
	}
}

func TestHeadersMarshalJSON(t *testing.T) {
	headers := make(map[string][]string)
	headers["Content-Type"] = []string{"9"}
	headers["User-Agent"] = []string{"Dredd/1.4.0 (Darwin 15.4.0; x64)"}
	headers["Set-Cookie"] = []string{"Test=Yo", "Another=This"}
	wrapper := &HeadersWrapper{
		Headers: headers,
	}

	data, err := json.Marshal(wrapper)

	if err != nil {
		t.Errorf("json.Marshal failed with %v", err)
	}

	var expected interface{}
	err = json.Unmarshal([]byte(`{"headers":{"Content-Type":"9","User-Agent":"Dredd/1.4.0 (Darwin 15.4.0; x64)","Set-Cookie":["Test=Yo","Another=This"]}}`), &expected)

	if err != nil {
		t.Errorf("failed to unmarshal expected json string")
	}

	var actual interface{}
	err = json.Unmarshal(data, &actual)

	if err != nil {
		t.Errorf("failed to unmarshal Marshaled json")
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("json.Marshal failed, data \n %#v did not matched expected \n %#v", actual, expected)

	}
}
