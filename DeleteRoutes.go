package main

import (
	"github.com/buaazp/fasthttprouter"
)

func DeleteRoutes(router *fasthttprouter.Router) {
	router.DELETE("/:db_name", GetDbCollection)
}
