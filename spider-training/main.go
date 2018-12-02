package main

import (
	"leeif.me/Go-fun/spider-training/domain/pexels"
	"leeif.me/Go-fun/spider-training/infra/initial"
	"leeif.me/Go-fun/spider-training/src/model"
)

func init() {
	initial.DBInit()
	createTable()
}

func createTable() {
	initial.DataBase.AutoMigrate(
		&model.Dynasty{},
		&model.Poet{},
		&model.PoetryInfo{},
		&model.PoetryType{},
		&model.ImageSize{},
		&model.ImageAddress{},
	)
}

func main() {
	pexels.Start()
}
