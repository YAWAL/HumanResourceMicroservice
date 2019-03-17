package database

type Config struct {
	Dialect      string `json:"dialect"`
	User         string `json:"user"`
	DataBaseName string `json:"data_base_name"`
	SSLMode      string `json:"disable"`
	Password     string `json:"password"`
}
