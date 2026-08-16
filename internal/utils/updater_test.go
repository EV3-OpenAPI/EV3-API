package utils

import "testing"

func TestGetServerAssetFindsCurrentServerBinary(t *testing.T) {
	assets := []asset{
		{Name: "server", BrowserDownloadUrl: "https://example.com/legacy"},
		{Name: SERVER_FILE_NAME, BrowserDownloadUrl: "https://example.com/current"},
	}

	url, err := getServerAsset(&assets)
	if err != nil {
		t.Fatalf("getServerAsset() error = %v", err)
	}
	if url != "https://example.com/current" {
		t.Fatalf("getServerAsset() = %q, want current server asset", url)
	}
}
