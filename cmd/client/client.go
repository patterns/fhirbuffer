package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	pb "github.com/patterns/fhirbuffer"
	"google.golang.org/grpc"
)

var (
	tlsflag            = flag.Bool("tls", false, "Connection uses TLS if true, else vanilla TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The GRPC server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "fhirbuffer", "The server name to verify against the value returned by TLS handshake")

	rid        = flag.String("resource_id", "d3af67c9-0c02-45f2-bc91-fea45af3ee83", "The resource key as UUID string")
	rtype      = flag.String("resource_type", "Patient", "The resource type (default: \"Patient\")")
	changeFile = flag.String("change_file", "", "The JSON data for update")
)

func main() {
	flag.Parse()
	climode := enableCreateDelete()
	opts := enableTLSConfig()
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("GRPC server connect error: %v", err)
	}
	defer conn.Close()

	client := pb.NewFhirbufferClient(conn)

	updateResource(client, climode)

	printResource(client, climode)
}

func printResource(client pb.FhirbufferClient, mode string) {
	req := &pb.Search{Id: *rid, Type: *rtype}
	var err error
	var resultset *pb.Record
	if mode == "delete" {
		resultset, err = client.Delete(context.Background(), req)
	} else {
		resultset, err = client.Read(context.Background(), req)
	}
	if err != nil {
		log.Fatalf("Read, %v", err)
	}
	log.Println(resultset)
}

func updateResource(client pb.FhirbufferClient, mode string) {
	if len(*changeFile) == 0 {
		return
	}
	ap, err := filepath.Abs(*changeFile)
	if err != nil {
		log.Fatalf("Abs path, %v", err)
	}

	// TODO Need the streaming equivalent to send JSON
	json, err := ioutil.ReadFile(ap)
	if err != nil {
		log.Fatalf("ReadFile, %v", err)
	}

	req := &pb.Change{Resource: json}
	var apierr error
	var resultset *pb.Record
	if mode == "create" {
		resultset, apierr = client.Create(context.Background(), req)
	} else {
		resultset, apierr = client.Update(context.Background(), req)
	}
	if apierr != nil {
		log.Fatalf("Update, %v", apierr)
	}
	log.Println(resultset)
}

func enableCreateDelete() string {
	switch strings.ToLower(filepath.Base(os.Args[0])) {
	case "fhirrm":
		// Naming the binary executable "fhirrm" enables delete mode.
		return "delete"
	case "fhirmk":
		// Naming the binary executable "fhirmk" enables create mode.
		return "create"
	default:
		return "default"
	}
}
