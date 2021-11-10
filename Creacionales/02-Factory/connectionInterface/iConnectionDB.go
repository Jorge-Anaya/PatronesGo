package connectionInterface

import "time"

type DBConnection interface {
	Connection() error
	GetNow() (*time.Time, error)
	Close() error
}


