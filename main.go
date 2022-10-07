package main

import (
	"github.com/henriquecursino/gateway/common/env"
	"github.com/henriquecursino/gateway/router"
)

func main() {
	env.Load()
	router.Router()
}
