package emailnotificationclientv1

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	notificationemailv1 "github.com/cuemby/by-email-notification-sdk/gen/go/by/notificationemail/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var client notificationemailv1.NotificationEmailAPIServiceClient
var doOnce sync.Once

var byEmailNotificationServiceUri string
var byEmailNotificationServiceTimeout string

type EmailStruct struct {
	Name  string
	Email string
}

type EmailData struct {
	To      []EmailStruct
	Cc      []EmailStruct
	Bbc     []EmailStruct
	ReplyTo []EmailStruct
	Subject string
	Tag     string
	Text    string
	Html    string
}

func init() {
	doOnce.Do(func() {
		byEmailNotificationServiceTimeout = os.Getenv("BY_EMAIL_NOTIFICATION_SERVICE_TIMEOUT")
		if byEmailNotificationServiceTimeout == "" {
			byEmailNotificationServiceTimeout = "30s"
		}
		byEmailNotificationServiceUri = os.Getenv("BY_EMAIL_NOTIFICATION_SERVICE_URI")
		con, err := grpc.Dial(byEmailNotificationServiceUri, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		client = notificationemailv1.NewNotificationEmailAPIServiceClient(con)
	})
}

func SendEmail(data EmailData) (response *notificationemailv1.SendEmailResponse, err error) {
	d, err := time.ParseDuration(byEmailNotificationServiceTimeout)
	if err != nil {
		return
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(d))
	defer cancel()

	var arrayEmailTo []*notificationemailv1.To
	var arrayEmailCc []*notificationemailv1.Cc
	var arrayEmailBbc []*notificationemailv1.Bbc
	var arrayEmailReplyTo []*notificationemailv1.ReplyTo

	if len(data.To) == 0 {
		log.Printf("%s: ", "To is required")
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("%s: ", "To is required"),
		)
	}

	for _, item := range data.To {
		arrayEmailTo = append(arrayEmailTo, &notificationemailv1.To{
			Name:  item.Name,
			Email: item.Email,
		})
	}

	for _, item := range data.Cc {
		arrayEmailCc = append(arrayEmailCc, &notificationemailv1.Cc{
			Name:  item.Name,
			Email: item.Email,
		})
	}

	for _, item := range data.Bbc {
		arrayEmailBbc = append(arrayEmailBbc, &notificationemailv1.Bbc{
			Name:  item.Name,
			Email: item.Email,
		})
	}

	for _, item := range data.ReplyTo {
		arrayEmailReplyTo = append(arrayEmailReplyTo, &notificationemailv1.ReplyTo{
			Name:  item.Name,
			Email: item.Email,
		})
	}

	response, err = client.SendEmail(ctx, &notificationemailv1.SendEmailRequest{
		To:      arrayEmailTo,
		Cc:      arrayEmailCc,
		Bbc:     arrayEmailBbc,
		ReplyTo: arrayEmailReplyTo,
		Subject: data.Subject,
		Tag:     data.Tag,
		Text:    data.Text,
		Html:    data.Html,
	})

	if err != nil {
		log.Printf("%s: %v", "Error send email", err)
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("%s: %v", "Error send email", err),
		)
	}
	return response, nil
}
