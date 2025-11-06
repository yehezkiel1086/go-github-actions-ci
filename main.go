package main

import (
	"context"
	"fmt"

	"github.com/yehezkiel1086/go-github-actions-ci/config"
	"github.com/yehezkiel1086/go-github-actions-ci/controller"
	"github.com/yehezkiel1086/go-github-actions-ci/model"
	"github.com/yehezkiel1086/go-github-actions-ci/router"
	"github.com/yehezkiel1086/go-github-actions-ci/storage/postgres"
)

func main() {
	// load .env config
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ .env config loaded successfully")

	// init context
	ctx := context.Background()

	// connect to database
	db, err := postgres.InitDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ database connected successfully")

	// migrate database
	err = db.Migrate(&model.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ database migrated successfully")

	// dependency injection
	userController := controller.InitUserController(db)

	// init router
	r := router.InitRouter(userController)

	if err := r.Serve(conf.HTTP); err != nil {
		panic(err)
	}
}
