package aws

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/tae2089/devops-platform-backend/domain"
)

// S3CloudFrontPolicy for cloudfront polict with bucket
type S3CloudFrontPolicy struct {
	BucketName             string
	OriginAccessIdentityID string
}

func (a *utilImpl) createBucket(bucketName string) error {
	_, err := a.s3Client.CreateBucket(context.Background(), &s3.CreateBucketInput{
		Bucket: &bucketName,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintApNortheast2,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *utilImpl) putBucketPublicAccess(bucketName string) error {
	_, err := a.s3Client.PutPublicAccessBlock(context.Background(), &s3.PutPublicAccessBlockInput{
		Bucket: &bucketName,
		PublicAccessBlockConfiguration: &types.PublicAccessBlockConfiguration{
			BlockPublicAcls:       true,
			BlockPublicPolicy:     true,
			IgnorePublicAcls:      true,
			RestrictPublicBuckets: true,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *utilImpl) putBucketPolicy(bucketName, policy string) error {
	_, err := a.s3Client.PutBucketPolicy(context.Background(), &s3.PutBucketPolicyInput{
		Bucket: &bucketName,
		Policy: &policy,
	})
	if err != nil {
		return err
	}
	return nil
}

// BucketExists checks whether a bucket exists in the current account.
func (a *utilImpl) bucketExists(bucketName string) (bool, error) {
	_, err := a.s3Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	exists := true
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.(type) {
			case *types.NotFound:
				log.Printf("Bucket %v is available.\n", bucketName)
				exists = false
				err = nil
			default:
				log.Printf("Either you don't have access to bucket %v or another error occurred. "+
					"Here's what happened: %v\n", bucketName, err)
			}
		}
	} else {
		log.Printf("Bucket %v exists and you already own it.", bucketName)
	}

	return exists, err
}

func (a *utilImpl) createS3PublicGetPolicy(bucketName string) (string, error) {
	readOnlyAnonUserPolicy := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Sid":       "AddPerm",
				"Effect":    "Allow",
				"Principal": "*",
				"Action": []string{
					"s3:GetObject",
				},
				"Resource": []string{
					fmt.Sprintf("arn:aws:s3:::%s/*", bucketName),
				},
			},
		},
	}
	policy, err := json.Marshal(readOnlyAnonUserPolicy)
	if err != nil {
		return "", err
	}
	return string(policy), nil
}

func (a *utilImpl) createS3OnlyCloudFrontAccessPolicy(bucketName, originAccessIdentityID string) (string, error) {

	policy := S3CloudFrontPolicy{
		BucketName:             bucketName,
		OriginAccessIdentityID: originAccessIdentityID,
	}

	tmpl, err := template.New("Create Cloud front bucket policy").Parse(domain.CloudFrontBucketPolicy)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, policy); err != nil {
		return "", err
	}
	fmt.Println(tpl.String())
	return tpl.String(), nil
}
