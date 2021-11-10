package connectionImpl

import (
	"database/sql"
	"fmt"
	"time"
	//_ "github.com/lib/pq"
)

type Postgres struct {
	dB *sql.DB
}

func (postgres *Postgres) Connection() error {
	connection := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres",
		"12345678",
		"127.0.0.1",
		"5432",
		"postgres",
	)
	dB, err := sql.Open("postgres", connection)
	if err != nil {
		return err
	}
	postgres.dB = dB
	return nil
}

func (postgres *Postgres) GetNow() (*time.Time, error) {
	myTime := &time.Time{}
	err := postgres.dB.QueryRow("select CURRENT_DATE").Scan(myTime)
	if err != nil {
		fmt.Printf("error al leer la fecha del servidor: %v", err)
		return nil, err
	}
	return myTime, nil
}

func (postgres *Postgres) Close() error {
	return postgres.dB.Close()
}