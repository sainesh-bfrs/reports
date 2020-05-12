package aws

/*
 * File: aws.go
 * File Created: Tuesday, 12th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"reports/config"
	"reports/helpers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Upload file to S3
func Upload(filepath string, s3filepath string) string {
	file, err := os.Open(filepath)
	helpers.LogError("Unable to open file", err)
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size) // read file content to buffer

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := s3filepath

	creds := credentials.NewStaticCredentials(config.Config.AwsKey, config.Config.AwsSecret, "")
	_, err = creds.Get()
	helpers.LogError("Unable to get credentials ", err)
	cfg := aws.NewConfig().WithRegion(config.Config.AwsRegion).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	params := &s3.PutObjectInput{
		Bucket:               aws.String(config.Config.AwsBucket),
		Key:                  aws.String(path),
		Body:                 fileBytes,
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(fileType),
		ACL:                  aws.String("public-read"),
		ServerSideEncryption: aws.String("AES256"),
	}

	_, err = svc.PutObject(params)

	helpers.LogError("Unable to put object: ", err)

	return fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", config.Config.AwsBucket, config.Config.AwsRegion, path)
}
