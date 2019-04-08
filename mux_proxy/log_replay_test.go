package sexpr

import (
	logreplay "LogReplay-sdk-go/client/record"
	logreplayhttp "LogReplay-sdk-go/client/record/tars/http"
	"fmt"
	"log"
	"net/http"
	"tars"
	"testing"
)

func TestLogReplay(t *testing.T) {
	var mux *tars.TarsHttpMux = &tars.TarsHttpMux{}
	mux.HandleFunc("/", handle1)

	fmt.Println(*mux)
	logreplayhttp.AddHttpIntercepter(mux, *logreplayhttp.NewHttpHandleIntercepter().AddFirst(
		logreplayhttp.NewRecordHttpIntercepter()).AddFirst(
		logreplay.NewUserFilterIntercepter()))
	fmt.Println(*mux)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
