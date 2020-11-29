package archive

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
	"path/filepath"
)

type BucketClient struct {
	session  *session.Session
	Uploader *s3manager.Uploader

	bucket string
	key    string

	CallbackUrl string
}

var instance *BucketClient

func GetBucketInstance() *BucketClient {
	if instance == nil {
		instance = new(BucketClient)
	}
	return instance
}

func (c *BucketClient) Connect(bucket, key, secret string) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	})
	if err != nil {
		panic(err)
	} else {
		log.Println("S3 uploader session created")
	}

	c.session = sess
	uploader := s3manager.NewUploader(sess)
	c.Uploader = uploader
	c.bucket = bucket
	c.key = key
}

func (c *BucketClient) Upload(path string) bool {
	log.Println(fmt.Sprintf("Starting upload file %s", path))

	file, err := os.Open(path)
	if err != nil {
		log.Println("Uploader cannot open file")
		return false
	}
	defer file.Close()

	basename := filepath.Base(path)
	uploadName := "videos/" + basename
	fmt.Println(uploadName)
	_, err = c.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(uploadName),
		Body:   file,
	})

	if err == nil {
		log.Println(fmt.Sprintf("Successfully uploaded file %s", path))
		return true
	} else {
		log.Println(err)
		return false
	}

}
