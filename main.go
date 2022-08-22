package main

import (
	"fmt"
)

var _mainAppVersion = "v0.1.3"
var _mainInternalBuildNumber = "10006"

func main() {
	fmt.Println("fyneApp gRPC")

	grpcClientInit()
	fyneApp()
	grpcClientCleanup()

}
