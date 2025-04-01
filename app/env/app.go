package env

const local = "local"
const staging = "staging"
const production = "production"

type AppEnvironment struct {
	Name string `validate:"required,min=4"`
	Type string `validate:"required,lowercase,oneof=local production staging"`
}

func (e AppEnvironment) isProduction() bool {
	return e.Type == production
}

func (e AppEnvironment) isStaging() bool {
	return e.Type == staging
}

func (e AppEnvironment) isLocal() bool {
	return e.Type == local
}
