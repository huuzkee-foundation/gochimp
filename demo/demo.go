package main

import (
    "fmt"
    "github.com/mattbaird/gochimp"
    "os"
)

func main() {
    apiKey := os.Getenv("MANDRILL_KEY")
    mandrillApi, err := gochimp.NewMandrill(apiKey)

    if err != nil {
        fmt.Println("Error instantiating client")
    }

    templateName := "welcome email"
    contentVar := gochimp.Var{"main", "<h1>Welcome aboard!</h1>"}
    content := []gochimp.Var{contentVar}

    renderedTemplate, err := mandrillApi.TemplateRender(templateName, content, nil)

    if err != nil {
        fmt.Println("Error rendering template")
    }

    recipients := []gochimp.Recipient{
        gochimp.Recipient{Email: "newemployee@example.com"},
    }

    message := gochimp.Message{
        Html:      renderedTemplate,
        Subject:   "Welcome aboard!",
        FromEmail: "bossman@example.com",
        FromName:  "Boss Man",
        To:        recipients,
    }

    _, err = mandrillApi.MessageSend(message, false)

    if err != nil {
        fmt.Println("Error sending message")
    }
}