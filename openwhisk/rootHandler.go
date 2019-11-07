package openwhisk

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	Deadline       int    `json:"deadline,omitempty"`
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

	var s jsonString
	var jsonData []byte
	jsonData, err := json.Marshal(s)
	if err != nil {
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s", jsonData)

	}
	return jsonData
}

/*

// prende un json e ritorna una response
func postProcess(res *http.Response) (jsonString string) {

}
*/
