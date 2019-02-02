package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx"
	pb "github.com/patterns/fhirbuffer"
)

var (
	databaseConfig *pgx.ConnConfig
)

type fhirbuffer struct{}

func (s *fhirbuffer) Read(ctx context.Context, req *pb.Search) (*pb.Record, error) {

	conn, err := pgx.Connect(*databaseConfig)
	if err != nil {
		log.Printf("Database connection, %v", err)
		return &pb.Record{}, err
	}
	defer conn.Close()

	qr := conn.QueryRow("SELECT PUBLIC.fhirbase_read( $1 , $2 )", req.Type, req.Id)
	return s.runStmt(ctx, qr)
}

func (s *fhirbuffer) Update(ctx context.Context, req *pb.Change) (*pb.Record, error) {

	conn, err := pgx.Connect(*databaseConfig)
	if err != nil {
		log.Printf("Database connection, %v", err)
		return &pb.Record{}, err
	}
	defer conn.Close()

	qr := conn.QueryRow("SELECT PUBLIC.fhirbase_update( $1 )", req.Resource)
	return s.runStmt(ctx, qr)
}

func (s *fhirbuffer) runStmt(ctx context.Context, qryrow *pgx.Row) (*pb.Record, error) {
	var resource string

	err := qryrow.Scan(&resource)
	if err != nil {
		log.Printf("Database error, %v", err)
		return &pb.Record{}, err
	}

	resultset := &pb.Record{
		Resource: []byte(resource),
	}
	return resultset, nil
}

func newServer() *fhirbuffer {
	readDatabaseConf()
	s := &fhirbuffer{}
	return s
}

func readDatabaseConf() {
	url := os.Getenv("DATABASE_URL")
	host := os.Getenv("DB_HOST")
	if (len(host) == 0 && len(url) == 0) {
		// Hardcoded default conf fields for local development environment
		databaseConfig = &pgx.ConnConfig{Host: "127.0.0.1", User: "postgres", Password: "postgres", Database: "fhirbase"}
	} else if len(url) == 0 {
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		name := os.Getenv("DB_NAME")
		databaseConfig = &pgx.ConnConfig{Host: host, User: user, Password: pass, Database: name}
	} else {

		conf, err := pgx.ParseConnectionString(url)
		if err != nil {
			log.Fatal(err)
		}
		databaseConfig = &conf
	}
}
