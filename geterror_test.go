package qbsllm

import (
	"fmt"
	"os"
)

func ExampleErrorReturn() {
	log := New(Lnormal, "qbsllm", os.Stdout, testFormat)
	var err error = log.Get().Errora("foo `bar`", 4711)
	fmt.Println(err)
	// Output:
	// qbsllm ERROR <foo `bar:4711`>
	// foo `bar:4711`
}
