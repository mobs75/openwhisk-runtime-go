package openwhisk

import (
	"fmt"
)

func ExampleHello() {

	ts, cur, log := startTestServer("")

	res, _, _ := doGet(ts.URL + "/hello")

	fmt.Println(res)

	// Output: hellox

	stopTestServer(ts, cur, log)
}
