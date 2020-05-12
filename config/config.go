package config

/*
 * @Script = config.go
 * @Author = Jayanta
 * @Description = This is description.
 */

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// ConfigFile ...
var ConfigFile *ini.File

// Config ...
var Config = new(ServerConfig)

// ServerConfig ...
type ServerConfig struct {
	RunMode        string
	HTTPPort       int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	Type           string
	User           string
	Password       string
	Host           string
	DbName         string
	TablePrefix    string
	QueueUser      string
	QueuePass      string
	QueueHost      string
	RedisHost      string
	RedisPass      string
	RedisIndex     string
	JwtSecret      string
	JwtTokenExpire int64
	AwsKey         string
	AwsSecret      string
	AwsRegion      string
	AwsBucket      string
	MailHost       string
	MailUsername   string
	MailPassword   string
	MailFrom       string
	MailFromName   string
	MailPort       string
}

// LoadServerConfig ...
func init() {

	var err error

	ConfigFile, err = ini.Load("app.ini")
	if err != nil {
		log.Fatal(2, "Fail to parse 'app.ini': %v", err)
	}

	server, _ := ConfigFile.GetSection("server")
	database, _ := ConfigFile.GetSection("database")
	queue, _ := ConfigFile.GetSection("queue")
	redis, _ := ConfigFile.GetSection("redis")
	jwt, _ := ConfigFile.GetSection("jwt")
	aws, _ := ConfigFile.GetSection("aws")
	mail, _ := ConfigFile.GetSection("mail")

	Config.RunMode = ConfigFile.Section("").Key("RUN_MODE").MustString("debug")
	Config.HTTPPort = server.Key("HTTP_PORT").MustInt()
	Config.ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	Config.WriteTimeout = time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	Config.Type = database.Key("TYPE").MustString("")
	Config.User = database.Key("USER").MustString("")
	Config.Password = database.Key("PASSWORD").MustString("")
	Config.Host = database.Key("HOST").MustString("")
	Config.DbName = database.Key("NAME").MustString("")
	Config.TablePrefix = database.Key("TABLE_PREFIX").MustString("")
	Config.QueueUser = queue.Key("USER").MustString("")
	Config.QueuePass = queue.Key("PASSWORD").MustString("")
	Config.QueueHost = queue.Key("HOST").MustString("")
	Config.RedisHost = redis.Key("HOST").MustString("")
	Config.RedisPass = redis.Key("PASSWORD").MustString("")
	Config.RedisIndex = redis.Key("INDEX").MustString("")
	Config.JwtSecret = jwt.Key("JWTSECRET").MustString("")
	Config.JwtTokenExpire = jwt.Key("TOKEN_EXPIRE").MustInt64(0)

	// S3 config

	Config.AwsKey = aws.Key("S3_KEY").MustString("")
	Config.AwsSecret = aws.Key("S3_SECRET").MustString("")
	Config.AwsRegion = aws.Key("S3_REGION").MustString("")
	Config.AwsBucket = aws.Key("S3_BUCKET").MustString("")

	//Mail Config

	Config.MailHost = mail.Key("MAIL_HOST").MustString("")
	Config.MailUsername = mail.Key("MAIL_USERNAME").MustString("")
	Config.MailPassword = mail.Key("MAIL_PASSWORD").String()
	Config.MailFrom = mail.Key("MAIL_FROM").MustString("")
	Config.MailFromName = mail.Key("MAIL_FROM_NAME").MustString("")
	Config.MailPort = mail.Key("MAIL_PORT").MustString("")

	log.Println("Server Config Initialized with following values")
	log.Printf("%+v\n", Config)
}
