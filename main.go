package main

import (
	"fmt"
)

var _mainAppVersion = "v0.1.2"
var _mainInternalBuildNumber = "10004"

func main() {
	fmt.Println("fyne gRPC")

	grpcClientInit()
	fyneApp()
	grpcClientCleanup()

}
