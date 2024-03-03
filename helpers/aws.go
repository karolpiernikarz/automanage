package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/karolpiernikarz/automanage/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func AwsConfig() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(viper.GetString("aws.region")),
		Credentials: credentials.NewStaticCredentials(viper.GetString("aws.accesskeyid"), viper.GetString("aws.secretaccesskey"), viper.GetString("aws.token")),
	})
	return sess, err
}

func IsUserExist(bucketname string) bool {
	sess, err := AwsConfig()
	if err != nil {
		log.WithFields(log.Fields{"bucketname": bucketname}).Error("error connecting to aws")
		return true
	}
	svc := iam.New(sess)
	_, err = svc.GetUser(&iam.GetUserInput{
		UserName: aws.String(bucketname),
	})
	return err != nil
}

// IsPolicyExist wip
func IsPolicyExist(bucketname string) bool {
	sess, err := AwsConfig()
	if err != nil {
		log.WithFields(log.Fields{"bucketname": bucketname}).Error("error connecting to aws")
		return true
	}
	svc := iam.New(sess)
	arn := "arn:aws:iam::aws:policy/AWSLambdaExecute"
	_, err = svc.GetPolicy(&iam.GetPolicyInput{
		PolicyArn: &arn,
	})
	return err != nil
}

func IsBucketExist(bucketname string) bool {
	sess, err := AwsConfig()
	if err != nil {
		log.WithFields(log.Fields{"bucketname": bucketname}).Error("error connecting to aws")
		return true
	}
	svc := s3.New(sess)
	input := &s3.HeadBucketInput{
		Bucket: aws.String(bucketname),
	}
	_, err = svc.HeadBucket(input)
	return err != nil
}

func CreateS3Bucket(bucketname string, restaurantid string, domain string) (err error) {
	sess, err := AwsConfig()
	if err != nil {
		log.WithFields(log.Fields{"bucketname": bucketname}).Error("error connecting to aws")
		return err
	}
	svc := s3.New(sess)
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketname),
	}
	_, err = svc.CreateBucket(input)
	if err != nil {
		return err
	}
	taginput := &s3.PutBucketTaggingInput{
		Bucket: aws.String(bucketname),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String("restaurantid"),
					Value: aws.String(restaurantid),
				},
				{
					Key:   aws.String("domain"),
					Value: aws.String(domain),
				},
			},
		},
	}
	_, err = svc.PutBucketTagging(taginput)
	if err != nil {
		return err
	}
	return
}

func CreateIAMUser(bucketname string) (err error) {
	sess, err := AwsConfig()
	if err != nil {
		log.WithFields(log.Fields{"bucketname": bucketname}).Error("error connecting to aws")
		return err
	}
	svc := iam.New(sess)
	_, err = svc.CreateUser(&iam.CreateUserInput{
		UserName: aws.String(bucketname),
	})
	if err != nil {
		return err
	}
	return
}

func CreateS3BucketPolicy(bucketname string) (err error) {
	sess, err := AwsConfig()
	if err != nil {
		log.WithFields(log.Fields{"bucketname": bucketname}).Error("error connecting to aws")
		return err
	}
	svc := s3.New(sess)
	input := &s3.PutPublicAccessBlockInput{
		Bucket: aws.String(bucketname),
		PublicAccessBlockConfiguration: &s3.PublicAccessBlockConfiguration{
			BlockPublicAcls:       aws.Bool(false),
			BlockPublicPolicy:     aws.Bool(false),
			IgnorePublicAcls:      aws.Bool(false),
			RestrictPublicBuckets: aws.Bool(false),
		},
	}
	svc.PutPublicAccessBlock(input)
	s3policy := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Sid":    "AllowPublicRead",
				"Effect": "Allow",
				"Principal": map[string]interface{}{
					"AWS": "*",
				},
				"Action": []string{
					"s3:GetObject",
				},
				"Resource": []string{
					fmt.Sprintf("arn:aws:s3:::%s/*", bucketname),
				},
			},
		},
	}
	policy, err := json.Marshal(s3policy)
	if err != nil {
		return err
	}
	_, err = svc.PutBucketPolicy(&s3.PutBucketPolicyInput{
		Bucket: aws.String(bucketname),
		Policy: aws.String(string(policy)),
	})
	if err != nil {
		return err
	}
	return
}

