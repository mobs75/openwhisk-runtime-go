package openwhisk

import (
	"bytes"
	"fmt"
	"net/http"
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

	// Output:
	// xxx

	stopTestServer(ts, cur, log)
}

func ExamplePreprocess() {

	os.Setenv("__OW_NAMESPACE", "__namespace__")
	os.Setenv("__OW_ACTION_NAME", "__action_name__")
	os.Setenv("__OW_API_HOST", "__api_host__")
	os.Setenv("__OW_API_KEY", "__api_key__")
	os.Setenv("__OW_ACTIVATION_ID", "__activation_id__")
	os.Setenv("__OW_TRANSACTION_ID", "__transaction_id__")
	os.Setenv("__OW_DEADLINE", "__deadline__")

	data := bytes.NewBuffer([]byte(`{"hello":"world"}`))
	r, _ := http.NewRequest("POST", "", data)
	out, _ := preProcess(r)
	fmt.Printf("%s", out)
	// Output:
	// {"value":{"hello":"world"},"namespace":"__namespace__","action_name":"__action_name__","api_host":"__api_host__","api_key":"__api_key__","activation_id":"__activation_id__","transaction_id":"__transaction_id__"}

}
