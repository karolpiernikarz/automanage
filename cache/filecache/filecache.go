package filecache

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/karolpiernikarz/automanage/cache"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/karolpiernikarz/automanage/services/docker"
	"gopkg.in/yaml.v3"
)

func DockerComposeFiles() (composeFiles []models.DockerCompose, err error) {
	// check if key exist
	var value string
	if value, err = cache.GetValueFromKey("fileCache_fromEmails"); err == nil {
		// if key exist, return value
		err = json.Unmarshal([]byte(value), &composeFiles)
		return
	}
	if errors.Is(badger.ErrKeyNotFound, err) {
		files, err := docker.FindFiles()
		if err != nil {
			fmt.Println("Error finding Docker files:", err)
			return nil, err
		}
		for _, filePath := range files {
			dcFile := models.DockerCompose{}
			file, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", filePath, err)
				continue
			}
			// Unmarshal the yaml file
			err = yaml.Unmarshal(file, &dcFile)
			if err != nil {
				fmt.Println("Error unmarshaling file:", filePath, err)
				continue
			}
			composeFiles = append(composeFiles, dcFile)
		}
		// marshal to json
		composeFilesJson, _ := json.Marshal(composeFiles)
		err = cache.SetKeyValue([]byte("fileCache_fromEmails"), composeFilesJson, 5*time.Minute)
		if err != nil {
			fmt.Println("Error setting cache value:", err)
			return nil, err
		}
	} else {
		fmt.Println("Error retrieving cache value:", err)
	}
	return composeFiles, nil
}
