package server

import (
	"github.com/gin-gonic/gin"
	"helm_service/helm_actions"
	"net/http"
)

func Routes() {
	r := gin.Default()

	r.GET("/charts", func(context *gin.Context) {
		rels, err := helm_actions.ListReleases()
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
		}
		context.JSON(http.StatusOK, rels)
	})

	r.POST("/install", func(context *gin.Context) {
		name := context.Query("name")
		chart := context.Query("url")
		ns := context.Query("namespace")
		release, err := helm_actions.InstallChart(ns, name, chart)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
		}
		context.JSON(http.StatusOK, release)
	})

	r.Run()
}
