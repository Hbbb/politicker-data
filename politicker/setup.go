package politicker

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

type appEnv struct {
	resource string // members, bills, votes
	chamber  string // house/senate
	persist  bool
	print    bool
	db       *pgx.Conn
	timeout  time.Duration
}

func (app *appEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet("politicker", flag.ContinueOnError)
	fl.StringVar(&app.resource, "resource", "", "Resource type to fetch (members, bills, votes)")
	fl.StringVar(&app.resource, "chamber", "", "Either Senate or House, required when fetching members")
	fl.DurationVar(&app.timeout, "timeout", 30*time.Second, "Import timeout")
	fl.BoolVar(&app.print, "print", false, "Print results to stdout")

	persist := fl.Bool(
		"persist", true,
		"If overriden to false, data will print to stdout instead of writing to a database (default true)")

	if err := fl.Parse(args); err != nil {
		return err
	}

	if *persist == true {
		db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			return err
		}

		app.db = db
	}

	return nil
}
