package config

//Properties Configuration properties based on env variables.
type Properties struct {
	Port           string `env:"MY_APP_PORT" env-default:"3333"`
	Host           string `env:"HOST" env-default:"0.0.0.0"`
	DBHost         string `env:"DB_HOST" env-default:"localhost"`
	DBPort         string `env:"DB_PORT" env-default:"27017"`
	DBName         string `env:"DB_NAME" env-default:"koho"`
	AppHome        string `env:"APP_HOME" env-default:"/Users/rodrigosantos/Documents/Koho/backend"`
	CollectionName string `env:"COLLECTION_NAME" env-default:"loads"`
}

//PropertiesTest Configuration properties based on env variables for test environment.
type PropertiesTest struct {
	Port           string `env:"MY_APP_PORT" env-default:"3333"`
	Host           string `env:"HOST" env-default:"0.0.0.0"`
	DBHost         string `env:"DB_HOST" env-default:"localhost"`
	DBPort         string `env:"DB_PORT" env-default:"27017"`
	DBName         string `env:"DB_NAME" env-default:"kohotest"`
	AppHome        string `env:"APP_HOME" env-default:"/Users/rodrigosantos/Documents/Koho/backend"`
	CollectionName string `env:"COLLECTION_NAME" env-default:"loads"`
}
