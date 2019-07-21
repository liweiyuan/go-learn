package main

import (
	"github.com/liweiyuan/go-learn/model"
	"github.com/liweiyuan/go-learn/infra/init"
)

func main() {

	// 1. create table by gorm auto migrate
	StartTables()

	// 2. push data into db
	PushDataIntoDB()
}


func StartTables() {

	initiator.INSTANCE.AutoMigrate(&model.Match{})
}

func PushDataIntoDB() {

}
