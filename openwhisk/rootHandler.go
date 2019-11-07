package openwhisk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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

	valueJSON, err := json.Marshal(value{})

	req, err := http.NewRequest("GET", "http://localhost:8088", bytes.NewBuffer(valueJSON))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", valueJSON)

	/*
		var s jsonString
		var jsonData []byte
		jsonData, err := json.Marshal(s)
		if err != nil {
			fmt.Printf("%s", jsonData)
		}
	*/
	return xxx

}

/*

// prende un json e ritorna una response
func postProcess(res *http.Response) (jsonString string) {

}
*/
