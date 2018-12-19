package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Request : the struct for the request by the client
type Request struct {
	LangID int    `json:"lang_id"`
	Code   string `json:"code"`
	Input  string `json:"input"`
}

// ContainerOutput : the struct for the response to the client
type ContainerOutput struct {
	Output  string `json:"output"`
	RunTime int    `json:"runtime"`
	Memory  int    `json:"memory"`
}

var router *gin.Engine

func init() {
	os.Setenv("DOCKER_API_VERSION", "1.39")
}

func main() {
	router = gin.Default()
	router.POST("/compile", compileEndpoint)
	router.Run()
}

func compileEndpoint(c *gin.Context) {
	status := http.StatusOK
	var req Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		status = http.StatusBadRequest
	}

	// TODO
	// - Make a temp file

	// DUMMY DATA
	res := compileCode(req)

	// - delete the temp file
	c.JSON(status, res)
}
