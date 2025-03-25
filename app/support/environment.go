package support

type Environment struct {
	App      AppEnvironment
	DB       DBEnvironment
	Admin    GlobalAdmin // Naive users' permissions
	HttpHost string      `validate:"required"`
	HttpPort string      `validate:"required,numeric,oneof=8080"`
}

type AppEnvironment struct {
	Name     string `validate:"required,min=4"`
	Type     string `validate:"required,oneof=local production staging"`
	LogLevel string `validate:"required,oneof=debug info warn error"`
}

type GlobalAdmin struct {
	Salt  string `validate:"required"`
	Token string `validate:"required,sha256"`
}

type DBEnvironment struct {
	UserName      string `validate:"required,alpha"`
	UserPassword  string `validate:"required"` //,min=10
	DatabaseName  string `validate:"required,alpha"`
	Port          int    `validate:"required,numeric,gt=0"`
	PortSecondary int    `validate:"required,numeric,gt=0"`
	Host          string `validate:"required,alphanumunicode"`
	DriverName    string `validate:"required,oneof=postgres"`
	BinDir        string `validate:"required"`
	URL           string `validate:"required,startswith=postgres"`
}

//"localhost:8080"

func (e Environment) GetHttpPort() string {
	return e.HttpPort
}

func (e Environment) GetHttpHost() string {
	return e.HttpHost
}

func (e Environment) GetHostURL() string {
	return e.HttpHost + ":" + e.HttpPort
}
