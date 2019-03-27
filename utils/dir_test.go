package sexpr

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestRel(t *testing.T) {

	pwd, _ := os.Getwd()

	fmt.Println(filepath.Rel("/User/yongfuchen", pwd))

	fmt.Println(filepath.Rel("/User/yongfuchen", "/User/yongfuchen/"))

	fmt.Println(filepath.Rel("/User/yongfuchen", "/root/yongfuchen/"))

	fmt.Println(filepath.Abs("../"))
}
