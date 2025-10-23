package main

import (
	"log"

	"github.com/adhyttungga/bri-life-testteknikbe/config"
	"github.com/adhyttungga/bri-life-testteknikbe/server"
)

func main() {
	log.Println("Starting server...")

	// Initialize MySQL connection
	db, err := config.NewMySql()
	if err != nil {
		log.Fatalf("failed to connect with MySQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failes to get underlying *sql.DB: %v", err)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("failed to close database connection: %v", err)
		}
	}()

	// Start the server
	s := server.NewServer(db)
	if err := s.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
