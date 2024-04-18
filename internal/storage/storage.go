package storage

import (
	"bytes"
	"context"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

// Uplaod bytes
func UploadBytes(data []byte, bucket, path string) error {

	//create buffer
	buffer := bytes.NewBuffer(data)

	//Create client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	//Create Writer
	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()
	writer := client.Bucket(bucket).Object(path).NewWriter(ctx)

	//Copy Data from Buffer to Google Cloud Storage
	_, err = io.Copy(writer, buffer)
	if err != nil {
		return err
	}

	//Close Writer
	err = writer.Close()
	if err != nil {
		return err
	}

	//Return Result
	return nil

}
