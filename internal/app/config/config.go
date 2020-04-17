package config

type Config struct {
	BindAddr       string `toml:"bind_addr"`
	LogLevel       string `toml:"log_level"`
	DatabaseURL    string `toml:"database_url"`
	ApiTmdbBaseUrl string `toml:"api_tmdb_base_url"`
	ApiTmdbKey     string `toml:"api_tmdb_key"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
	}
}
