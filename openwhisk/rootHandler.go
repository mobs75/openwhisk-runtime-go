package openwhisk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type value struct {
	name   string            `json:"name,omitempty"`
	main   string            `json:"main,omitempty"`
	code   string            `json:"code,omitempty"`
	binary bool              `json:"binary,omitempty"`
	env    map[string]string `json:"env,omitempty"`
}

type jsonString struct {
	value          value  `json:"value,omitempty"`
	namespace      string `json:"namespace,omitempty"`
	action_name    string `json:"action_name,omitempty"`
	api_host       string `json:"api_host,omitempty"`
	api_key        string `json:"api_key,omitempty"`
	activation_id  string `json:"activation_id,omitempty"`
	transaction_id string `json:"transaction_id,omitempty"`
	deadline       int    `json:"deadline,omitempty"`
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
		fmt.Println(string(jsonData))

	}
	return jsonData
}

/*

// prende un json e ritorna una response
func postProcess(res *http.Response) (jsonString string) {

}
*/
