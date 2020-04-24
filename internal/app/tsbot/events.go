package tsbot

import (
	"log"
	"strings"

	"github.com/darfk/ts3"
)

func logTest(client *ts3.Client, notification chan ts3.Notification) {
	for i := range notification {
		switch i.Type {
		case "notifyclientmoved", "notifycliententerview":
			// log.Println(client, i.Params[0]["clid"])
			// client.Exec(SendMessageUser("test", i.Params[0]["clid"]))
			// client.Exec(PokeMessageUser("test", i.Params[0]["clid"]))
			data, _ := client.Exec(GetClientInfo(i.Params[0]["clid"]))
			cldbid := data.Params[0]["client_database_id"]
			groups := strings.Split(data.Params[0]["client_servergroups"], ",")
			nickname := data.Params[0]["client_nickname"]
			log.Println(cldbid)
			log.Println(groups)
			log.Println(nickname)
		}
	}
}

func AddClientListeners(client *ts3.Client) {
	addJoinAndMoveListener(client)

	notification := make(chan ts3.Notification)
	go logTest(client, notification)

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
