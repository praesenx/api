package boost

import (
	"github.com/gocanto/blog/env"
	"github.com/gocanto/blog/pkg"
	"github.com/joho/godotenv"
)

func Spark(envPath string) (*env.Environment, *pkg.Validator) {
	validate := GetDefaultValidate()

	envMap, err := godotenv.Read(envPath)

	if err != nil {
		panic("failed to read the .env file: " + err.Error())
	}

	return MakeEnv(envMap, validate), validate
}
