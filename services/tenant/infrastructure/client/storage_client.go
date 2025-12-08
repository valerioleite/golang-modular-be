package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"services/tenant/infrastructure/client/dto"
	"time"
)

const (
	uploadTimeout = 3 * time.Second
)

type StorageClient struct {
	baseURL string
	client  *http.Client
}

func NewStorageClient() *StorageClient {
	baseURL := os.Getenv("STORAGE_BASE_URL")
	return &StorageClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *StorageClient) UploadFile(ctx context.Context, bucket, filename string, file multipart.File) (*dto.UploadResponse, error) {
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	if err := writer.WriteField("bucket", bucket); err != nil {
		return nil, fmt.Errorf("failed to write bucket field: %w", err)
	}

	if err := writer.WriteField("filename", filename); err != nil {
		return nil, fmt.Errorf("failed to write filename field: %w", err)
	}

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("failed to copy file: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, uploadTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"storage", &requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("storage service returned status %d: %s", resp.StatusCode, string(body))
	}

	var uploadResp dto.UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &uploadResp, nil
}
