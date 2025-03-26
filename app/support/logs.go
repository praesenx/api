package support

import (
	"fmt"
	"github.com/gocanto/blog/app/contracts"
	"github.com/gocanto/blog/app/env"
	"log/slog"
	"os"
	"time"
)

type FileLog struct {
	path            string
	file            *os.File
	logger          *slog.Logger
	LogsEnvironment env.LogsEnvironment
}

func MakeDefaultFileLogs(environment env.LogsEnvironment) (contracts.LogsDriver, error) {
	file := FileLog{}
	file.LogsEnvironment = environment
	file.path = file.DefaultPath()

	resource, err := os.OpenFile(file.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return FileLog{}, err
	}

	logger := slog.New(slog.NewTextHandler(resource, nil))
	slog.SetDefault(logger)

	file.file = resource
	file.logger = logger

	return file, nil
}

func (receiver FileLog) DefaultPath() string {
	env := receiver.LogsEnvironment

	return fmt.Sprintf(
		env.Dir,
		time.Now().UTC().Format(env.DateFormat),
	)
}

func (receiver FileLog) Close() bool {
	if err := receiver.file.Close(); err != nil {
		receiver.logger.Error("error closing file: " + err.Error())

		return false
	}

	return true
}
