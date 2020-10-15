package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"strings"
	"time"
	"unicode"
)

var dbConn *sqlx.DB

const ConnString = "postgres://postgresql:postgresql@postgresql:5432/livechat-example?sslmode=disable"

func toLowerPothole(input string) string {
	if input == "" {
		return ""
	}

	runes := []rune(input)

	lastWasUpper := true
	var out strings.Builder
	out.WriteRune(unicode.ToLower(runes[0]))
	for _, char := range runes[1:] {
		if unicode.IsUpper(char) {
			if !lastWasUpper {
				out.WriteRune('_')
			}
			out.WriteRune(unicode.ToLower(char))
			lastWasUpper = true
		} else {
			out.WriteRune(char)
			lastWasUpper = false
		}
	}
	return out.String()
}


func Init(maxConns int) error {
	if dbConn != nil {
		return fmt.Errorf("double initialize")
	}
	var err error
	dbConn, err = sqlx.Open("postgres", ConnString)
	if err != nil {
		return errors.Wrapf(err, "could not connect to database")
	}
	dbConn.SetMaxOpenConns(maxConns)
	dbConn.SetConnMaxLifetime(time.Hour)
	dbConn.Mapper = reflectx.NewMapperFunc("db", toLowerPothole)
	return nil
}

func Get(dest interface{}, query string, args ...interface{}) error {
	return dbConn.Get(dest, query, args...)
}

func Exec(query string, args ...interface{}) error {
	_, err := dbConn.Exec(query, args...)
	return err
}

func Select(dest interface{}, query string, args ...interface{}) error {
	return dbConn.Select(dest, query, args...)
}