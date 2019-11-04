package database

// MongoConfig is a structure which conteins details for creating MongoDB client
type MongoConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
