package sexpr

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/", handle1)

	fmt.Println(*mux)
	fmt.Println(*mux)
	http.ListenAndServe("localhost:8000", mux)
}
