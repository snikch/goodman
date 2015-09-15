package goodman

import "encoding/json"

// Transaction represents a Dredd transaction object.
// http://dredd.readthedocs.org/en/latest/hooks/#transaction-object-structure
type Transaction struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
	FullPath string `json:"fullPath"`
	Request  struct {
		Body    string                 `json:"body"`
		Headers map[string]interface{} `json:"headers"`
		URI     string                 `json:"uri"`
		Method  string                 `json:"method"`
	} `json:"request"`
	Expected *json.RawMessage `json:"expected,omitempty"`
	Real     struct {
		Body       string                 `json:"body"`
		Headers    map[string]interface{} `json:"headers"`
		StatusCode int                    `json:"statusCode"`
	} `json:"real,omitempty"`
	Origin  *json.RawMessage `json:"origin,omitempty"`
	Test    *json.RawMessage `json:"test,omitempty"`
	Results *json.RawMessage `json:"results,omitempty"`
	Skip    bool             `json:"skip,omitempty"`
	Fail    interface{}      `json:"fail,omitempty"`

	TestOrder []string `json:"hooks_modifications"`
}

// AddTestOrderPoint adds a value to the hooks_modification key used when
// running dredd with TEST_DREDD_HOOKS_HANDLER_ORDER enabled.
func (t *Transaction) AddTestOrderPoint(value string) {
	if t.TestOrder == nil {
		t.TestOrder = []string{}
	}
	t.TestOrder = append(t.TestOrder, value)
}
