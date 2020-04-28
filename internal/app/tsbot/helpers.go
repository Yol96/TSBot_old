package tsbot

import (
	"log"
	"strconv"
	"strings"

	"github.com/Yol96/TSBot/internal/app/model"
	"github.com/darfk/ts3"
)

// Adds client to given group
func AddClientServerGroup(sgid, cldbid string) ts3.Command {
	return ts3.Command{
		Command: "servergroupaddclient",
		Params: map[string][]string{
			"sgid":   []string{sgid},
			"cldbid": []string{cldbid},
		},
	}
}

// Deletes client from given group
func DeleteClientServerGroup(sgid, cldbid string) ts3.Command {
	return ts3.Command{
		Command: "servergroupdelclient",
		Params: map[string][]string{
			"sgid":   []string{sgid},
			"cldbid": []string{cldbid},
		},
	}
}

// Sends message to given client
func SendMessageClient(message, clid string) ts3.Command {
	return ts3.Command{
		Command: "sendtextmessage",
		Params: map[string][]string{
			"targetmode": []string{"1"},
			"target":     []string{clid},
			"msg":        []string{message},
		},
	}
}

// Pokes client with given message
func PokeMessageClient(message, clid string) ts3.Command {
	return ts3.Command{
		Command: "clientpoke",
		Params: map[string][]string{
			"msg":  []string{message},
			"clid": []string{clid},
		},
	}
}

// Get info of given client
func GetClientInfo(clid string) ts3.Command {
	return ts3.Command{
		Command: "clientinfo",
		Params: map[string][]string{
			"clid": []string{clid},
		},
	}
}

// Sets client privileges
func SetClientPrivileges(client *ts3.Client, clid string, cluid string) {
	u, err := model.GetUserByTsId(cluid)
	if err != nil {
		client.Exec(PokeMessageClient("Can`t find user`s TeamspeakID in database", clid))
		LockClient(client, clid)
		log.Println(err)
	}
	data, err := client.Exec(GetClientInfo(clid))
	if err != nil {
		log.Fatal(err)
	}

	cldbid := data.Params[0]["client_database_id"]
	client.Exec(AddClientServerGroup(strconv.Itoa(u.GroupID), cldbid))
}

func LockClient(client *ts3.Client, clid string) {
	data, err := client.Exec(GetClientInfo(clid))
	if err != nil {
		log.Fatal(err)
	}

	cldbid := data.Params[0]["client_database_id"]
	groups := strings.Split(data.Params[0]["client_servergroups"], ",")
	client.Exec(AddClientServerGroup("8", "cldbid"))
	for i := range groups {
		if groups[i] == "8" { // TODO: Change guest group id
			continue
		}

		_, err := client.Exec(DeleteClientServerGroup(groups[i], cldbid))
		if err != nil {
			log.Println(err)
		}
	}
}
