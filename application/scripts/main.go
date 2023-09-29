package main

import (
	"bistleague-be/application"
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/services/repository/admin"
	admin2 "bistleague-be/services/usecase/admin"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	rsc, err := application.NewCommonResource(cfg, ctx)
	if err != nil {
		panic(err)
	}
	repo := admin.New(cfg, rsc.Db, rsc.QBuilder)
	uc := admin2.New(cfg, repo, nil, nil)
	_, err = uc.InsertNewAdmin(ctx, dto.RegisterAdminRequestDTO{
		Username: "bistmin",
		Password: "g4C0rb1St1!",
		FullName: "admin bistlig",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("admin has been created!")
}
