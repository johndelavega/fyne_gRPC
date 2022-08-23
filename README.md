# fyne_gRPC
### fyne.io (Golang cross-platform native UI) with gRPC client, includes gRPC server app <br/> <br/>

*This is a demo app* <br/><br/>

![alt text](https://raw.githubusercontent.com/johndelavega/assets/master/fyne_gRPC/screen_shot1.png)
Screenshots running on desktop
<br/> <br/>

**Setup fyne dev environment: (assumes familiarity with fyne.io)** <br/>
  https://developer.fyne.io/started/  <br/> <br/>


**server app (run from Desktop)** <br/>
run ./serverApp/main.go from a separate terminal

```
go run serverApp/main.go
```

**fyne client app (run from Desktop)**  <br/>
run ./main.go from another terminal

```
go run .
```

**Original gRPC code with added int32 message response or reply** <br/>
  https://github.com/grpc/grpc-go/tree/master/examples/helloworld  <br/><br/>
  

Repo is self-contained  <br/>
both client and server app uses local module|folder ./helloworldProto <br/><br/>


**Android build (from Windows machine), create .apk file:** <br/>
 download NDK 
https://github.com/android/ndk/wiki
~~~
 set ANDROID_NDK_HOME=c:\somefolder\ndk\...
~~~
Install fyne command <br/>
~~~
go install fyne.io/fyne/v2/cmd/fyne@latest
~~~


To test the Android fyne client app with a cloud server

~~~
update localhost with IPv4 from grpcClient.go before building
~~~

cd to project folder, then build apk file using fyne command
~~~
cd fyne_gRPC 
fyne package -os android -appID com.fynetest.fyne_gRPC
~~~
~~~
fyne_gRPC.apk is created, about 127MB file size
~~~
