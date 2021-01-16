package main

import (
	"encoding/json"
	"github.com/DipandaAser/deserve/uploader"
	"github.com/DipandaAser/deserve/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
)



func main() {

	var ProjectConfig config.Configuration
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Missing configuration file.")
	}

	err = json.Unmarshal(b, &ProjectConfig)
	if err != nil {
		log.Fatal("Somethings went wrong with the config file.")
	}


	// Check if the path provide exist and is a directory
	if fileFolderInfo, err := os.Stat(ProjectConfig.Folder) ; os.IsNotExist(err)  {
		log.Fatal("Folder does not exist. Please provide an existing folder")
	}else if !fileFolderInfo.IsDir(){
		log.Fatal("Please provide a folder")
	}

	uploader.ProjectConfig = &ProjectConfig

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.Static("/download", ProjectConfig.Folder)
	router.POST("/upload", uploader.Upload)


	if err = router.Run(":" + ProjectConfig.Port); err != nil {
		log.Fatalf("Server shutdown. Caused by : %v ", err)
	}
}
