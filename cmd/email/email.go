package email

import (
    "log"
    "fmt"
    "context"
    "github.com/mailgun/mailgun-go/v3"
    "time"
)

func sendVerifyEmail(name string, link string, toEmail string){
    mg := mailgun.NewMailgun("sandboxd9e463e80cce4faf8a1485695622a12e.mailgun.org", "bea10734978742be4446bba47e3ecc85-e49cc42c-25380117")
    m := mg.NewMessage(
        name + ` <valami@valami.com>`,
        "Hello",
        "",
        toEmail,
    )
    m.SetHtml(`
    <html>
    <h3>Welcome, ` + name + `!</h3>
    <h4>Account Verification</h4>
    <form action=` + link +`>
        <input type="submit" value="Verify yout email addres" />
    </form>
    </html>`)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()

    resp, id, err := mg.Send(ctx, m)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("ID: %s Resp: %s\n", id, resp)
}