package integration

import (
	"bytes"
	"context"
	storage "github.com/supabase-community/storage-go"
	"io"
	"net/http"
)

type StorageClient interface {
	CreateImage(ctx context.Context, bucketName string, fileUrl string) (string, error)
}

type StorageClientService struct {
	storageClient *storage.Client
}

func NewStorageClientService(supabaseUrl, supabaseKey string) *StorageClientService {
	return &StorageClientService{storageClient: storage.NewClient(supabaseUrl, supabaseKey, nil)}
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