func CreateIAMPolicy(bucketname string, restaurantdid string, domain string, fromaddress string) (err error) {
	sess, err := AwsConfig()
	if err != nil {
		log.WithFields(log.Fields{"bucketname": bucketname}).Error("error connecting to aws")
		return err
	}

	svc := iam.New(sess)
	policy := models.PolicyDocument{
		Version: "2012-10-17",
		Statement: []models.StatementEntry{
			{
				Effect: "Allow",
				Action: []string{
					"s3:GetBucketLocation",
					"s3:ListAllMyBuckets",
				},
				Resource: []string{
					"arn:aws:s3:::*",
				},
				Condition: map[string]interface{}{},
			},
			{
				Effect: "Allow",
				Action: []string{
					"s3:*",
				},
				Resource: []string{
					fmt.Sprintf("arn:aws:s3:::%s/*", bucketname),
					fmt.Sprintf("arn:aws:s3:::%s", bucketname),
				},
				Condition: map[string]interface{}{},
			},
			{
				Effect: "Allow",
				Action: []string{
					"ses:SendRawEmail",
				},
				Resource: []string{
					"*",
				},
				Condition: map[string]interface{}{
					"StringEquals": map[string]interface{}{
						"ses:FromAddress": fromaddress,
					},
				},
			},
		},
	}
	s3taginput := &iam.TagPolicyInput{
		PolicyArn: aws.String("arn:aws:iam::" + viper.GetString("aws.arn") + ":policy/" + bucketname),
		Tags: []*iam.Tag{
			{
				Key:   aws.String("ses"),
				Value: aws.String(""),
			},
			{
				Key:   aws.String("web"),
				Value: aws.String(""),
			},
			{
				Key:   aws.String("s3"),
				Value: aws.String(""),
			},
			{
				Key:   aws.String("restaurantid"),
				Value: aws.String(restaurantdid),
			},
			{
				Key:   aws.String("domain"),
				Value: aws.String(domain),
			},
		},
	}

	b, err := json.Marshal(&policy)
	if err != nil {
		return err
	}
	_, err = svc.CreatePolicy(&iam.CreatePolicyInput{
		PolicyDocument: aws.String(string(b)),
		PolicyName:     aws.String(bucketname),
	})
	if err != nil {
		return err
	}
	_, err = svc.TagPolicy(s3taginput)

	if err != nil {
		return err
	}
	return
}

func CreateAccessKey(username string) (accesskey string, secretkey string, err error) {
	sess, err := AwsConfig()
	if err != nil {
		return "", "", err
	}
	svc := iam.New(sess)
	result, err := svc.CreateAccessKey(&iam.CreateAccessKeyInput{
		UserName: aws.String(username),
	})
	if err != nil {
		return "", "", err
	}
	return *result.AccessKey.AccessKeyId, *result.AccessKey.SecretAccessKey, nil
}

func AttachIAMPolicy(username string) (err error) {
	sess, err := AwsConfig()
	if err != nil {
		return err
	}
	svc := iam.New(sess)
	policyArn := "arn:aws:iam::" + viper.GetString("aws.arn") + ":policy/" + username
	_, err = svc.AttachUserPolicy(&iam.AttachUserPolicyInput{
		PolicyArn: &policyArn,
		UserName:  aws.String(username),
	})
	if err != nil {
		return err
	}
	return
}

func GenerateSmtpCredentials(secretkey string, awsregion string) (smtppassword string, err error) {
	command := "./generatesmtp " + secretkey + " " + awsregion
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		log.WithFields(log.Fields{"output": output}).Error("error creating smtp creadentials")
		return "", err
	}
	return bytes.NewBuffer(output).String(), nil
}

// ChangeS3PolicyTag wip
func ChangeS3PolicyTag(domain string, bucketname string) error {
	sess, err := AwsConfig()
	if err != nil {
		return err
	}
	svc := s3.New(sess)
	taginput := &s3.PutBucketTaggingInput{
		Bucket: aws.String(bucketname),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String("domain"),
					Value: aws.String(domain),
				},
			},
		},
	}
	_, err = svc.PutBucketTagging(taginput)
	if err != nil {
		return err
	}
	return nil
}
