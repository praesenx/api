package env

type Environment struct {
	App     AppEnvironment
	DB      DBEnvironment
	Logs    LogsEnvironment
	Network NetEnvironment
}
