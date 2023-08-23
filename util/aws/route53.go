package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

func (a *utilImpl) createRecordSet(hostedZoneID, domain, cloudfrontDomain string) error {
	_, err := a.routeClient.ChangeResourceRecordSets(
		context.Background(),
		&route53.ChangeResourceRecordSetsInput{
			HostedZoneId: aws.String(hostedZoneID),
			ChangeBatch: &types.ChangeBatch{
				Changes: []types.Change{
					{
						Action: types.ChangeActionUpsert,
						ResourceRecordSet: &types.ResourceRecordSet{
							Type: types.RRTypeA,
							Name: aws.String(domain),
							AliasTarget: &types.AliasTarget{
								DNSName:              aws.String(cloudfrontDomain),
								HostedZoneId:         aws.String("Z2FDTNDATAQYW2"), // cloudfront의 기본 hostedZone ID 는 Z2FDTNDATAQYW2
								EvaluateTargetHealth: false,
							},
						},
					},
				},
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
