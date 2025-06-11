package boost

import (
	"github.com/joho/godotenv"
	"github.com/oullin/api/env"
	"github.com/oullin/api/pkg"
)

func Spark(envPath string) (*env.Environment, *pkg.Validator) {
	validate := GetDefaultValidate()

	envMap, err := godotenv.Read(envPath)

	if err != nil {
		panic("failed to read the .env file: " + err.Error())
	}

	return MakeEnv(envMap, validate), validate
}
