package media

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(ctx context.Context, imageFile string) error {
	//should use cloudinary
	cld, _ := cloudinary.NewFromParams("n07t21i7", "123456789012345", "abcdeghijklmnopqrstuvwxyz12")
	_, err := cld.Upload.Upload(ctx, "my_picture.jpg", uploader.UploadParams{PublicID: "my_image"})
	return err

}
