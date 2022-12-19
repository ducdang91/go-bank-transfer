package main

import (
	"os"
	"time"

	"github.com/ducdang91/go-bank-transfer/infrastructure"
	"github.com/ducdang91/go-bank-transfer/infrastructure/database"
	"github.com/ducdang91/go-bank-transfer/infrastructure/log"
	"github.com/ducdang91/go-bank-transfer/infrastructure/router"
	"github.com/ducdang91/go-bank-transfer/infrastructure/validation"
)

func main() {
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		Validator(validation.InstanceGoPlayground).
		DbSQL(database.InstancePostgres).
		DbNoSQL(database.InstanceMongoDB)

	app.WebServerPort(os.Getenv("APP_PORT")).
		//WebServer(router.InstanceGorillaMux).
		WebServer(router.InstanceGin).
		Start()
}
