package settings

import (
	"fmt"
	"os"
	"strconv"
)

var environments = map[string]string{
	"production":    "./settings/prod.json",
	"preproduction": "./settings/pre.json",
	"tests":         "./settings/tests.json",
}

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

var settings Settings = Settings{}
var env = "tests"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "tests"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	// content, err := ioutil.ReadFile(environments[env])
	// if err != nil {
	// 	fmt.Println("Error while reading config file", err)
	// }
	settings = Settings{}
	settings.PublicKeyPath = os.Getenv("PublicKeyPath")
	settings.PrivateKeyPath = os.Getenv("PrivateKeyPath")
	i, err := strconv.Atoi(os.Getenv("JWTExpirationDelta"))
	if err != nil {
		panic("Error setup key JWTExpirationDelta")
	}
	settings.JWTExpirationDelta = i

	// jsonErr := json.Unmarshal(content, &settings)
	// if jsonErr != nil {
	// 	fmt.Println("Error while parsing config file", jsonErr)
	// }
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	// if &settings == nil {
	// 	Init()
	// }
	Init()
	return settings
}

func IsTestEnvironment() bool {
	return env == "tests"
}
