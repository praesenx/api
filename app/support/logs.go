package support

import (
	"fmt"
	"github.com/gocanto/blog/app/contracts"
	"log/slog"
	"os"
	"time"
)

type FileLog struct {
	path   string
	file   *os.File
	logger *slog.Logger
}

func MakeDefaultFileLogs() (contracts.LogsDriver, error) {
	file := FileLog{}
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
	return fmt.Sprintf("./storage/logs/logs_%s.log", time.Now().UTC().Format("2006_02_01"))
}

func (receiver FileLog) Close() {
	receiver.file.Close()
}
