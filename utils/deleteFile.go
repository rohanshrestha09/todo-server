package utils

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	"github.com/rohanshrestha09/todo/configs"
)

func DeleteFile(fileName string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Second)
	defer cancel()

	object := configs.Bucket.Object(fileName)

	attrs, err := object.Attrs(ctx)
	if err != nil {
		return err
	}

	object = object.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if err := object.Delete(ctx); err != nil {
		return err
	}

	return nil
}
