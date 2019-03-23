package sexpr

import (
	"fmt"
	echoutil "go_test/echo"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestReplaceField(t *testing.T) {
	var mux *echoutil.MyServMux = &echoutil.MyServMux{}
	mux.HandleFunc("/", handle1)

	fmt.Println(*mux)
	addIntercepter(mux, NewRecordHttpIntercepter())
	fmt.Println(*mux)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func handle1(res http.ResponseWriter, req *http.Request) {
	fmt.Println("handle1")
	//fmt.Fprintf(res, "handle 1")
	res.Write([]byte("handle 1"))

	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println(body)
	res.WriteHeader(500)
	//var sl = []string{"a", "b"}
	//fmt.Println(sl[3]);
}

func handle2(res http.ResponseWriter, req *http.Request) {
	fmt.Println("handle2")
	res.Write([]byte("handle 2"))
}
