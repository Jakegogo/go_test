package wrap

import (
	"testing"
)

func TestWrapper(t *testing.T) {
	//u := &UploadInform{}
	//w := &Wraps_impVideoUploadInform{Orign:u}
	//caller(w)
}

func caller(impl _impVideoUploadInform) {
	impl.InformVideoUploadOver(nil, nil)
}
