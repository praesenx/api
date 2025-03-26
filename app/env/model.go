package env

import "fmt"

type Environment struct {
	App     AppEnvironment
	DB      DBEnvironment
	Logs    LogsEnvironment
	Admin   GlobalAdmin // Naive users' permissions
	Network NetEnvironment
}

type NetEnvironment struct {
	HttpHost string `validate:"required,lowercase,min=8"`
	HttpPort string `validate:"required,numeric,oneof=8080"`
}

type AppEnvironment struct {
	Name string `validate:"required,min=4"`
	Type string `validate:"required,lowercase,oneof=local production staging"`
}

type LogsEnvironment struct {
	Level      string `validate:"required,lowercase,oneof=debug info warn error"`
	Dir        string `validate:"required,lowercase"`
	DateFormat string `validate:"required,lowercase"`
}

type GlobalAdmin struct {
	Salt  string `validate:"required,min=8"`
	Token string `validate:"required,sha256"`
}

type DBEnvironment struct {
	UserName      string `validate:"required,lowercase,min=10"`
	UserPassword  string `validate:"required,min=10"`
	DatabaseName  string `validate:"required,lowercase,min=10"`
	Port          int    `validate:"required,numeric,gt=0"`
	PortSecondary int    `validate:"required,numeric,gt=0"`
	Host          string `validate:"required,lowercase,hostname"`
	DriverName    string `validate:"required,lowercase,oneof=postgres"`
	BinDir        string
	URL           string `validate:"required,lowercase,startswith=postgres"`
	SSLMode       string `validate:"required,lowercase,oneof=disable enabled"`
	TimeZone      string `validate:"required"`
}

func (e NetEnvironment) GetHttpPort() string {
	return e.HttpPort
}

func (e NetEnvironment) GetHttpHost() string {
	return e.HttpHost
}

func (e NetEnvironment) GetHostURL() string {
	return e.HttpHost + ":" + e.HttpPort
}

// GetDSN GORM DSN Connection String
func (e DBEnvironment) GetDSN() string {
	return fmt.Sprintf(
		"host=%s users=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		e.Host,
		e.UserName,
		e.UserPassword,
		e.DatabaseName,
		e.Port,
		e.SSLMode,
		e.TimeZone,
	)
}
