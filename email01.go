/* 
   Sending an Email (SMTP) message 
   Gmail account -Click your profile name in the upper right corner
Click Manage your Google Account
Click on Security on the left
Scroll down until you see the slider to activate access for "less secure" apps
   */
   
package main

import (
    "log"
    "net/smtp"
)

func main() {

    // Configuration
    from := "my_email@gmail.com"            // update this to reflect your value
    password := "super_secret_password"     // update this to reflect your value
    to := []string{"recipient@email.com"}   // update this to reflect your value
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    // update this to be your message body 
    message := []byte("From: my_email@gmail.com\r\n" + "To: recipient@email.com\r\n" + "Subject: Testing email from go\r\n\n" + "My super secret message.")

    // Create authentication
    auth := smtp.PlainAuth("", from, password, smtpHost)

    // Send actual message
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
    if err != nil {
        log.Fatal(err)
    }
}
