package support

type Environment struct {
	App      AppEnvironment
	DB       DBEnvironment
	Admin    GlobalAdmin // Naive users' permissions
	HttpHost string      `validate:"required,lowercase,min=8"`
	HttpPort string      `validate:"required,numeric,oneof=8080"`
}

type AppEnvironment struct {
	Name     string `validate:"required,min=4"`
	Type     string `validate:"required,lowercase,oneof=local production staging"`
	LogLevel string `validate:"required,lowercase,oneof=debug info warn error"`
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
}

func (e Environment) GetHttpPort() string {
	return e.HttpPort
}

func (e Environment) GetHttpHost() string {
	return e.HttpHost
}

func (e Environment) GetHostURL() string {
	return e.HttpHost + ":" + e.HttpPort
}
