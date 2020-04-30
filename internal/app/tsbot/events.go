package tsbot

import (
	"log"

	"github.com/darfk/ts3"
)

func userProcess(client *ts3.Client, notification chan ts3.Notification) {
	for i := range notification {
		switch i.Type {
		case "notifyclientmoved", "notifycliententerview":
			clid := i.Params[0]["clid"]
			data, _ := client.Exec(GetClientInfo(clid))
			CheckClientPrivileges(client, data, clid)
			CheckClientNickname(client, data, clid)
		}
	}
}

func AddClientListeners(client *ts3.Client) {
	addJoinAndMoveListener(client)

	notification := make(chan ts3.Notification)
	go userProcess(client, notification)

	client.NotifyHandler(func(n ts3.Notification) {
		notification <- n
	})
}

// Start listening to client join/move events
func addJoinAndMoveListener(client *ts3.Client) {
	_, err := client.Exec(ts3.Command{
		Command: "servernotifyregister",
		Params: map[string][]string{
			"event": []string{"channel"},
			"id":    []string{"0"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
