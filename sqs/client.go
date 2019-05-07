package sqs

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/pkg/errors"
)

func client() (*sqs.SQS, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_ACCESS_SECRET_KEY"), ""),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to init client")
	}

	return sqs.New(sess), nil
}

func deleteMessage(queueUrl string, msg *sqs.Message) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl),
		ReceiptHandle: aws.String(*msg.ReceiptHandle),
	}
	sqsClient, err := client()
	if err != nil {
		return errors.Wrap(err, "failed to init client")
	}
	_, err = sqsClient.DeleteMessage(params)

	if err != nil {
		return err
	}
	return nil
}

func ReceiveMessage(queueUrl string) (string, error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(1),
	}
	sqsClient, err := client()
	if err != nil {
		return "", errors.Wrap(err, "failed to init client")
	}

	response, err := sqsClient.ReceiveMessage(params)
	if err != nil {
		return "", errors.Wrap(err, "failed to receive message")
	}

	if len(response.Messages) == 0 {
		return "", nil
	}

	message := response.Messages[0]
	body := message.Body

	err = deleteMessage(queueUrl, message)
	if err != nil {
		return "", errors.Wrap(err, "failed to delete message")
	}

	return *body, nil
}
