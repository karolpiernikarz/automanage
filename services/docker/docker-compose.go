package docker

import (
	"fmt"
	"github.com/spf13/viper"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

// UpdateAllApps WIP
func UpdateAllApps() (err error) {
	// get the docker-compose.yml files from /docker/web/
	// for each file, run docker-compose up -d
	files, err := FindFiles()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(files)
	return
}

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		println(s)
	}
	if d.Name() == "docker-compose.yaml" {
		println(s)
	}
	return nil
}

func FindFiles() (response []string, err error) {
	libRegEx, e := regexp.Compile(`docker-compose\.yaml`)
	if e != nil {
		return nil, e
	}
	e = filepath.Walk(viper.GetString("app.workdir"), func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			response = append(response, path)
		}
		return nil
	})
	return
}
