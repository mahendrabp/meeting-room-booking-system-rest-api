package gcsbucket

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"io"
	"net/url"
	"os"
)

var (
	storageClient *storage.Client
)

// HandleFileUploadToBucket uploads file to bucket
func HandleFileUploadToBucket(c *gin.Context) (string, error) {
	var err error

	//err = godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error getting env, %v", err)
	//} else {
	//	fmt.Println("We are getting values")
	//}

	bucket := os.Getenv("BUCKET_NAME")

	ctx := appengine.NewContext(c.Request)

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return "", err
	}

	f, uploadedFile, err := c.Request.FormFile("photo")
	if err != nil {
		return "", nil
	}

	defer f.Close()

	sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", nil
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return "", nil
	}

	return u.EscapedPath(), nil

}
