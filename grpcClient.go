/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "helloworld" // helloworld is a local module using go.mod and replace
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the 'socket address' to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

var (
	conn   *grpc.ClientConn
	err    error
	client pb.GreeterClient

	reply *pb.HelloReply
)

func grpcClientInit() {

	flag.Parse()

	// Set up a connection to the server.
	conn, err = grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client = pb.NewGreeterClient(conn)

	log.Printf("gRPC initialized")

}

func grpcClientCleanup() {
	defer conn.Close()
}

func SayHello() (string, int32) {

	flag.Parse()

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err = client.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("Error: could not greet, make sure ./serverApp/main.go is running: %v", err)
	}

	return reply.GetMessage(), reply.GetRandNum()
}
