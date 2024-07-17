package media

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofrs/uuid"
	"github.com/the-Jinxist/tukio-api/config"
)

type ImageMeta struct {
	AssociatedID uuid.UUID
	ImageType    string
}

func (i *ImageMeta) getPublicID() string {
	return fmt.Sprintf("%s_%s", i.AssociatedID.String(), i.ImageType)
}

func UploadImage(ctx context.Context, imageFile string, meta ImageMeta) (string, error) {
	cfg := config.GetCurrentConfig()

	//should use cloudinary
	cld, err := cloudinary.NewFromParams("eventsly", cfg.CloudinaryAPIKey, cfg.CloudinarySecret)
	if err != nil {
		return "", nil
	}
	res, err := cld.Upload.Upload(ctx, imageFile, uploader.UploadParams{PublicID: meta.getPublicID()})
	if err != nil {
		return "", err
	}

	return res.URL, nil

}
