package configs

type Config struct {
	BasePath     string `json:"base_path"`
	ApiKey       string `json:"api_key"`
	Timeout      int    `json:"timeout"`
	InsecureSkip bool   `json:"insecure"`
	UseSSL       bool   `json:"use_ssl"`
}
