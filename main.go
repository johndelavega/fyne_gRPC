package main

import (
	"fmt"
)

var _mainAppVersion = "v0.0.0"
var _mainInternalBuildNumber = "10000"

func main() {
	fmt.Println("fyne gRPC")

	grpcApp()
	fyneApp()

}
