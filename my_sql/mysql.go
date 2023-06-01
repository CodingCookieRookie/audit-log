package my_sql

import (
	"database/sql"
	"os"
	"sync"

	"github.com/CodingCookieRookie/audit-log/log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var dbName = "sql12622940"
var once sync.Once

func Init() {
	once.Do(func() {
		newDb, err := sql.Open("mysql", getMySqlUri())
		if err != nil {
			panic(err.Error())
		}
		db = newDb

		if err := Exec("CREATE DATABASE IF NOT EXISTS " + dbName); err != nil {
			log.Errorf("error creating database, error: %v", err)
		}

		if err := Exec("USE " + dbName); err != nil {
			log.Errorf("error using table, error: %v", err)
		}

		if err := createEventsTable(); err != nil {
			log.Errorf("error creating table, error: %v", err)
		}
	})
}

func getMySqlUri() string {
	uri := os.Getenv("MYSQL_URI")
	if uri == "" {
		uri = "root:@tcp(localhost:3306)/"
	}
	return uri
}

func Exec(query string, args ...interface{}) error {
	_, err := db.Exec(query, args...)
	return err
}

func Query[T any](
	fieldPtrs func(*T) []interface{},
	s string, args ...interface{},
) ([]*T, error) {
	rows, err := db.Query(s, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*T

	for rows.Next() {
		var t T
		err = rows.Scan(fieldPtrs(&t)...)
		if err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}

func QueryRow[T any](
	fieldPtrs func(*T) []interface{},
	s string, args ...interface{},
) (*T, error) {
	var t T
	err := db.QueryRow(s, args...).Scan(fieldPtrs(&t)...)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func returnPlaceHolderString(args []any) string {
	if len(args) == 0 {
		return ""
	}
	var str []byte = make([]byte, len(args)*2+1)
	str[0] = '('
	for i := range args {
		str[i*2+1] = '?'
		str[i*2+2] = ','
	}
	str[len(str)-1] = ')'
	return string(str)
}
