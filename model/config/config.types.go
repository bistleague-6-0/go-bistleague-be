package config

type Config struct {
	Stage          string         `json:"stage"`
	Server         ServerConfig   `json:"server"`
	Database       DatabaseConfig `json:"database"`
	Firebase       FirebaseConfig `json:"firebase"`
	Secret         SecretConfig   `json:"secret"`
	ServiceAccount ServiceAccount `json:"service-account-gcp"`
	Storage        StorageConfig  `json:"storage"`
	SMTP           SMTPConfig     `json:"smtp"`
}

type SecretConfig struct {
	TokenSecret string `json:"token_secret"`
	JWTSecret   string `json:"jwt_secret"`
	AdminSecret string `json:"admin_secret"`
	AdminJWT    string `json:"admin_jwt"`
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

type ServiceAccount struct {
	Type                string `json:"type"`
	ProjectID           string `json:"project_id"`
	PrivateKeyID        string `json:"private_key_id"`
	PrivateKey          string `json:"private_key"`
	ClientEmail         string `json:"client_email"`
	ClientID            string `json:"client_id"`
	AuthProviderX509URL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL   string `json:"client_x509_cert_url"`
	UniverseDomain      string `json:"universe_domain"`
}

type StorageConfig struct {
	StorageURlBase string `json:"storage_url_base"`
	ProjectID      string `json:"projectID"`
	BucketName     string `json:"bucketName"`
}

type SMTPConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}
