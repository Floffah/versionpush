package fetch

import (
	"os"
	"path/filepath"
	"strings"
	"time"
	"versionpush/src/util"
)

func findLatestJar(ver string) string {
	dir, err := os.Getwd()
	if err != nil || dir == "" {
		util.Fatal("Error while finding current working directory.")
		panic(err)
	}
	root := filepath.Join(dir, "target")
	latest := ""
	var latestTime time.Time

	listerr := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".jar") {
			if latest == "" {
				latest = path
				latestTime = info.ModTime()
			} else if info.ModTime().After(latestTime) {
				latest = path
				latestTime = info.ModTime()
			}
		}
		return nil
	})
	if listerr != nil {
		util.Fatal("Error while finding files in a directory")
		panic(err)
	}
	return latest
}