package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/devops-platform-backend/usecase"
)

type JenkinsHandler struct {
	JenkinsUsecase usecase.JenkinsUsecase
}

func (j *JenkinsHandler) CheckHealth(c *gin.Context) {
	err := j.JenkinsUsecase.RegisterLunchPayment(c.Request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}

func (j *JenkinsHandler) CreateJob(c *gin.Context) {
	err := j.JenkinsUsecase.RegistJob(c.Request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}

func (j *JenkinsHandler) CreateProject(c *gin.Context) {
	err := j.JenkinsUsecase.RegistProject(c.Request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}
