package mail

import (
	"context"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

const domain string = ""
const privateAPIKey string = ""

func SendNewIpNotification(ip string, recipient string) {
	mg := mailgun.NewMailgun(domain, privateAPIKey)

	sender := "notify" + domain
	subject := "IP Change detected!"
	body := "Your account is being accessed from new IP:" + ip + "."

	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mg.Send(ctx, message)

	if err != nil {
		log.Print(err)
	}
}
