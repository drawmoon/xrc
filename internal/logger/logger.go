package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	once sync.Once
	lv   *slog.LevelVar
)

// Setup initializes the global logger with the specified verbosity level.
func Setup(verbose bool) {
	once.Do(func() {
		lv = &slog.LevelVar{}
		if verbose {
			lv.Set(slog.LevelDebug)
		} else {
			lv.Set(slog.LevelWarn)
		}

		handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: lv, AddSource: false})
		l := slog.New(handler)
		slog.SetDefault(l)
	})
}
