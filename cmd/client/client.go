package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	pb "github.com/patterns/fhirbuffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else vanilla TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The GRPC server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "fhirbuffer", "The server name to verify against the value returned by TLS handshake")

	pid        = flag.String("patient_id", "d3af67c9-0c02-45f2-bc91-fea45af3ee83", "The patient key as UUID string")
	changeFile = flag.String("change_file", "", "The JSON data for update")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		base, err := os.Getwd()
		if err != nil {
			log.Fatalf("Working directory error: %v", err)
		}
		if *caFile == "" {
			*caFile = filepath.Join(base, "test_cert_auth.crt")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("GRPC server connect error: %v", err)
	}
	defer conn.Close()

	client := pb.NewFhirbufferClient(conn)

	updatePatient(client)

	printPatient(client)
}

func printPatient(client pb.FhirbufferClient) {
	req := &pb.Search{Id: *pid, Type: "Patient"}
	resultset, err := client.Read(context.Background(), req)
	if err != nil {
		log.Fatalf("Read, %v", err)
	}
	log.Println(resultset)
}

func updatePatient(client pb.FhirbufferClient) {
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
	resultset, err := client.Update(context.Background(), req)
	if err != nil {
		log.Fatalf("Update, %v", err)
	}
	log.Println(resultset)
}
