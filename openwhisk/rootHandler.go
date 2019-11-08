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

type value struct {
	Name   string            `json:"name,omitempty"`
	Main   string            `json:"main,omitempty"`
	Code   string            `json:"code,omitempty"`
	Binary bool              `json:"binary,omitempty"`
	Env    map[string]string `json:"env,omitempty"`
}

type jsonString struct {
	Value          value  `json:"value,omitempty"`
	Namespace      string `json:"namespace,omitempty"`
	Action_name    string `json:"action_name,omitempty"`
	Api_host       string `json:"api_host,omitempty"`
	Api_key        string `json:"api_key,omitempty"`
	Activation_id  string `json:"activation_id,omitempty"`
	Transaction_id string `json:"transaction_id,omitempty"`
	Deadline       int64  `json:"deadline,omitempty"`
}

func (ap *ActionProxy) rootHandler(w http.ResponseWriter, r *http.Request) {

	var jsonByte []byte = preProcess(w, r)
	var jsonStr string = fmt.Sprintf("%s", jsonByte)
	if jsonStr != "" {
		fmt.Println(jsonStr)
	}

}

// prede una request e ritorna un json
func preProcess(w http.ResponseWriter, r *http.Request) []byte {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var val value
	// var val map[string]interface{}

	err = json.Unmarshal(b, &val)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n%v\n", err, val)

	output, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err)
	}

	var jsonStr jsonString

	valuStr := bytes.NewBuffer(output).String()
	fmt.Println("valuStr: ", valuStr)

	jsonStr.Value = val

	fmt.Println("jsonStr.Value: ", jsonStr.Value)

	jsonStr.Namespace = os.Getenv("__OW_NAMESPACE")
	jsonStr.Action_name = os.Getenv("__OW_ACTION_NAME")
	jsonStr.Api_host = os.Getenv("__OW_API_HOST")
	jsonStr.Api_key = os.Getenv("__OW_API_KEY")
	jsonStr.Activation_id = os.Getenv("__OW_ACTIVATION_ID")
	jsonStr.Transaction_id = os.Getenv("__OW_TRANSACTION_ID")
	jsonStr.Deadline, err = strconv.ParseInt(os.Getenv("__OW_DEADLINE"), 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	output2, err := json.Marshal(jsonStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("jsonStr.Value: ", output2)

	return output2

}

/*

// prende un json e ritorna una response
func postProcess(res *http.Response) (jsonString string) {

}
*/
