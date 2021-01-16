package uploader

import (
	"fmt"
	"github.com/DipandaAser/deserve/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path/filepath"
)

var ProjectConfig *config.Configuration

func Upload(context *gin.Context){
	// We check the api key
	if ProjectConfig.UploadKey != context.GetHeader("UploadKey") {
		context.String(401, "Bad UploadKey.")
		return
	}

	// We check the filename
	fileName := context.GetHeader("filename")
	if fileName == "" {
		context.String(400, "Provide a file name.")
		return
	}

	// We read the body
	bufData, err := ioutil.ReadAll(context.Request.Body)
	if err != nil && err.Error() != "EOF" {
		context.String(406, "Bad content")
		return
	}

	// We store the file locally
	err = ioutil.WriteFile(filepath.Join(ProjectConfig.Folder, fileName), bufData, 0777)
	if err != nil {
		context.String(500, fmt.Sprintf("We were not able to store the file \n %s", err.Error()))
		return
	}

	context.String(200,"")
}