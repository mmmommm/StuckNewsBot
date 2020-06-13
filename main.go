package main

import (
    "fmt"
    "os"
    "net/smtp"
)


func main() {

    from := "username@gmail.com"
    to := "test@gmail.com"

    // func PlainAuth(identity, username, password, host string) Auth
    auth := smtp.PlainAuth("", from, "password", "smtp.gmail.com")

    msg := []byte("" +
        "From: 送信した人 <" + from + ">\r\n" +
        "To: " + to + "\r\n" +
        "Subject: 件名 subject です\r\n" +
        "\r\n" +
        "テスト\r\n" +
    "")

    // func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
    err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, msg)
    if err != nil {
        fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
        return
    }

    fmt.Print("success")
}