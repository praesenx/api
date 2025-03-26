package filesmanager

import (
	"fmt"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/logger"
	"log/slog"
	"os"
	"time"
)

type FilesManager struct {
	path            string
	file            *os.File
	logger          *slog.Logger
	LogsEnvironment env.LogsEnvironment
}

func MakeFilesManager(environment env.LogsEnvironment) (logger.Managers, error) {
	manager := FilesManager{}
	manager.LogsEnvironment = environment
	manager.path = manager.DefaultPath()

	resource, err := os.OpenFile(manager.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return FilesManager{}, err
	}

	handler := slog.New(slog.NewTextHandler(resource, nil))
	slog.SetDefault(handler)

	manager.file = resource
	manager.logger = handler

	return manager, nil
}

func (receiver FilesManager) DefaultPath() string {
	logsEnvironment := receiver.LogsEnvironment

	return fmt.Sprintf(
		logsEnvironment.Dir,
		time.Now().UTC().Format(logsEnvironment.DateFormat),
	)
}

func (receiver FilesManager) Close() bool {
	if err := receiver.file.Close(); err != nil {
		receiver.logger.Error("error closing file: " + err.Error())

		return false
	}

	return true
}
