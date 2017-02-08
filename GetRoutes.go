package main

import (
	"github.com/buaazp/fasthttprouter"
)

func GetRoutes(router *fasthttprouter.Router) {
	router.GET("/:db_name", GetDbCollection)
}
