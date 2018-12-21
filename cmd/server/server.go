// See GRPC beyond the basics guide (grpc.io)
//
//
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/jackc/pgx"
	pb "github.com/patterns/fhirbuffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls = flag.Bool("tls", false, "Connection uses TLS if true, else vanilla TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile = flag.String("key_file", "", "The TLS key file")
	port = flag.Int("port", 10000, "The server port")

	databaseConfig *pgx.ConnConfig = &pgx.ConnConfig{Host: "127.0.0.1", User: "postgres", Password: "postgres", Database: "fhirbase"}
)

type fhirbuffer struct{}

func (s *fhirbuffer) Read(ctx context.Context, req *pb.Search) (*pb.Record, error) {
	var resource string

	conn, err := pgx.Connect(*databaseConfig)
	if err != nil {
		log.Printf("Database connection, %v", err)
		return &pb.Record{}, err
	}

	qerr := conn.QueryRow("select public.fhirbase_read( $1 , $2 )", req.Type, req.Id).Scan(&resource)
	if qerr != nil {
		log.Printf("Database read, %v", err)
		return &pb.Record{}, err
	}

	resultset := &pb.Record{
		Resource: []byte(resource),
	}
	return resultset, nil
}

func newServer() *fhirbuffer {
	s := &fhirbuffer{}
	return s
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on ", addr)
	var opts []grpc.ServerOption
	if *tls {
		base, err := os.Getwd()
		if err != nil {
			log.Fatalf("Working directory error: %v", err)
		}
		if *certFile == "" {
			*certFile = filepath.Join(base, "fhirbuffer.crt")
		}
		if *keyFile == "" {
			*keyFile = filepath.Join(base, "fhirbuffer.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFhirbufferServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
