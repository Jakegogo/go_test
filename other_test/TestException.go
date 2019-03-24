package other_test

import (
	"fmt"
	"log"
	"os"
)

var (
	logger *log.Logger
)

func init() {
	f, _ := os.Create("log")
	logger = log.New(f, "", log.LstdFlags)
}

func main() {
	fmt.Println("Begin.")

	var sl = []string{"a", "b"}
	defer recoverPanic()
	fmt.Println(sl[3])

	fmt.Println("End.") // NO End.
}

func recoverPanic() {
	if rec := recover(); rec != nil {
		err := rec.(error)
		logger.Printf("Unhandled error: %v\n", err.Error())
		fmt.Fprintf(os.Stderr, "Program quit unexpectedly; please check your logs\n", err.Error())
		//os.Exit(1)
	}
}
