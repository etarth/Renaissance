package config

type Config interface {
	GetServer() Server
	GetDb() Db
	GetJwt() Jwt
	GetAws() Aws
}

type Server struct {
	Name string `mspstructure:"server_name"`
	Env  string `mapstructure:"server_env"`
	Url  string `mapstructure:"server_url"`
	Host string `mapstructure:"server_host"`
	Port int    `mapstructure:"server_port"`
}

type Db struct {
	URI          string `mapstructure:"DB_URI"`
	DatabaseName string `mapstructure:"DB_DATABASE_NAME"`
	Username     string `mapstructure:"DB_USERNAME"`
	Password     string `mapstructure:"DB_PASSWORD"`
}

type Jwt struct {
	ApiSecretKey           string `mapstructure:"jwt_api_secret_key"`
	AccessTokenSecret      string `mapstructure:"jwt_access_token_secret"`
	RefreshTokenSecret     string `mapstructure:"jwt_refresh_token_secret"`
	AccessTokenExpiration  int    `mapstructure:"jwt_access_token_expiration"`
	RefreshTokenExpiration int    `mapstructure:"jwt_refresh_token_expiration"`
}

type Aws struct {
	BucketName      string `mapstructure:"aws_bucket_name"`
	AccessKeyId     string `mapstructure:"aws_access_key_id"`
	SecretAccessKey string `mapstructure:"aws_secret_access_key"`
	Region          string `mapstructure:"aws_region"`
}
