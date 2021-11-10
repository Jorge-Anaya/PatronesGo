package _2_Factory

import (
	"github.com/Jorge-Anaya/PatronesGo/Creacionales/02-Factory/connectionInterface"
	"github.com/Jorge-Anaya/PatronesGo/Creacionales/02-Factory/connectionInterface/connectionImpl"
)

func Factory(typeConnection int) connectionInterface.DBConnection {
	switch typeConnection {
	case 1:
		return &connectionImpl.MySql{}
	case 2:
		return &connectionImpl.Postgres{}
	default:
		return nil
	}
}
