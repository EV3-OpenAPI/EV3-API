package utils

import (
	"encoding/json"
	"fmt"
	"golang.org/x/mod/semver"
	"io"
	"log"
	"net/http"
	"os"
)

const RELEASE_URL = "https://api.github.com/repos/EV3-OpenAPI/EV3-API/releases/latest"

type release struct {
	Id              int      `json:"id"`
	TagName         string   `json:"tag_name"`
	TargetCommitish string   `json:"target_commitish"`
	Name            string   `json:"name"`
	Draft           bool     `json:"draft"`
	PreRelease      bool     `json:"pre_release"`
	Assets          *[]asset `json:"assets"`
}

type asset struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	ContentType        string `json:"content_type"`
	Size               int    `json:"size"`
	BrowserDownloadUrl string `json:"browser_download_url"`
	State              string `json:"state"`
}

func CheckForNewVersion() {
	res, err := http.Get(RELEASE_URL)
	if err != nil {
		log.Printf("ERROR - Failed to fetch newest version, continung with existing. %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("ERROR - Failed to fetch newest version, continung with existing. %v", err)
		return
	}

	rel := release{}
	err = json.Unmarshal(body, &rel)
	if err != nil {
		log.Printf("ERROR - Failed to fetch newest version, continung with existing. %v", err)
		return
	}

	currVer := getCurrentVersion()
	if semver.Compare(currVer, rel.TagName) == -1 {
		log.Printf("INFO - Newer version found. Current: %s, New: %s", currVer, rel.TagName)

		downloadUrl, err := getAssetWithName(rel.Assets, "server")
		if err != nil {
			log.Printf("ERROR - No server binary in release")
			return
		}
		downloadVersion(downloadUrl)
	}
}

func getAssetWithName(assets *[]asset, name string) (string, error) {
	for _, a := range *assets {
		if a.Name == name {
			return a.BrowserDownloadUrl, nil
		}
	}

	return "", fmt.Errorf("no asset with the given name found")
}

func downloadVersion(url string) {
	// Create blank, executable file
	file, err := os.Create("ev3api-server.new")
	os.Chmod(file.Name(), 0744)
	if err != nil {
		log.Printf("ERROR - cannot create new server binary file. %v", err)
		return
	}

	// Put content on file
	res, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR - cannot create new server binary file. %v", err)
		return
	}
	defer res.Body.Close()

	size, err := io.Copy(file, res.Body)
	if err != nil {
		log.Printf("ERROR - cannot write new server binary file. %v", err)
	}
	defer file.Close()

	if !isExecutable(file) {
		log.Printf("ERROR - new server binary is not executable, aborting")
		os.Remove(file.Name())
	}

	fmt.Printf("Downloaded a file %s with size %d", file.Name(), size)
	// TODO: replace current server file with new one
}

func isExecutable(file *os.File) bool {
	stat, err := file.Stat()
	log.Printf("stat: %v, err: %v", stat, err)
	return false
}

func getCurrentVersion() string {
	file, err := os.ReadFile("version")
	if err != nil {
		return "0.0.0"
	}

	return string(file)
}
