package utils

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rohanshrestha09/todo/configs"
	"github.com/rohanshrestha09/todo/enums"
)

func UploadFile(uploadedFile *multipart.FileHeader, fileDir enums.FileDIR, acceptedFileType enums.FileType) (string, string, error) {

	fileType := strings.Split(uploadedFile.Header.Get("Content-Type"), "/")

	if !strings.HasPrefix(fileType[0], string(acceptedFileType)) {
		return "", "", errors.New("please provide an image")
	}

	file, err := uploadedFile.Open()

	if err != nil {
		return "", "", err
	}

	defer file.Close()

	uuid := uuid.New().String()

	fileName := uuid + "." + fileType[1]

	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Second)
	defer cancel()

	object := configs.Bucket.Object(string(fileDir) + fileName)
	writer := object.NewWriter(ctx)

	//Set the attribute
	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": uuid}
	defer writer.Close()

	if _, err := io.Copy(writer, file); err != nil {
		return "", "", err
	}

	fileUrl := "https://firebasestorage.googleapis.com/v0/b/dev-synapse-0.appspot.com/o/" + url.QueryEscape(object.ObjectName()) + "?alt=media&token=" + uuid

	return fileUrl, fileName, nil
}
