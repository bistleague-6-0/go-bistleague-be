package config

type Config struct {
	Stage    string         `json:"stage"`
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	Firebase FirebaseConfig `json:"firebase"`
}

type ServerConfig struct {
	Name string `json:"name"`
	Port string `json:"port"`
}

type DatabaseConfig struct {
	Host         string `json:"host"`
	DatabaseType string `json:"type"`
}

type FirebaseConfig struct {
	ProjectID     string `json:"project_id"`
	ApiKey        string `json:"api_key"`
	AuthDomain    string `json:"auth_domain"`
	StorageBucket string `json:"storage_bucket"`
	AppID         string `json:"app_id"`
}
