package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx"
	pb "github.com/patterns/fhirbuffer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *fhirbuffer) Create(ctx context.Context, req *pb.Change) (*pb.Record, error) {
	conn, err := pgx.Connect(*databaseConfig)
	if err != nil {
		log.Printf("Database connection, %v", err)
		return &pb.Record{}, err
	}
	defer conn.Close()

	qr := conn.QueryRow("SELECT PUBLIC.fhirbase_create( $1 )", req.Resource)
	return s.runStmt(ctx, qr)
}

func (s *fhirbuffer) Delete(ctx context.Context, req *pb.Search) (*pb.Record, error) {
	conn, err := pgx.Connect(*databaseConfig)
	if err != nil {
		log.Printf("Database connection, %v", err)
		return &pb.Record{}, err
	}
	defer conn.Close()

	qr := conn.QueryRow("SELECT PUBLIC.fhirbase_delete( $1 , $2 )", req.Type, req.Id)
	return s.runStmt(ctx, qr)
}

func (s *fhirbuffer) List(req *pb.Search, stream pb.Fhirbuffer_ListServer) error {
	res := strings.ToLower(req.Type)

	if _, ok := resourceHistory[res]; !ok {
		return status.Error(codes.Unimplemented, "Unsupported resource type")
	}

	conn, err := pgx.Connect(*databaseConfig)
	if err != nil {
		return status.Error(codes.Unknown, err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT Resource FROM " + res)
	if err != nil {
		return status.Error(codes.Unknown, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var fetched string
		err := rows.Scan(&fetched)
		switch err {
		case nil:
			rec := &pb.Record{Resource: []byte(fetched)}
			if err := stream.Send(rec); err != nil {
				return status.Error(codes.Unknown, err.Error())
			}
		case pgx.ErrNoRows:
			return status.Error(codes.NotFound, "data not found")
		default:
			return status.Error(codes.Unknown, err.Error())
		}

	}

	if err := rows.Err(); err != nil {
		return status.Error(codes.Unknown, err.Error())
	}
	return nil
}

func (s *fhirbuffer) runStmt(ctx context.Context, qryrow *pgx.Row) (*pb.Record, error) {
	var resource string

	err := qryrow.Scan(&resource)

	switch err {
	case nil:
		resultset := &pb.Record{
			Resource: []byte(resource),
		}
		return resultset, nil

	case pgx.ErrNoRows:
		return &pb.Record{}, status.Error(codes.NotFound, "id was not found")

	default:
		log.Printf("Database error, %v", err)
		return &pb.Record{}, err
	}
}

func newServer() *fhirbuffer {
	readDatabaseConf()
	s := &fhirbuffer{}
	err := s.loadHistory(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func readDatabaseConf() {
	url := os.Getenv("DATABASE_URL")
	host := os.Getenv("DB_HOST")
	if len(host) == 0 && len(url) == 0 {
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
