package main

import "fmt"

type Notifier interface {
	notify()
}

type MailNotifier struct {
	To string
	From string
	Server string
}

func (mailNotifier MailNotifier) notify()  {
	fmt.Println(mailNotifier.From)
}

func main() {
	var notifier Notifier = MailNotifier{
		"jacquirhatim@gmail.com",
		"hello.notifications@domain.com",
		"smtp.localhost.com",
	}

	notifier.notify()
}