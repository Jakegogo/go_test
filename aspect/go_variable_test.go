package aspect

import (
	"fmt"
	"testing"
)

func TestGoVariable(t *testing.T) {

	for i := 1; i <= 10; i++ {
		var n = i
		var finish bool
		go func(n1 int) {
			if n1 == 2 {
				Finish(&finish)
			}
			fmt.Printf("n %d isFinish %t\n", n, finish)
		}(n)
	}

}

func Finish(i *bool) {
	*i = true
}
