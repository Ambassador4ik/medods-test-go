package mail

import (
	"context"
	"github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

const domain string = "sandbox595731283ef1463d8a70dc0c40ec35b0.mailgun.org"
const privateAPIKey string = "3b58b56a6468ccc42c4a19b17d7ceb9d-2b755df8-3aaddf4e"

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
