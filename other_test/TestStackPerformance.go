package other_test

import (
	"fmt"
	"github.com/roserocket/xerrs"
	"time"
)

func main() {

	//....

	var time1 = time.Now().Second()
	fmt.Println(time1)

	for a := 0; a < 1000000; a++ {

		if nil, err := MyFunc1(); err != nil {
			xerrs.Details(err, 5)
		}

	}

	var time2 = time.Now().Second()
	fmt.Println(time2)
	fmt.Println("cost time:%s", (time2 - time1))

}

func MyFunc1() (interface{}, error) {
	return "result", xerrs.New("error occur")
}
