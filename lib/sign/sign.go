package sign

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

// URL...
func URL(key string) ([]byte, error) {
	accessKeyID := "AKIAIP7LDPSLWSEZIWQQ"
	secretAccessKey := "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"

	sess := session.New(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("zoocasa.importer-go"),
		Key:    aws.String(key),
	})

	urlStr, err := req.Presign(2 * time.Minute)
	if err != nil {
		return nil, errors.Wrap(err, "error presigning")
	}

	return []byte(urlStr), nil
}
