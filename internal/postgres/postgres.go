package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PG interface {
	InsertHash(text, hash string) error
}

type stmts struct {
	insertHash *sql.Stmt
}

type pg struct {
	db       *sql.DB
	logger   *log.Logger
	host     string
	user     string
	pass     string
	database string
	args     string
	stmts    *stmts
}

func mustPrepare(db *sql.DB, query string, logger *log.Logger) (stmt *sql.Stmt) {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("cannot prepare statment %s. %s", query, err.Error())
	}
	return
}

func NewPG(host, user, password, database, uri_args string, logger *log.Logger) *pg {
	db := MustGetConnection(user, password, host, database, uri_args, logger)
	return &pg{
		db:   db,
		host: host,
		user: user,
		pass: password,
		args: uri_args,
		stmts: &stmts{
			insertHash: mustPrepare(db, `INSERT INTO hashes(text, hash) VALUES($1, $2)`, logger),
		},
	}
}

func (c *pg) InsertHash(text, hash string) error {
	_, err := c.stmts.insertHash.Exec(text, hash)
	return err
}

func MustGetConnection(user, password, host, database, uri_args string, logger *log.Logger) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s/%s?%s", user, password, host, database, uri_args))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL server: %s", err.Error())
	}
	logger.Print("connected to postgres server")
	return db
}
