package openwhisk

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleRootHandler() {

	ts, cur, log := startTestServer("")

	res, _, _ := doGet(ts.URL + "/")

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

	var jsonData = `{"name":"TEST-name","main":"TEST-main","code":"TEST-code","binary":"true","env":"{"hello":"world","hi":"all"}"}`

	os.Setenv("namespace", "TEST-namespace")
	os.Setenv("action_name", "TEST-action_name")
	os.Setenv("api_host", "TEST-action_name")
	os.Setenv("api_key", "TEST-")
	os.Setenv("activation_id", "TEST-api_key")
	os.Setenv("transaction_id", "TEST-transaction_id")
	os.Setenv("deadline", "9999")

	value := value{}
	json.Unmarshal([]byte(jsonData), &value)
	fmt.Printf("%s", value)

	// Output: xxx

}
