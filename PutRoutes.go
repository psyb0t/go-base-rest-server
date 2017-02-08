package main

import (
	"github.com/buaazp/fasthttprouter"
)

func PutRoutes(router *fasthttprouter.Router) {
	router.PUT("/:db_name", GetDbCollection)
}
