package main

import (
	"github.com/liweiyuan/go-learn/infra/model"
	"github.com/liweiyuan/go-learn/infra/init"
	"github.com/liweiyuan/go-learn/infra/config"
	"github.com/liweiyuan/go-learn/infra/download"
	"fmt"
	"github.com/liweiyuan/go-learn/domain"

	"github.com/liweiyuan/go-learn/ui"

	_ "github.com/liweiyuan/go-learn/docs"
)

func main() {

	// 1. create table by gorm auto migrate
	StartTables()

	// 2. push data into db
	//PushDataIntoDB()

	//3. start server
	api_server.New().Start()
}

func StartTables() {

	initiator.INSTANCE.AutoMigrate(&model.Match{})
}

func PushDataIntoDB() {

	//matches
	urlMatches := config.MatchesURLGroupPhase

	docMatches, err := download.Download(urlMatches)

	if err != nil {
		fmt.Println(err)
	}
	domain.MatchesGroupPhase(docMatches)

}
