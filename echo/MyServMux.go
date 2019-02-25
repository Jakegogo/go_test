package echo

import (
	"net/http"
)

type MyServMux struct {
	http.ServeMux
}

//func (mux *MyServMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	h, _ := mux.Handler(r)
//	h.ServeHTTP(w, r)
//}
