package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

func (a *utilImpl) createContainerRegistry(repositoryName string) (*ecr.CreateRepositoryOutput, error) {
	result, err := a.ecrClient.CreateRepository(context.Background(), &ecr.CreateRepositoryInput{
		RepositoryName:     aws.String(repositoryName),
		ImageTagMutability: types.ImageTagMutabilityImmutable,
		ImageScanningConfiguration: &types.ImageScanningConfiguration{
			ScanOnPush: true,
		},
		Tags: []types.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String(repositoryName),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
