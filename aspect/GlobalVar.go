package aspect

import (
	"go_test/another"
	"log"
)

var GlobalVarIns *another.GlobalVar = &another.GlobalVar{
	Key: "true",
}

func init() {
	log.Fatal("errorq")
}
