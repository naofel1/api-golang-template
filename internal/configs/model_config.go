package configs

import (
	"time"
)

// Config take Connection info and CORS origin info
type Config struct {
	Mariadb    *Mariadb    `json:"mariadb" yaml:"mariadb"`
	Certs      *Certs      `json:"certs" yaml:"certs"`
	Discord    *Discord    `json:"discord" yaml:"discord"`
	AWS        *AWS        `json:"aws" yaml:"aws"`
	Host       *Host       `json:"host" yaml:"host"`
	Mailgun    *Mailgun    `json:"mailgun" yaml:"mailgun"`
	Jwt        *Jwt        `json:"jwt" yaml:"jwt"`
	Centrifugo *Centrifugo `json:"centrifugo" yaml:"centrifugo"`
	Jaeger     *Jaeger     `json:"jaeger" yaml:"jaeger"`
	AppInfo    *AppInfo    `json:"app_info" yaml:"app_info"`
	Cors       *Cors       `json:"cors" yaml:"cors"`
	Redis      *Redis      `json:"redis" yaml:"redis"`
	Server     *Server     `json:"server" yaml:"server"`
}

// Host struct get the host info in the config file
type Host struct {
	Mode    string `json:"mode" yaml:"mode"`
	Address string `json:"address" yaml:"address"`
	BaseURL string `json:"base_url" yaml:"base_url"`
	Port    int    `json:"port" yaml:"port"`
}

// Server struct get the server info in the config file
type Server struct {
	ClientTimeout       time.Duration `json:"client_timeout" yaml:"client_timeout"`
	ReadTimeout         time.Duration `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout        time.Duration `json:"write_timeout" yaml:"write_timeout"`
	ShutdownGracePeriod time.Duration `json:"shutdown_grace_period" yaml:"shutdown_grace_period"`
}

// Jwt struct get the JWT Token info in the config file
type Jwt struct {
	RefreshSecret   string        `json:"refresh_secret" yaml:"refresh_secret"`
	TokenDuration   time.Duration `json:"token_duration" yaml:"token_duration"`
	RefreshDuration time.Duration `json:"refresh_duration" yaml:"refresh_duration"`
}

// Certs struct get the certificate info in the config file
type Certs struct {
	PubStudent  string `json:"pub_student" yaml:"pub_student"`
	PrivStudent string `json:"priv_student" yaml:"priv_student"`
	PubAdmin    string `json:"pub_admin" yaml:"pub_admin"`
	PrivAdmin   string `json:"priv_admin" yaml:"priv_admin"`
}

// Redis struct get the redis info in the config file
type Redis struct {
	ConnectionString string `json:"connection_string" yaml:"connection_string"`
	Password         string `json:"password" yaml:"password"`
	SelectedDB       int    `json:"selected_db" yaml:"selected_db"`
}

// AppInfo struct get the app info in the config file
type AppInfo struct {
	Mode string `json:"mode" yaml:"mode"`
}

// Discord struct get the discord in the config file
type Discord struct {
	Channels *DiscordChannels `json:"channels" yaml:"channels"`

	AuthToken string `json:"token" yaml:"token"`
}

// DiscordChannels struct get the discord channels in the config file
type DiscordChannels struct {
	Session string `json:"session" yaml:"session"`
}

// TaskManager struct get the redis info in the config file
type TaskManager struct {
	ConnectionString string `json:"connection_string" yaml:"connection_string"`
	Password         string `json:"password" yaml:"password"`
	SelectedDB       int    `json:"selected_db" yaml:"selected_db"`
	Concurrency      int    `json:"concurrency" yaml:"concurrency"`
	Enabled          bool   `json:"enabled" yaml:"enabled"`
}

// Jaeger struct get the jaeger info in the config file
type Jaeger struct {
	ConnectionString string `json:"connection_string" yaml:"connection_string"`
}

// Centrifugo struct get the centrifugo info in the config file
type Centrifugo struct {
	ConnectionString string `json:"connection_string" yaml:"connection_string"`
	APIKey           string `json:"api_key" yaml:"api_key"`
}

// Mariadb struct get the mariadb info in the config file
type Mariadb struct {
	MariaParams MariaParams `json:"params" yaml:"params"`
	DBName      string      `json:"db_name" yaml:"db_name"`
	Host        string      `json:"host" yaml:"host"`
	Port        string      `json:"port" yaml:"port"`
	Net         string      `json:"net" yaml:"net"`
	User        string      `json:"user" yaml:"user"`
	Password    string      `json:"password" yaml:"password"`
	ParseTime   bool        `json:"parseTime" yaml:"parseTime"`
}

// MariaParams struct get the params info in the config file
type MariaParams struct {
	Charset  string `json:"charset" yaml:"charset"`
	Location string `json:"loc" yaml:"loc"`
}

// AWS struct get the AWS info in the config file
type AWS struct {
	SCWAccessKey string `json:"scw_access_key" yaml:"scw_access_key"`
	SCWSecretKey string `json:"scw_secret_key" yaml:"scw_secret_key"`
	Region       string `json:"region" yaml:"region"`
	Endpoint     string `json:"endpoint" yaml:"endpoint"`
	BucketName   string `json:"bucket_name" yaml:"bucket_name"`
}

// GDrive struct get the Google Drive info in the config file
type GDrive struct {
	ClientID     string `json:"client_id" yaml:"client_id"`
	ClientSecret string `json:"client_secret" yaml:"client_secret"`
}

// MailgunResetPassword struct get the Mailgun reset password info in the config file
type MailgunResetPassword struct {
	Sender  string `json:"sender" yaml:"sender"`
	Subject string `json:"subject" yaml:"subject"`
}

// MailgunMailer struct get the Mailgun mailer info in the config file
type MailgunMailer struct {
	ResetPassword *MailgunResetPassword `json:"reset_password" yaml:"reset_password"`
}

// Mailgun struct get the Mailgun info in the config file
type Mailgun struct {
	MailConfig   *MailgunMailer `json:"mailer" yaml:"mailer"`
	ClientDomain string         `json:"client_domain" yaml:"client_domain"`
	ClientSecret string         `json:"client_secret" yaml:"client_secret"`
}

// Cors struct get the cors info in the config file
type Cors struct {
	AllowedOrigins []string `json:"allowed_origins" yaml:"allowed_origins"`
	AllowedHeaders []string `json:"allowed_headers" yaml:"allowed_headers"`
	AllowedMethods []string `json:"allowed_methods" yaml:"allowed_methods"`
}
