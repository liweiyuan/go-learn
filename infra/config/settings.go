package config

import (
	"github.com/spf13/viper"
	"fmt"
)

var (
	MatchesURLGroupPhase = "https://www.fifa.com/worldcup/matches/?#groupphase"
)

func init() {
	viper.SetConfigName("settings")
	viper.AddConfigPath("E:/gopath/src/github.com/liweiyuan/go-learn/infra/config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetMysqlGreConfig() string {

	var (
		dbName   string
		port     string
		user     string
		password string
		host     string
	)

	dbName = viper.GetString("db.dbname")
	port = viper.GetString("db.port")
	user = viper.GetString("db.user")
	password = viper.GetString("db.password")
	host = viper.GetString("db.host")
	//return fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", host, user, dbName, port, password)

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)
}
