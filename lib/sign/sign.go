package sign

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

// URL...
func URL(key string) ([]byte, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_ID"), os.Getenv("ACCESS_KEY"), ""),
	})
	if err != nil {
		return nil, errors.Wrap(err, "error logging in")
	}
	svc := s3.New(sess)

	req, res := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("zoocasa.importer-go"),
		Key:    aws.String(key),
	})

	if res.ContentLength == nil {
		return nil, errors.New("no file found with provided key")
	}

	urlStr, err := req.Presign(2 * time.Minute)
	if err != nil {
		return nil, errors.Wrap(err, "error presigning")
	}

	return []byte(urlStr), nil
}
