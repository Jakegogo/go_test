package logrus

import (
	"github.com/Jakegogo/xerrs"
	"testing"
)

func BenchmarkStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithFields(Fields{"stack": xerrs.Stack(nil)}).Infof("test stack")
	}
}