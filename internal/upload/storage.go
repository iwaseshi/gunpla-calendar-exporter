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

func CloudStorage(targetFilePath string) error {

	ctx := context.Background()

	file, err := os.Open(targetFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bucketName := "gunpla-calendar-exporter"
	credentialsFile := "storage_key.json"

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return err
	}

	now := time.Now()
	objectPath := fmt.Sprintf("%d%s.ics", now.Year(), now.Month().String())
	obj := client.Bucket(bucketName).Object(objectPath)

	wc := obj.NewWriter(ctx)
	wc.ContentType = "text/calendar"
	if _, err := io.Copy(wc, file); err != nil {
		return err
	}
	defer wc.Close()

	acl := client.Bucket(bucketName).Object(objectPath).ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}

	fmt.Println("file upload success:" + fmt.Sprintf("https://storage.cloud.google.com/%s/%s", bucketName, objectPath))
	return nil
}
