// Package aws is for Aws Services
package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Service is the interface for AWS services
type Util interface {
	CreateDistribution(bucketName, domain, certificateSSLArn string) (*cloudfrontTypes.Distribution, error)
	CreateBucket(bucketName string, isPublish bool) (string, error)
	CreateRouterRecordSet(hostedZoneID, domain, cloudfrontDomain string) error
	CreateContainerRegistry(repositoryName string) (string, error)
}

// NewAwsUtil creates a new AwsUtil
func NewAwsUtil(profile string) Util {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(profile))
	if err != nil {
		return nil
	}

	s3Client := s3.NewFromConfig(cfg)
	routeClient := route53.NewFromConfig(cfg)
	cloudfrontClient := cloudfront.NewFromConfig(cfg)
	acmClient := acm.NewFromConfig(cfg)
	ecrClient := ecr.NewFromConfig(cfg)
	return &utilImpl{cloudfrontClient, s3Client, routeClient, ecrClient, acmClient}
}
