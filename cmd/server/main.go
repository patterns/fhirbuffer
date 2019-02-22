// See GRPC beyond the basics guide (grpc.io)
//
//
package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/patterns/fhirbuffer"
	"google.golang.org/grpc"
)

var (
	tlsflag  = flag.Bool("tls", false, "Connection uses TLS if true, else vanilla TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.String("port", "10000", "The server port")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort("localhost", *port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on ", addr)

	opts := enableTLSConfig()
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFhirbufferServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
