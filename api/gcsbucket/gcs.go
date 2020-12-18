package gcsbucket

import (
	"cloud.google.com/go/storage"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"io"
	"net/url"
	"os"
	"path/filepath"
)

var (
	storageClient *storage.Client
)

// HandleFileUploadToBucket uploads file to bucket
func HandleFileUploadToBucket(c *gin.Context) (string, error) {
	var err error

	bucket := os.Getenv("BUCKET_NAME")

	ctx := appengine.NewContext(c.Request)

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("google-credentials.json"))
	if err != nil {
		return "", err
	}

	f, uploadedFile, err := c.Request.FormFile("photo")
	if err != nil {
		return "", err
	}

	fileExtension := filepath.Ext(uploadedFile.Filename)

	if fileExtension != ".png" && fileExtension != ".jpg" && fileExtension != ".jpeg" {
		return "", errors.New("silahkan upload dengan ekstensi png/jpg/jpeg")
	}

	defer f.Close()

	sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return "", err
	}

	return u.EscapedPath(), nil

}
