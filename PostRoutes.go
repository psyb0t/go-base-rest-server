package main

import (
	"github.com/buaazp/fasthttprouter"
)

func PostRoutes(router *fasthttprouter.Router) {
	router.POST("/:db_name", GetDbCollection)
}
