package main

var _mainAppVersion = "v0.1.5"
var _mainInternalBuildNumber = "10011"

func main() {
	// fmt.Println("fyneApp gRPC")

	grpcClientInit()

	fyneApp()

	grpcClientCleanup()

}
