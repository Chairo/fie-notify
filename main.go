package main

import (
	"file-notify/notifier"
)

func main() {
	notify := notifier.NewNotifier()
	notify.Notify()
	for {
		msg, _ := notify.PPubSub.ReceiveMessage()
		notify.Update(msg.Payload)
	}
}
