package forceexport

import "testing"

func TestTimeNow1(t *testing.T) {
	var timeNowFunc func() (int64, int32)
	GetFunc(&timeNowFunc, "time.now")
	sec, nsec := timeNowFunc()
	if sec == 0 || nsec == 0 {
		t.Error("Expected nonzero result from time.now().")
	}
}


func TestGetg(t *testing.T) {
	var timeNowFunc func() (int64, int32)
	GetFunc(&timeNowFunc, "runtime.getg")
	sec, nsec := timeNowFunc()
	if sec == 0 || nsec == 0 {
		t.Error("Expected nonzero result from time.now().")
	}
}