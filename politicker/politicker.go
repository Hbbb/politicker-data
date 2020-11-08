package politicker

import (
	"context"
	"fmt"
	"os"
)

func (app *appEnv) run() error {
	var r Response
	err := app.fetchJSON("https://api.propublica.org/congress/v1/116/senate/members.json", &r)
	if err != nil {
		return err
	}

	if app.print {
		app.printResponse(r)
	}

	if app.persist {
		app.saveResponse(r)
	}

	return nil
}

// TODO: This needs to take an interface so we can handle different response types
func (app *appEnv) saveResponse(r Response) error {
	err := app.db.Ping(context.Background())
	if err != nil {
		return err
	}

	// TODO: Save the response

	return nil
}

func (app *appEnv) printResponse(r Response) {
	for _, member := range r.Results[0].Members {
		fmt.Println(member)
	}
}

// CLI exposes the CLI interface to the app
func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Startup error: %e\n", err)
		return 2
	}

	if err := app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %e\n", err)
		return 1
	}
	return 0
}
