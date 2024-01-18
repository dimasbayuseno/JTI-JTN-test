package configuration

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"math/rand"
	"time"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func NewDatabase() (*sql.DB, error) {
	maxPoolOpen := Env().Database.PoolMaxConn
	maxPoolIdle := Env().Database.PoolIdleConn
	maxPollLifeTime := Env().Database.PoolLifeTime

	db, err := sql.Open("postgres", Env().Database.Dsn())
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		err = fmt.Errorf("unable to connect to database, %w", err)
		return nil, err
	}

	// Set database connection pool settings
	db.SetMaxOpenConns(maxPoolOpen)
	db.SetMaxIdleConns(maxPoolIdle)
	db.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	// Check if the connection is still alive
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging the database: %v", err)
		err = fmt.Errorf("unable to connect to database, %w", err)
		return nil, err
	}
	log.Print("Database connection established successfully")
	return db, nil
}
