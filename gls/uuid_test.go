package gls

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestUUid(t *testing.T) {

	for i := 0; i < 1000000; i++ {
		// or error handling
		_, err := uuid.NewV4()
		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
			return
		}
	}

}

func TestAutoInc(t *testing.T) {
	inc := NewAutoInc(0, 1)
	for i := 0; i < 1000000; i++ {
		fmt.Print(inc.Id(), " ")
	}
}
