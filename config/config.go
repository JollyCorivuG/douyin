package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

// 服务器配置
type Server struct {
	Ip   string `toml:"ip"`
	Port int    `toml:"port"`
}

// mysql配置
type Mysql struct {
	UserName  string `toml:"username"`
	PassWord  string `toml:"password"`
	Host      string `toml:"host"`
	Port      int    `toml:"port"`
	Database  string `toml:"database"`
	Charset   string `toml:"charset"`
	ParseTime bool   `toml:"parse_time"`
	Loc       string `toml:"loc"`
}

// redis配置
type Redis struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	PassWord string `toml:"password"`
	Database int    `toml:"database"`
}

// 总的配置
type Config struct {
	Server  `toml:"server"`
	MysqlDb Mysql `toml:"mysql"`
	RedisDb Redis `toml:"redis"`
}

var Info Config

func init() {
	if _, err := toml.DecodeFile("/project/douyin/config/conf.toml", &Info); err != nil {
		log.Fatal(err)
	}
}

// mysql连接
func MysqlDbConnectString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Info.MysqlDb.UserName, Info.MysqlDb.PassWord, Info.Ip, Info.MysqlDb.Port, Info.MysqlDb.Database,
		Info.MysqlDb.Charset, Info.MysqlDb.ParseTime, Info.MysqlDb.Loc)
}
