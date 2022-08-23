package main

var _mainAppVersion = "v0.1.4"
var _mainInternalBuildNumber = "10010"

func main() {
	// fmt.Println("fyneApp gRPC")

	grpcClientInit()

	fyneApp()

	grpcClientCleanup()

}
