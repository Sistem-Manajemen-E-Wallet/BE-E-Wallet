package upload

import (
	"context"
	"e-wallet/app/configs"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func ImageUploadHelper(input multipart.File) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(configs.CLOUDINARY_CLOUD_NAME, configs.CLOUDINARY_API_KEY, configs.CLOUDINARY_API_SECRET)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: configs.CLOUDINARY_UPLOAD_FOLDER})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
