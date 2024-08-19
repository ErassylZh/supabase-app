package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Storage interface {
	CreateImage(ctx context.Context, bucketName string, fileName, fileUrl string) (string, error)
}

type StorageClient struct {
	supabaseURL string
	supabaseKey string
}

func NewStorageClient(supabaseURL string, supabaseKey string) StorageClient {
	return StorageClient{supabaseURL: supabaseURL, supabaseKey: supabaseKey}
}

func (s *StorageClient) CreateImage(ctx context.Context, bucketName string, fileName, fileUrl string) (string, error) {
	fileName = strings.Replace(fileName, "·", "_", len(fileName))
	fileName = time.Now().String() + " " + fileName
	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", s.supabaseURL, bucketName, fileName)
	file, err := downloadFileInMemory(fileUrl)
	if err != nil {
		return "", err
	}

	body := bytes.NewBuffer(file)
	req, err := http.NewRequest(http.MethodPost, uploadURL, body)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	contentType := http.DetectContentType(file)

	req.Header.Set("Authorization", "Bearer "+s.supabaseKey)
	req.Header.Set("Content-Type", contentType)
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		responseBody, _ := io.ReadAll(resp.Body)
		fmt.Printf("Response Body: %s\n %s", string(responseBody), fileName)

		return "", fmt.Errorf("failed to upload file, status: %s", resp.Status)
	}

	var uploadResponse UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResponse); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return uploadResponse.Key, nil
}

func downloadFileInMemory(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type UploadResponse struct {
	Key string `json:"Key"`
}
