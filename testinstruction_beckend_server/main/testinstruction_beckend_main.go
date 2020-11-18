package main

import (
	"jlambert/FenixInception3/FenixTestInstructionBuilder/testinstruction_beckend_server"
	"time"
)

func main() {
	time.Sleep(15 * time.Second)
	testinstruction_beckend_server.BackendServer_main()
}
