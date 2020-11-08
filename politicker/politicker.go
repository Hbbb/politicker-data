package politicker

import (
	"context"
	"fmt"
	"os"
)

func (app *appEnv) run() error {
	var response Response
	err := app.fetchJSON("https://api.propublica.org/congress/v1/116/senate/members.json", &response)
	if err != nil {
		return err
	}

	for _, member := range response.Results[0].Members {
		fmt.Println(member)
	}

	return nil
}

func (app *appEnv) saveJSON() error {
	if app.persist != true {
		return nil
	}

	err := app.db.Ping(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// CLI exposes the CLI interface to the app
func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		fmt.Println(err)
		return 2
	}

	if err := app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %e\n", err)
		return 1
	}
	return 0
}
