package sexpr

import (
	"fmt"
	"testing"
)

func TestOverFlow(t *testing.T) {
	var i = 1000000 * 10000000000000000
	fmt.Println(i)
}
