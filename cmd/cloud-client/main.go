package main

import (
	"github.com/byteso/Xcloud/api/cloud-client/v1/router"
	"github.com/byteso/Xcloud/internal/auth"
	"github.com/byteso/Xcloud/internal/config"
	"github.com/byteso/Xcloud/internal/database"
)

func main() {
	//init config
	config.InitConfig()

	// init auth
	auth.GenerateKeyForRsa()

	// init database
	database.InitEngine()
	defer database.Close()

	// init Object Server
	database.InitObjectServerEngine()

	// init router
	router.Router()
}
