package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hitman99/k8s-sandbox/internal/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "performs postgresql database migrations for the application",
	Run: func(cmd *cobra.Command, args []string) {
		l := log.New(os.Stderr, "[migrations] ", log.Ltime)
		m, err := migrate.New(
			"file://./migrations/postgres",
			fmt.Sprintf("postgres://%s:%s@%s/%s?%s",
				utils.MustGetEnv("POSTGRES_USER", l),
				utils.MustGetEnv("POSTGRES_PASS", l),
				utils.MustGetEnv("POSTGRES_HOST", l),
				utils.MustGetEnv("POSTGRES_DB", l),
				utils.MustGetEnv("POSTGRES_URI_ARGS", l)))
		if err != nil {
			l.Fatalf(err.Error())
		}
		err = m.Up()
		if err != nil {
			l.Fatalf(err.Error())
		} else {
			l.Print("migrated")
		}
	},
}
