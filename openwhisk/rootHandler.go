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

type actionWrapperResponse struct {
	__ow_method  string                 `json:"__ow_method,omitempty"`
	__ow_query   string                 `json:"__ow_query,omitempty"`
	__ow_body    string                 `json:"__ow_body,omitempty"`
	__ow_headers map[string]interface{} `json:"__ow_headers,omitempty"`
	__ow_path    string                 `json:"__ow_path,omitempty"`
}

func (ap *ActionProxy) rootHandler(w http.ResponseWriter, r *http.Request) {

	jsonByte, err := preProcess(r)
	if err != nil {
		fmt.Printf("%s", jsonByte)
	} else {
		fmt.Println(err)
	}

}

// preProcess transforms a request in an action value
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

// https://github.com/apache/openwhisk/blob/master/docs/webactions.md
// prende un json e ritorna una response
func postProcess(bt []byte, w http.ResponseWriter) (res *http.Response) {

	fmt.Printf("%s", bt)

	awr := actionWrapperResponse{}

	err := json.NewDecoder(res.Body).Decode(&awr)
	if err != nil {
		panic(err)
	}

	awrJSON, err := json.Marshal(awr)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(awrJSON)

	return res

}
