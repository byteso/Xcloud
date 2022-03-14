package main

import (
	"github.com/byteso/Xcloud/api/cloud-server/v1/router"
	"github.com/byteso/Xcloud/internal/auth"
	"github.com/byteso/Xcloud/internal/config"
	"github.com/byteso/Xcloud/internal/database"
)

func main() {

	// init config
	config.InitConfig()

	// init auth
	auth.GenerateKeyForRsa()

	// init database
	database.InitEngine()
	defer database.Close()

	// init object server
	database.InitObjectServerEngine()

	// init router
	router.Router()
}
