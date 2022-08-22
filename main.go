package main

var _mainAppVersion = "v0.1.3"
var _mainInternalBuildNumber = "10009"

func main() {
	// fmt.Println("fyneApp gRPC")

	grpcClientInit()

	fyneApp()

	grpcClientCleanup()

}
