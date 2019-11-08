package openwhisk

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleRootHandler() {

	ts, cur, log := startTestServer("")

	res, _, _ := doGet(ts.URL + "/")

	os.Setenv("value", "JSON")
	os.Setenv("namespace", "__OW_NAMESPACE")
	os.Setenv("action_name", "__OW_ACTION_NAME")
	os.Setenv("api_host", "__OW_API_HOST")
	os.Setenv("api_key", "__OW_API_KEY")
	os.Setenv("activation_id", "__OW_ACTIVATION_ID")
	os.Setenv("transaction_id", "__OW_TRANSACTION_ID")
	os.Setenv("deadline", "__OW_DEADLINE")

	fmt.Println(res)

	// Output: xxx

	stopTestServer(ts, cur, log)
}

func ExamplePreprocess() {

	var jsonData = `{"name":"TEST-name","main":"TEST-main","code":"TEST-code","binary":"true","env":"{"hello":"world","hi":"all"}"}`

	os.Setenv("value", "JSON")
	os.Setenv("namespace", "__OW_NAMESPACE")
	os.Setenv("action_name", "__OW_ACTION_NAME")
	os.Setenv("api_host", "__OW_API_HOST")
	os.Setenv("api_key", "__OW_API_KEY")
	os.Setenv("activation_id", "__OW_ACTIVATION_ID")
	os.Setenv("transaction_id", "__OW_TRANSACTION_ID")
	os.Setenv("deadline", "__OW_DEADLINE")

	value := value{}
	json.Unmarshal([]byte(jsonData), &value)
	fmt.Println(value)

	// Output: xxx

}
