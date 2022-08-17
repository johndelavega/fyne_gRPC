package main

import (
	"fmt"
)

var _mainAppVersion = "v0.1.1"
var _mainInternalBuildNumber = "10002"

func main() {
	fmt.Println("fyne gRPC")

	// grpcClient()

	grpcClientInit()
	fyneApp()
	grpcClientCleanup()

}
