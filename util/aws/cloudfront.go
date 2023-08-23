package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	awshttp "github.com/aws/smithy-go/transport/http"
)

func (a *utilImpl) createDistribution(bucketName, domain, certificateSSLArn string) (*cloudfront.CreateDistributionOutput, string, error) {
	bucket, err := a.s3Client.HeadBucket(context.Background(), &s3.HeadBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		return nil, "", err
	}
	bucketResponseMetadata := middleware.GetRawResponse(bucket.ResultMetadata).(*awshttp.Response)

	region := bucketResponseMetadata.Header.Get("x-amz-bucket-region")
	originDomain := bucketName + ".s3." + region + ".amazonaws.com"
	oaiID, err := a.createoriginAccessIdentity(originDomain)
	// fmt.Println(oaiID)
	if err != nil {
		return nil, "", err
	}

	cloudfrontResponse, err := a.cloudfrontClient.CreateDistribution(context.Background(), &cloudfront.CreateDistributionInput{
		DistributionConfig: &cloudfrontTypes.DistributionConfig{
			Enabled:         aws.Bool(true),
			CallerReference: &originDomain,
			Comment:         &originDomain,
			//IPV6 지원 여부
			IsIPV6Enabled: aws.Bool(false),
			//가격 분류
			PriceClass: cloudfrontTypes.PriceClassPriceClass100,
			//지원되는 HTTP 버전
			HttpVersion: cloudfrontTypes.HttpVersionHttp11,
			//기본값 루트 객체
			DefaultRootObject: aws.String("index.html"),
			// 대체 도메인
			Aliases: &cloudfrontTypes.Aliases{
				Quantity: aws.Int32(1),
				Items:    []string{domain},
			},
			// 사용자 SSL 인증서
			ViewerCertificate: &cloudfrontTypes.ViewerCertificate{
				ACMCertificateArn: aws.String(certificateSSLArn),
				SSLSupportMethod:  cloudfrontTypes.SSLSupportMethodSniOnly,
			},

			CustomErrorResponses: &cloudfrontTypes.CustomErrorResponses{
				Quantity: aws.Int32(1),
				Items: []cloudfrontTypes.CustomErrorResponse{
					{
						ErrorCode:          aws.Int32(403),
						ResponseCode:       aws.String("200"),
						ErrorCachingMinTTL: aws.Int64(10),
						ResponsePagePath:   aws.String("/index.html"),
					},
				},
			},

			//원본
			Origins: &cloudfrontTypes.Origins{
				Quantity: aws.Int32(1),
				Items: []cloudfrontTypes.Origin{
					{
						DomainName: &originDomain,
						Id:         &originDomain,
						S3OriginConfig: &cloudfrontTypes.S3OriginConfig{
							OriginAccessIdentity: aws.String("origin-access-identity/cloudfront/" + oaiID),
						},
					},
				},
			},
			CacheBehaviors: nil,
			// 기본 캐시 동작
			DefaultCacheBehavior: &cloudfrontTypes.DefaultCacheBehavior{
				TargetOriginId:       aws.String(originDomain),
				Compress:             aws.Bool(true),
				ViewerProtocolPolicy: cloudfrontTypes.ViewerProtocolPolicyRedirectToHttps,
				CachePolicyId:        aws.String(CachingOptimized.String()),
			},
		},
	},
	)
	if err != nil {
		return nil, "", err
	}

	return cloudfrontResponse, oaiID, err
}

func (a *utilImpl) createoriginAccessIdentity(domainName string) (string, error) {
	ctx := context.Background()
	oai, err := a.cloudfrontClient.CreateCloudFrontOriginAccessIdentity(ctx, &cloudfront.CreateCloudFrontOriginAccessIdentityInput{
		CloudFrontOriginAccessIdentityConfig: &cloudfrontTypes.CloudFrontOriginAccessIdentityConfig{
			CallerReference: aws.String(domainName),
			Comment:         aws.String(domainName),
		},
	})
	if err != nil {
		return "", err
	}
	return *oai.CloudFrontOriginAccessIdentity.Id, nil
}
