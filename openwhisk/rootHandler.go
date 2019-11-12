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

func (ap *ActionProxy) rootHandler(w http.ResponseWriter, r *http.Request) {

	var jsonByte []byte = preProcess(r)
	var jsonStr string = fmt.Sprintf("%s", jsonByte)
	if jsonStr != "" {
		fmt.Printf("%s", jsonStr)
	}

}

// prede una request e ritorna un json
func preProcess(r *http.Request) []byte {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	var val map[string]interface{}

	err = json.Unmarshal(b, &val)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%v\n%v\n", err, val)
	fmt.Printf("%s", val)

	output, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err)
	}

	var aw actionWrapper

	valuStr := bytes.NewBuffer(output).String()
	fmt.Printf("1")
	fmt.Printf("%s", valuStr)

	aw.Value = val
	fmt.Printf("2")
	fmt.Printf("%s", aw.Value)

	aw.Namespace = os.Getenv("__OW_NAMESPACE")
	aw.Action_name = os.Getenv("__OW_ACTION_NAME")
	aw.Api_host = os.Getenv("__OW_API_HOST")
	aw.Api_key = os.Getenv("__OW_API_KEY")
	aw.Activation_id = os.Getenv("__OW_ACTIVATION_ID")
	aw.Transaction_id = os.Getenv("__OW_TRANSACTION_ID")
	aw.Deadline, err = strconv.ParseInt(os.Getenv("__OW_DEADLINE"), 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	output2, err := json.Marshal(aw)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("3")
	fmt.Printf("%s", output2)

	return output2

}

/*

// prende un json e ritorna una response
func postProcess(res *http.Response) (jsonString string) {

}
*/
