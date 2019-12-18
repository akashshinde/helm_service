package server

import (
	"github.com/gin-gonic/gin"
	"helm_service/helm_actions"
	"net/http"
)

func Start() {
	r := gin.Default()

	r.GET("/charts", func(context *gin.Context) {
		rels, err := helm_actions.ListReleases()
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
		}
		context.JSON(http.StatusOK, rels)
	})

	r.POST("/install", func(context *gin.Context) {
		ns := context.Query("namespace")
		name := context.Query("name")
		chart := context.Query("url")
		release, err := helm_actions.InstallChart(ns, name, chart)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
		}
		context.JSON(http.StatusOK, release)
	})

	r.GET("/template", func(context *gin.Context) {
		name := context.Query("name")
		chart := context.Query("url")
		release, err := helm_actions.RenderManifests(name, chart)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
		}
		context.JSON(http.StatusOK, release)
	})

	r.PUT("/upgrade", func(context *gin.Context) {
		ns := context.Query("namespace")
		name := context.Query("name")
		chart := context.Query("url")
		release, err := helm_actions.UpgradeRelease(ns, name, chart)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
		}
		context.JSON(http.StatusOK, release)
	})

	r.PUT("/rollback", func(context *gin.Context) {
		name := context.Query("name")
		version := context.Query("version")
		res,err := helm_actions.RollbackRelease(name, version)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
		}
		context.JSON(http.StatusOK, res)
	})

	r.Run()
}
