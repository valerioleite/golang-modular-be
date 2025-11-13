package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"io"
	"log/slog"
	"os"
	"services/storage/internal/storage/domain"
	"services/storage/internal/storage/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageRepositoryS3 struct {
	uploader *manager.Uploader
}

func NewStorageRepositoryS3() repository.StorageRepository {
	return &StorageRepositoryS3{}
}

func (r *StorageRepositoryS3) Init() error {
	client, err := r.setupS3Client()
	if err != nil {
		return err
	}

	r.uploader = manager.NewUploader(client)
	return nil
}

func (r *StorageRepositoryS3) Upload(ctx context.Context, storage *domain.Storage, file io.Reader) error {
	_, err := r.uploader.Upload(ctx, &awsS3.PutObjectInput{
		Bucket: aws.String(storage.Bucket),
		Key:    aws.String(storage.Filename),
		Body:   file,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *StorageRepositoryS3) setupS3Client() (*awsS3.Client, error) {
	region := os.Getenv("AWS_REGION")
	endpoint := os.Getenv("AWS_ENDPOINT")

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
		//config.WithClientLogMode(aws.LogRetries|aws.LogRequest),
	)

	if err != nil {
		return nil, err
	}

	var client = awsS3.NewFromConfig(cfg, func(o *awsS3.Options) {
		o.UsePathStyle = true
		if endpoint != "" {
			o.BaseEndpoint = aws.String(endpoint)
		}
	})

	slog.Info("AWS setup loaded successfully.")
	return client, nil
}
