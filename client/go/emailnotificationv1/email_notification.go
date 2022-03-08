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

func SendEmail(
	receiver []EmailStruct,
	cc []EmailStruct,
	bbc []EmailStruct,
	replyTo []EmailStruct,
	subject string,
	tag string,
	text string,
	html string,
) (response *notificationemailv1.SendEmailResponse, err error) {
	d, err := time.ParseDuration(byEmailNotificationServiceTimeout)
	if err != nil {
		return
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(d))
	defer cancel()

	var arrayEmailReceiver []*notificationemailv1.EmailStruct
	var arrayEmailCc []*notificationemailv1.EmailStruct
	var arrayEmailBbc []*notificationemailv1.EmailStruct
	var arrayEmailReplyTo []*notificationemailv1.EmailStruct

	if len(receiver) == 0 {
		log.Printf("%s: ", "To is required")
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("%s: ", "To is required"),
		)
	}

	for _, item := range receiver {
		arrayEmailReceiver = append(arrayEmailReceiver, &notificationemailv1.EmailStruct{
			Name:  item.Name,
			Email: item.Email,
		})
	}

	for _, item := range cc {
		arrayEmailCc = append(arrayEmailCc, &notificationemailv1.EmailStruct{
			Name:  item.Name,
			Email: item.Email,
		})
	}

	for _, item := range bbc {
		arrayEmailBbc = append(arrayEmailBbc, &notificationemailv1.EmailStruct{
			Name:  item.Name,
			Email: item.Email,
		})
	}

	for _, item := range replyTo {
		arrayEmailReplyTo = append(arrayEmailReplyTo, &notificationemailv1.EmailStruct{
			Name:  item.Name,
			Email: item.Email,
		})
	}

	response, err = client.SendEmail(ctx, &notificationemailv1.SendEmailRequest{
		EmailData: &notificationemailv1.EmailData{
			Receiver: arrayEmailReceiver,
			Cc:       arrayEmailCc,
			Bbc:      arrayEmailBbc,
			ReplyTo:  arrayEmailReplyTo,
			Subject:  subject,
			Tag:      tag,
			Text:     text,
			Html:     html,
		},
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
