package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Dynamodb struct {
	DynamodbSession *dynamodb.DynamoDB
	TableName string
}

type AWSSessionConfig struct {
	Session *session.Session
	Dynamodb Dynamodb
}

type DBConfig struct {
	Table string
}

type SNSConfig struct {
	ARN string
}

type AWSConfig struct {
	Region   string
	Dynamodb DBConfig
	SNS      SNSConfig
}

type AlphavantageConfig struct {
	Endpoint 	string
	Apikey 		string
}

type Config struct {
	Alphavantage 	AlphavantageConfig
	AWS 			AWSConfig
}

var GlobalConfig *Config

func (c *Config) readFromFile(filename string) {

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Cannot read configuration file: #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Fatalf("Cannot parse yml configuration file: %v", err)
	}
}

func decryptSensitiveData(c *Config, session *session.Session){
	c.Alphavantage.Apikey 	= decryptValue(c.Alphavantage.Apikey, session)
	c.AWS.SNS.ARN = decryptValue(c.AWS.SNS.ARN, session)
}

func getConfig() *Config {
	var config Config

	config.readFromFile("./resources/configs.yml")

	return &config
}

func initAwsSession(region string) (*session.Session, error) {

	return session.NewSession(&aws.Config{
		Region: aws.String(region),
		//Credentials: credentials.NewSharedCredentials("", "Personal"),
	})
}

func getAWSConfig() *AWSSessionConfig {
	session, _ := initAwsSession(getConfig().AWS.Region)
	db := dynamodb.New(session)
	GlobalConfig = getConfig()
	decryptSensitiveData(GlobalConfig, session)

	dynamodbConfig := Dynamodb{
		DynamodbSession: db,
		TableName: GlobalConfig.AWS.Dynamodb.Table,
	}

	config := AWSSessionConfig{
		Session: session,
		Dynamodb: dynamodbConfig,
	}

	return &config
}
