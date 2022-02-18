package main

import (
	"github.com/byteso/Xcloud/api/cloud-server/v1/router"
	"github.com/byteso/Xcloud/internal/config"
	"github.com/byteso/Xcloud/internal/database"
)

func main() {

	config.InitConfig()

	database.InitEngine()
	defer database.Close()

	router.Router()
}
