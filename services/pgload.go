package services

import (
	"context"
	"github.com/hitman99/k8s-sandbox/gen/pgload"
	"github.com/hitman99/k8s-sandbox/internal/postgres"
	"github.com/hitman99/k8s-sandbox/internal/utils"
	"log"
)

// pgload service example implementation.
// The example methods log the requests and return zero values.
type pgloadsrvc struct {
	logger *log.Logger
	db     postgres.PG
}

// NewPgload returns the pgload service implementation.
func NewPgload(logger *log.Logger) pgload.Service {
	host := utils.MustGetEnv("POSTGRES_HOST", logger)
	user := utils.MustGetEnv("POSTGRES_USER", logger)
	pass := utils.MustGetEnv("POSTGRES_PASS", logger)
	db := utils.MustGetEnv("POSTGRES_DB", logger)
	uri_args := utils.MustGetEnv("POSTGRES_URI_ARGS", logger)
	return &pgloadsrvc{
		logger: logger,
		db:     postgres.NewPG(host, user, pass, db, uri_args, logger),
	}
}

// Load implements load.
func (s *pgloadsrvc) Load(ctx context.Context, p *pgload.LoadPayload) (*pgload.JSONStatus, error) {
	s.logger.Printf("Loading postgres with %d records", p.Count)
	for i := 0; i < p.Count; i++ {
		text := utils.GetRandomString(10)
		hash := utils.GetSha256(text)
		s.db.InsertHash(text, string(hash))
	}
	return &pgload.JSONStatus{
		Code:   0,
		Status: "OK",
	}, nil
}
