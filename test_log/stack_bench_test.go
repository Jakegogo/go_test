package test_log

import (
	"go_test/xerrs"
	"testing"
	log "github.com/sirupsen/logrus"
)

func BenchmarkStack(b *testing.B) {
	log.SetLevel(log.DebugLevel)
	for i := 0; i < b.N; i++ {
		log.WithFields(log.Fields{"stack": xerrs.Details(xerrs.New("test err"), 30)}).Info("test stack")
	}
}

func TestStack(t *testing.T) {
	log.WithFields(log.Fields{"stack": xerrs.Details(xerrs.New("test err"), 30)}).Debugf("test stack")
}