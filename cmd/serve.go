package cmd

import (
	"log/slog"

	"github.com/munnaMia/ahlan/utility"
)

func Run() {
	// setting up new logger for app.
	newLogger := utility.NewLogger(false, false)
	slog.SetDefault(newLogger)

	
}
