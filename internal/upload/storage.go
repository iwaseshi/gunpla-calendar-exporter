package upload

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func CloudStorage(ctx context.Context, targetFilePath string) error {

	file, err := os.Open(targetFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bucketName := "gunpla-calendar-exporter"
	now := time.Now()
	objectPath := fmt.Sprintf("%d%s.ics", now.Year(), now.AddDate(0, 0, 20).Month().String())

	storage, err := newCloudStorage(ctx, bucketName)
	if err != nil {
		return err
	}
	if err = storage.WriteObject(file, objectPath); err != nil {
		return err
	}
	fmt.Println("file upload success:" + fmt.Sprintf("https://storage.cloud.google.com/%s/%s", bucketName, objectPath))
	return nil
}

type cloudStorage struct {
	bucket string
	client *storage.Client
	ctx    context.Context
}

func newCloudStorage(ctx context.Context, bucket string) (*cloudStorage, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("storage_key.json"))
	if err != nil {
		return nil, err
	}
	return &cloudStorage{
		bucket: bucket,
		client: client,
		ctx:    ctx,
	}, nil
}

func (st cloudStorage) WriteObject(file *os.File, objectPath string) error {
	obj := st.client.Bucket(st.bucket).Object(objectPath)
	wc := obj.NewWriter(st.ctx)
	wc.ContentType = "text/calendar"
	if _, err := io.Copy(wc, file); err != nil {
		return err
	}
	defer wc.Close()
	return nil
}
