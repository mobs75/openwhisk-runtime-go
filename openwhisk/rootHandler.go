package openwhisk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type actionWrapper struct {
	Value          map[string]interface{} `json:"value,omitempty"`
	Namespace      string                 `json:"namespace,omitempty"`
	Action_name    string                 `json:"action_name,omitempty"`
	Api_host       string                 `json:"api_host,omitempty"`
	Api_key        string                 `json:"api_key,omitempty"`
	Activation_id  string                 `json:"activation_id,omitempty"`
	Transaction_id string                 `json:"transaction_id,omitempty"`
	Deadline       int64                  `json:"deadline,omitempty"`
}

type actionResponse struct {
	Method  string                 `json:"__ow_method,omitempty"`
	Query   string                 `json:"__ow_query,omitempty"`
	Body    string                 `json:"__ow_body,omitempty"`
	Headers map[string]interface{} `json:"__ow_headers,omitempty"`
	Path    string                 `json:"__ow_path,omitempty"`
}

func (ap *ActionProxy) rootHandler(w http.ResponseWriter, r *http.Request) {

	jsonByte, err := preProcess(r)
	if err != nil {
		fmt.Printf("%s", jsonByte)
	} else {
		fmt.Println(err)
	}

}

// preProcess: transforms a request in an action value
func preProcess(r *http.Request) ([]byte, error) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	var val map[string]interface{}

	var aw actionWrapper

	err = json.Unmarshal(body, &val)
	if err != nil {
		return nil, err
	}

	aw.Value = val
	aw.Namespace = os.Getenv("__OW_NAMESPACE")
	aw.Action_name = os.Getenv("__OW_ACTION_NAME")
	aw.Api_host = os.Getenv("__OW_API_HOST")
	aw.Api_key = os.Getenv("__OW_API_KEY")
	aw.Activation_id = os.Getenv("__OW_ACTIVATION_ID")
	aw.Transaction_id = os.Getenv("__OW_TRANSACTION_ID")
	aw.Deadline, err = strconv.ParseInt(os.Getenv("__OW_DEADLINE"), 10, 64)

	return json.Marshal(aw)

}

// postProcess: transforms json in a response
func postProcess(bt []byte, w http.ResponseWriter) error {

	ar := actionResponse{}

	err := json.Unmarshal(bt, &ar)
	if err != nil {
		return err
	}

	// start set Headers fields
	accept := ar.Headers["accept"].(string)
	w.Header().Set("accept", accept)

	connection := ar.Headers["connection"].(string)
	w.Header().Set("connection", connection)

	contentType := ar.Headers["content-type"].(string)
	w.Header().Set("content-type", contentType)

	host := ar.Headers["host"].(string)
	w.Header().Set("host", host)

	userAgent := ar.Headers["user-agent"].(string)
	w.Header().Set("user-agent", userAgent)
	// end set Headers fields

	// write body
	body := []byte(ar.Body)
	w.Write(body)

	w.WriteHeader(http.StatusOK)

	return err

}
