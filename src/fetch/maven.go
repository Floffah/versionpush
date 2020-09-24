package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"versionpush/src/util"
)

func FindLatestJar(ver string, basepath string) *string {
	root := filepath.Join(basepath, "target")
	latest := ""
	var latestTime time.Time

	listerr := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".jar") && strings.Contains(path, ver) {
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
		panic(listerr)
	}
	if latest != "" {
		util.Info(fmt.Sprintf("Found pre-built \"%v\"", latest))
		return &latest
	}
	return nil
}