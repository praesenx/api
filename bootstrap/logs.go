package bootstrap

import (
	"log/slog"
	"os"
)

func LogInFile(file string) (*slog.Logger, error) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	logger := slog.New(slog.NewTextHandler(f, nil))
	slog.SetDefault(logger)

	return logger, nil
}
