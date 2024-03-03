package filecache

import (
	"encoding/json"
	"errors"
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
			return nil, err
		}
		for _, file := range files {
			dcFile := models.DockerCompose{}
			file, err := os.ReadFile(file)
			if err != nil {
				continue
			}
			// Unmarshal the yaml file
			err = yaml.Unmarshal(file, &dcFile)
			if err != nil {
				continue
			}
			composeFiles = append(composeFiles, dcFile)
		}
		// marshal to json
		composeFilesJson, _ := json.Marshal(composeFiles)
		err = cache.SetKeyValue([]byte("fileCache_fromEmails"), composeFilesJson, 5*time.Minute)
		if err != nil {
			return nil, err
		}
	}
	return composeFiles, nil
}
