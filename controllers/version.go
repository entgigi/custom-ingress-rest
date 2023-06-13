package controllers

import (
	"net/http"

	"github.com/entgigi/custom-ingress-rest/version"
	"github.com/gin-gonic/gin"
)

type Version struct {
	BuildTime string `json:"buildTime"`
	Commit    string `json:"commit"`
	Release   string `json:"release"`
}

// https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications

func GetVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"version": Version{version.BuildTime, version.Commit, version.Release}})
}
