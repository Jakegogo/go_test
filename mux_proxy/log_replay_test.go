package sexpr

import (
	logreplay "LogReplay-sdk-go/client/record/tars"
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
	logreplay.AddIntercepter(mux, logreplay.NewRecordHttpIntercepter())
	fmt.Println(*mux)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
