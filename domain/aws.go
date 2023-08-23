package domain

const CloudFrontBucketPolicy = `
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "2",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity {{.OriginAccessIdentityID}}"
            },
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::{{.BucketName}}/*"
        }
    ]
}
`
