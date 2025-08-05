package main

import (
	"fmt"
	"os"
)

type DatabaseConnection interface {
	Connect() error
	Close() error
	Query(q string, v any) error
	Exec(cmd string) (int64, error)
}

const (
	pg     = "postgres"
	sqlite = "sqlite"
)

type Postgres struct {
	isConnected bool
}

func (p *Postgres) Connect() error {
	// Simulate connection logic
	p.isConnected = true
	return nil
}

func (p *Postgres) Close() error {
	if p.isConnected {
		fmt.Println("Closing Postgres connection")
		p.isConnected = false
	}
	return nil
}

func (p *Postgres) Query(q string, v any) error {
	if !p.isConnected {
		return fmt.Errorf("not connected")
	}
	fmt.Printf("Running Postgres query: %s\n", q)
	// Simulate query result
	v.(*struct{ id int }).id = 2 // Example of setting a scanned field in the result struct
	return nil
}

func (p *Postgres) Exec(cmd string) (int64, error) {
	if !p.isConnected {
		return 0, fmt.Errorf("not connected")
	}
	fmt.Printf("Executing Postgres command: %s\n", cmd)
	// Simulate affected rows
	return 1, nil
}

type SQLite struct {
	isOpen bool
}

func (s *SQLite) Connect() error {
	// Simulate connection logic
	s.isOpen = true
	return nil
}

func (s *SQLite) Close() error {
	if s.isOpen {
		fmt.Println("Closing SQLite database")
		s.isOpen = false
	}
	return nil
}

func (s *SQLite) Query(q string, v any) error {
	if !s.isOpen {
		return fmt.Errorf("not connected")
	}
	fmt.Printf("Running SQLite query: %s\n", q)
	// Simulate query result
	v.(*struct{ id int }).id = 3 // Example of setting a scanned field in the result struct
	return nil
}

func (s *SQLite) Exec(cmd string) (int64, error) {
	if !s.isOpen {
		return 0, fmt.Errorf("not connected")
	}
	fmt.Printf("Executing SQLite command: %s\n", cmd)
	// Simulate affected rows
	return 1, nil
}

// Usage example
func main() {
	// Create connections
	var db DatabaseConnection
	os.Setenv("DB_TYPE", "postgres")

	switch os.Getenv("DB_TYPE") {
	case pg:
		db = &Postgres{}
	case sqlite:
		db = &SQLite{}
	default:
		fmt.Println("Unsupported database type")
		return
	}

	testResult := struct {
		id int
	}{}
	db.Connect()
	db.Query("SELECT * FROM test", &testResult)
	fmt.Printf("Query result: %d\n", testResult.id)
	db.Close()
}
