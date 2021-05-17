// Code generated by protoc-gen-go-spire. DO NOT EDIT.

package test

import (
	pluginsdk "github.com/spiffe/spire-plugin-sdk/pluginsdk"
	grpc "google.golang.org/grpc"
)

func SomeServiceServiceServer(server SomeServiceServer) pluginsdk.ServiceServer {
	return someServiceServiceServer{SomeServiceServer: server}
}

type someServiceServiceServer struct {
	SomeServiceServer
}

func (s someServiceServiceServer) GRPCServiceName() string {
	return "test.SomeService"
}

func (s someServiceServiceServer) RegisterServer(server *grpc.Server) interface{} {
	RegisterSomeServiceServer(server, s.SomeServiceServer)
	return s.SomeServiceServer
}

type SomeServiceServiceClient struct {
	SomeServiceClient
}

func (c *SomeServiceServiceClient) IsInitialized() bool {
	return c.SomeServiceClient != nil
}

func (c *SomeServiceServiceClient) GRPCServiceName() string {
	return "test.SomeService"
}

func (c *SomeServiceServiceClient) InitClient(conn grpc.ClientConnInterface) interface{} {
	c.SomeServiceClient = NewSomeServiceClient(conn)
	return c.SomeServiceClient
}