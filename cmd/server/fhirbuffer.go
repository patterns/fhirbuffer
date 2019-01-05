package main

import (
	"context"
	"log"

	"github.com/jackc/pgx"
	pb "github.com/patterns/fhirbuffer"
)

var (
	databaseConfig *pgx.ConnConfig = &pgx.ConnConfig{Host: "127.0.0.1", User: "postgres", Password: "postgres", Database: "fhirbase"}
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
	s := &fhirbuffer{}
	return s
}
