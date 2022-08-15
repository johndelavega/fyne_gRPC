package main

import (
	"fmt"
)

var _mainAppVersion = "v0.1.0"
var _mainInternalBuildNumber = "10001"

func main() {
	fmt.Println("fyne gRPC")

	grpcClient()
	fyneApp()

}
