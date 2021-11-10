package connectionImpl

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MySql struct {
	dB *sql.DB
}

func (mysql *MySql) Connection() error {
	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&allowNativePasswords=true&parseTime=true",
		"root",
		"12345678",
		"127.0.0.1",
		"3306",
		"mysql",
	)
	dB, err := sql.Open("mysql", connection)
	if err != nil {
		return err
	}
	err = dB.Ping()
	if err != nil {
		return err
	}
	mysql.dB = dB
	return nil
}

func (mysql *MySql) GetNow() (*time.Time, error) {
	myTime := &time.Time{}
	err := mysql.dB.QueryRow("select CURDATE()").Scan(myTime)
	if err != nil {
		fmt.Printf("error al leer la fecha del servidor: %v", err)
		return nil, err
	}
	return myTime, nil
}

func (mysql *MySql) Close() error {
	return mysql.dB.Close()
}
