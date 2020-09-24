package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"versionpush/src/fetch"
	"versionpush/src/util"
)

func main() {
	util.Info("VersionPush v1.0.0 by Floffah\n---")
	var lang, pm, ver string
	flag.StringVar(&pm, "pm", "maven", "Project manager")
	flag.StringVar(&lang, "lang", "java", "Language profile")
	flag.StringVar(&ver, "ver", "1.0.0", "Version of prebuilt")

	flag.Parse()

	_ = os.Mkdir(path.Join(files.CWD(), ".versionpush"), os.ModeDir)

	if lang == "java" {
		util.Info("Using language profile Java")
		java(pm, ver)
	}
}

func java(pm string, ver string) {
	if pm == "maven" {
		util.Info("Using java project manager Maven")
		mavens := mavenPrebuilt(ver)
		final,archive := files.GetFinal(mavens)
		fmt.Println(final,archive)
	} else {
		util.Fatal("Unknown builder type \"" + pm + "\"")
		os.Exit(1)
	}
}

func mavenPrebuilt(ver string) []string {
	var resources []string
	latestJar := files.FindLatestJar(ver, files.CWD())

	if latestJar != nil {
		resources = append(resources, *latestJar)
		return resources
	} else {
		util.Fatal("No maven pre-built found.")
		os.Exit(1)
		return resources
	}
}