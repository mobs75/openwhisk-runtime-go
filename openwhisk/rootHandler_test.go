package openwhisk

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
)

func ExampleRootHandler() {

	// set variables
	os.Setenv("__OW_NAMESPACE", "__namespace__")
	os.Setenv("__OW_ACTION_NAME", "__action_name__")
	os.Setenv("__OW_API_HOST", "__api_host__")
	os.Setenv("__OW_API_KEY", "__api_key__")
	os.Setenv("__OW_ACTIVATION_ID", "__activation_id__")
	os.Setenv("__OW_TRANSACTION_ID", "__transaction_id__")
	os.Setenv("__OW_DEADLINE", "__deadline__")

	comp, _ := filepath.Abs("../common/gobuild.py")
	ts, cur, log := startTestServer(comp)

	// inizialize action that return a body
	doInit(ts, initCode("_test/knative1.src", ""))

	// execute doPost to rootHandler
	res, _, _ := doPost(ts.URL+"/", `{}`)
	fmt.Println(res)

	// done
	stopTestServer(ts, cur, log)

	// Output:
	//200 {"ok":true}
	// Hello, World
	//
	// XXX_THE_END_OF_A_WHISK_ACTIVATION_XXX
	// XXX_THE_END_OF_A_WHISK_ACTIVATION_XXX

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

func ExamplePostprocess() {

	//	data := bytes.NewBuffer([]byte(`{"__ow_method":"post",
	//"__ow_query":"name=Jane",
	//"__ow_body":"eyJuYW1lIjoiSmFuZSJ9",
	//"__ow_headers":{"accept":"*/*",
	//"connection":"close",
	//"content-length":"15",
	//"content-type":"application/json",
	//"host":"172.17.0.1",
	//"user-agent":"curl/7.43.0"},
	//"__ow_path": ""}`))

	data := bytes.NewBuffer([]byte(`{
 "statusCode": 200,
 "headers": {"Content-Type": "text/plain"},
 "body": "Hello!"
}`))

	rw := httptest.NewRecorder()

	err := postProcess(data.Bytes(), rw)
	fmt.Println(err)

	fmt.Println(rw.Header())
	fmt.Println(rw.Body)
	fmt.Println(rw.Result().StatusCode)

	// Output:
	// <nil>
	// map[Content-Type:[text/plain]]
	// Hello!
	// 200
}
