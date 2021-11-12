package _2_Factory

import "time"

type DBConnection interface {
	Connection() error
	GetNow() (*time.Time, error)
	Close() error
}


