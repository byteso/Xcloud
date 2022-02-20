package main

import (
	"github.com/byteso/Xcloud/api/cloud-client/v1/router"
	"github.com/byteso/Xcloud/internal/auth"
)

func main() {

	// init auth
	auth.GenerateKeyForRsa()

	// init cloud-client router
	router.Router()
}
