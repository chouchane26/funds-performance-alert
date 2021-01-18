package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

func sendEmail(awsConfig *AWSSessionConfig, email Email) {
	sns1 := sns.New(awsConfig.Session)

	sns1.Publish(&sns.PublishInput{
		Subject: 	aws.String(email.Subject),
		Message:  	aws.String(email.Text),
		TopicArn: 	aws.String(GlobalConfig.AWS.SNS.ARN),
	})
}
