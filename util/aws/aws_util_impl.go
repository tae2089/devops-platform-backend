package aws

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var _ Util = (*utilImpl)(nil)

type utilImpl struct {
	cloudfrontClient *cloudfront.Client
	s3Client         *s3.Client
	routeClient      *route53.Client
	ecrClient        *ecr.Client
	acmClient        *acm.Client
}

// CreateDistribution implements AwsService
func (a *utilImpl) CreateDistribution(bucketName string, domain string, certificateSSLArn string) (*types.Distribution, error) {
	output, oaiID, err := a.createDistribution(bucketName, domain, certificateSSLArn)
	if err != nil {
		return nil, err
	}
	policy, err := a.createS3OnlyCloudFrontAccessPolicy(bucketName, oaiID)
	if err != nil {
		return nil, err
	}
	time.Sleep(15 * time.Second)
	err = a.putBucketPolicy(bucketName, policy)
	if err != nil {
		return nil, err
	}

	return output.Distribution, nil
}

func (a *utilImpl) CreateBucket(bucketName string, isPublish bool) (string, error) {
	if bucketName == "" {
		return "", errors.New("Bucket name cannot be empty")
	}
	err := a.createBucket(bucketName)

	if err != nil {
		return "", err
	}

	if isPublish {
		policy, err := a.createS3PublicGetPolicy(bucketName)
		if err != nil {
			return "", err
		}
		err = a.putBucketPolicy(bucketName, policy)
	} else {
		err = a.putBucketPublicAccess(bucketName)
	}

	if err != nil {
		return "", err
	}
	return "create succes bucket", nil
}

func (a *utilImpl) CreateRouterRecordSet(hostedZoneID, domain, cloudfrontDomain string) error {
	err := a.createRecordSet(hostedZoneID, domain, cloudfrontDomain)
	if err != nil {
		return err
	}
	return nil
}

func (a *utilImpl) CreateContainerRegistry(repositoryName string) (string, error) {
	result, err := a.createContainerRegistry(repositoryName)
	if err != nil {
		return "", err
	}
	return *result.Repository.RepositoryUri, nil
}
