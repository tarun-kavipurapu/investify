package aws

import (
	"context"
	"fmt"
	"investify/config"
	"investify/util"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

type S3Service struct {
	Client *s3.Client
}

func NewS3Service(ctx context.Context) *S3Service {
	cfg, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(config.EnvVars.AWS_REGION),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(config.EnvVars.AWS_ACCESS_TOKEN, config.EnvVars.AWS_SECRET_TOKEN_KEY, ""),
		),
	)
	if err != nil {
		log.Fatalf("Unable to load AWS SDK config: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return &S3Service{Client: client}
}
func (awsSvc *S3Service) UploadFileAndDeleteTemp(ctx *gin.Context, file multipart.File, header *multipart.FileHeader) (string, error) {

	tempUploadDir := "./assets/uploads/temp"
	if _, err := os.Stat(tempUploadDir); os.IsNotExist(err) {
		err = os.MkdirAll(tempUploadDir, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("unable to create temp upload directory: %w", err)
		}
	}

	// Save the file to the temp directory
	filePath := filepath.Join(tempUploadDir, header.Filename)
	if err := ctx.SaveUploadedFile(header, filePath); err != nil {
		return "", fmt.Errorf("unable to save the file: %w", err)
	}
	encryptedFileName := util.GenerateUniqueFilename(header.Filename)
	url, err := awsSvc.UploadFile(encryptedFileName, filePath)
	if err != nil {
		return "", err
	}

	if err := os.Remove(filePath); err != nil {
		log.Printf("failed to remove temp file: %v", err)
	}
	return url, nil

}

func (awsSvc *S3Service) UploadFile(bucketKey string, fileName string) (string, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	log.Println("after file open")

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	contentLength := fileInfo.Size()
	log.Println("after file Info")

	_, err = awsSvc.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(config.EnvVars.AWS_BUCKET_NAME),
		Key:           aws.String(bucketKey),
		Body:          file,
		ContentLength: &contentLength,
		ContentType:   aws.String("image/jpeg"),
	})
	if err != nil {
		return "", err
	}

	log.Println("after upload Info")

	// Construct the URL for the uploaded image
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", config.EnvVars.AWS_BUCKET_NAME, config.EnvVars.AWS_REGION, bucketKey)
	return url, nil
}
func (awsSvc *S3Service) DeleteImage(objectKey string) error {
	// Check if the object exists
	// headInput := &s3.HeadObjectInput{
	// 	Bucket: aws.String(config.EnvVars.AWS_BUCKET_NAME),
	// 	Key:    aws.String(objectKey),
	// }

	// _, err := awsSvc.Client.HeadObject(context.TODO(), headInput)
	// if err != nil {
	// 	if aerr, ok := err.(*smithy.GenericAPIError); ok && aerr.Code() == "NotFound" {
	// 		log.Printf("Object %s not found in bucket %s", objectKey, config.EnvVars.AWS_BUCKET_NAME)
	// 		return nil
	// 	}
	// 	return fmt.Errorf("failed to check if object %s exists in bucket %s: %w", objectKey, config.EnvVars.AWS_BUCKET_NAME, err)
	// }

	// If the object exists, delete it
	deleteInput := &s3.DeleteObjectInput{
		Bucket: aws.String(config.EnvVars.AWS_BUCKET_NAME),
		Key:    aws.String(objectKey),
	}

	_, err := awsSvc.Client.DeleteObject(context.TODO(), deleteInput)
	if err != nil {
		return fmt.Errorf("failed to delete object %s from bucket %s: %w", objectKey, config.EnvVars.AWS_BUCKET_NAME, err)
	}

	log.Printf("Successfully deleted object %s from bucket %s", objectKey, config.EnvVars.AWS_BUCKET_NAME)
	return nil
}
