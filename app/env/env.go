package env

type Environment struct {
	App     AppEnvironment
	DB      DBEnvironment
	Logs    LogsEnvironment
	Admin   GlobalAdmin // Naive users' permissions
	Network NetEnvironment
}
