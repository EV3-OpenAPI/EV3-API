package utils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"golang.org/x/mod/semver"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const RELEASE_URL = "https://api.github.com/repos/EV3-OpenAPI/EV3-API/releases/latest"
const VERSION_FILE_NAME = "version"
const SERVER_FILE_NAME = "ev3api-server"

var client = &http.Client{
	Transport: &http.Transport{
		TLSHandshakeTimeout: 20 * time.Second,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
		},
	},
	Timeout: 30 * time.Second,
}

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

// CheckForNewVersion checks if there is a new release, if there is, downloads it.
// After the download it starts the executable, and if that is successful it replaces
// the original executable and exits, causing systemd to restart this application.
func CheckForNewVersion() {
	log.Printf("INFO - Check for new versions")

	res, err := client.Get(RELEASE_URL)
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

	currVer := readCurrentVersion()
	if semver.Compare(currVer, rel.TagName) == -1 {
		log.Printf("INFO - Newer version found. Current: %s, New: %s", currVer, rel.TagName)

		downloadUrl, err := getAssetWithName(rel.Assets, "server")
		if err != nil {
			log.Printf("ERROR - No server binary in release. %v", err)
			return
		}
		err = downloadVersion(downloadUrl)
		if err != nil {
			log.Printf("ERROR - Update unsuccessful, continue with current version. %v", err)
			return
		}

		// Update version file
		err = writeCurrentVersion(rel.TagName)
		if err != nil {
			log.Printf("WARNING - new version string could not be written. %v", err)
		}

		log.Printf("INFO - Update successful, restarting")
		os.Exit(2) // exit with error code to cause systemd to restart the new executable
	} else {
		log.Printf("INFO - Already at newest version. %s", currVer)
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

func downloadVersion(url string) error {
	newFilePath := fmt.Sprintf("./%s.new", SERVER_FILE_NAME)
	oldFilePath := fmt.Sprintf("./%s", SERVER_FILE_NAME)

	// Create blank, executable file
	file, err := os.Create(newFilePath)
	if err != nil {
		return fmt.Errorf("cannot create new server binary file. %v", err)
	}

	// Change permission to be executable
	err = os.Chmod(newFilePath, 0744)
	if err != nil {
		return fmt.Errorf("cannot make new server binary file executable. %v", err)
	}

	// Download file content
	res, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("cannot create new server binary file. %v", err)
	}
	defer res.Body.Close()

	// Write download body to file
	size, err := io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("cannot write new server binary file. %v", err)
	}

	err = file.Close()
	if err != nil {
		return fmt.Errorf("could not close new executable file")
	}

	// Check if new file is executable
	if !isExecutable(newFilePath) {
		os.Remove(newFilePath)
		return fmt.Errorf("new server binary is not executable, aborting")
	}

	fmt.Printf("INFO - Downloaded a file %s with size %d", newFilePath, size)

	// Replace old executable with new one
	if err = os.Rename(newFilePath, oldFilePath); err != nil {
		return fmt.Errorf("")
	}

	return nil
}

func isExecutable(filePath string) bool {
	cmd := exec.Command(filePath, "-verify")

	if err := cmd.Start(); err != nil {
		log.Printf("ERROR - cmd.Start: %v", err)
		return false
	}

	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode() == 0
		}
	}

	return true
}

func readCurrentVersion() string {
	fileContent, err := os.ReadFile(VERSION_FILE_NAME)
	if err != nil {
		return "v0.0.0"
	}

	return string(fileContent)
}

func writeCurrentVersion(version string) error {
	file, err := os.Open(VERSION_FILE_NAME)
	if err != nil {
		file, err = os.Create(VERSION_FILE_NAME)
		if err != nil {
			return fmt.Errorf("could not open version file. %v", err)
		}
	}
	defer file.Close()

	_, err = file.WriteString(version)
	if err != nil {
		return fmt.Errorf("could write to version file. %v", err)
	}

	return nil
}
