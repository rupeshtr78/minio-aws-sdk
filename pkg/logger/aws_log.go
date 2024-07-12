package logger

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
)

// NewAWSLogger creates a new AWS logger
func NewAWSLogger() aws.Logger {
	return &awsLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

type awsLogger struct {
	logger *log.Logger
}

func (l *awsLogger) Log(args ...interface{}) {
	l.logger.Println(args...)
}

// WithAWSLogging adds logging to the AWS SDK client
func WithAWSLogging(client *s3.S3) {
	client.Handlers.Send.PushFront(func(r *request.Request) {
		log.Printf("Request: %s/%s, Params: %v\n",
			r.ClientInfo.ServiceName, r.Operation.Name, r.Params)
	})
}
