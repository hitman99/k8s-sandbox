package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
)

type PG interface {
	InsertHash(text, hash string) error
}

type stmts struct {
	insertHash *sql.Stmt
}

const insertHash = `INSERT INTO hashes(text, hash) VALUES($1, $2)`

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
		logger.Fatalf("cannot prepare statment %s. %s", query, err.Error())
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
			insertHash: mustPrepare(db, insertHash, logger),
		},
	}
}

func (c *pg) InsertHash(text, hash string) error {
	_, err := c.stmts.insertHash.Exec(text, hash)
	if err != nil {
		if err, ok := err.(*net.OpError); ok {
			if err, ok := err.Err.(*os.SyscallError); ok {
				if err, ok := err.Err.(syscall.Errno); ok {
					// connection error
					if err == syscall.Errno(10061) {
						c.reconnect()
					}
				}
			}
		}
	}
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

func (c *pg) reconnect() {
	db := MustGetConnection(c.user, c.pass, c.host, c.database, c.args, c.logger)
	c.stmts = &stmts{
		insertHash: mustPrepare(db, insertHash, c.logger),
	}
	c.db.Close()
	c.db = db
}
