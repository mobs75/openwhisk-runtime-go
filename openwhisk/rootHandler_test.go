package openwhisk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func ExampleRootHandler() {

	ts, cur, log := startTestServer("")

	res, _, _ := doGet(ts.URL + "/")

	os.Setenv("name", "TEST-name")
	os.Setenv("main", "TEST-main")
	os.Setenv("code", "TEST-code")
	os.Setenv("binary", "true")
	os.Setenv("env", `{ "hello": "world", "hi": "all"}`)

	os.Setenv("value", `{ "name": "TEST-name", "main": "TEST-main","code":"TEST-code","binary":"true","env": "{ "hello": "world", "hi": "all"}"}`)
	os.Setenv("namespace", "TEST-namespace")
	os.Setenv("action_name", "TEST-action_name")
	os.Setenv("api_host", "TEST-action_name")
	os.Setenv("api_key", "TEST-")
	os.Setenv("activation_id", "TEST-api_key")
	os.Setenv("transaction_id", "TEST-transaction_id")
	os.Setenv("deadline", "9999")

	fmt.Println(res)

	// Output: xxx

	stopTestServer(ts, cur, log)
}

func ExamplePreprocess() {

	var s1 jsonString

	os.Setenv("value", `{"name": "TEST-name", "main": "TEST-main","code":"TEST-code","binary":"true","env": "{ "hello": "world", "hi": "all"}"}`)
	os.Setenv("namespace", "TEST-namespace")
	os.Setenv("action_name", "TEST-action_name")
	os.Setenv("api_host", "TEST-action_name")
	os.Setenv("api_key", "TEST-")
	os.Setenv("activation_id", "TEST-api_key")
	os.Setenv("transaction_id", "TEST-transaction_id")
	os.Setenv("deadline", "9999")

	var jsonData []byte

	jsonData, err := json.Marshal(s1)

	req, err := http.NewRequest("POST", "http://localhost:8088", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", jsonData)

	// Output: xxx
}