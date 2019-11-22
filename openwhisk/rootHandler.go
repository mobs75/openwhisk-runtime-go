package openwhisk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type ErrResponseRootHandler struct {
	Error string `json:"error"`
}

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
	StatusCode int                    `json:"statusCode,omitempty"`
	Headers    map[string]interface{} `json:"headers,omitempty"`
	Body       string                 `json:"body,omitempty"`
}

func sendErrorRootHandler(w http.ResponseWriter, code int, cause string) {

	errResponse := ErrResponseRootHandler{Error: cause}
	b, err := json.Marshal(errResponse)

	if err != nil {
		b = []byte("error marshalling error response")
		Debug(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
	w.Write([]byte("\n"))
}

func (ap *ActionProxy) rootHandler(w http.ResponseWriter, r *http.Request) {

	// check if you have an action
	if ap.theExecutor == nil {
		sendErrorRootHandler(w, http.StatusInternalServerError, fmt.Sprintf("no action defined yet"))
		return
	}

	// check if the process exited
	if ap.theExecutor.Exited() {
		sendErrorRootHandler(w, http.StatusInternalServerError, fmt.Sprintf("command exited"))
		return
	}

	// call preProcess
	jsonByte, err := preProcess(r)

	if err != nil {
		sendErrorRootHandler(w, http.StatusBadRequest, fmt.Sprintf("Error preprocessing: %v", err))
		return
	}
	Debug("done reading %d bytes", len(jsonByte))

	// remove newlines
	jsonByte = bytes.Replace(jsonByte, []byte("\n"), []byte(""), -1)

	// execute the action
	response, err := ap.theExecutor.Interact(jsonByte)
	// check for early termination
	if err != nil {
		Debug("WARNING! Command exited")
		ap.theExecutor = nil
		sendErrorRootHandler(w, http.StatusBadRequest, fmt.Sprintf("command exited"))
		return
	}

	// call postProcess
	err = postProcess(response, w)
	DebugLimit("received:", response, 120)

	// flush output
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
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

// postProcess: transforms action value in a response
func postProcess(bt []byte, w http.ResponseWriter) error {

	ar := actionResponse{}

	err := json.Unmarshal(bt, &ar)
	if err != nil {
		return err
	}

	// write body
	body := []byte(ar.Body)
	w.Write(body)

	// write header
	for k, v := range ar.Headers {
		w.Header().Set(k, fmt.Sprintf("%v", v))
	}

	// write StatusCode
	if ar.StatusCode != 200 {
		http.Error(w, http.StatusText(ar.StatusCode), ar.StatusCode)
	}

	return err

}
