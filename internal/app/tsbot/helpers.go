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

// Gets info of given client
func GetClientInfo(clid string) ts3.Command {
	return ts3.Command{
		Command: "clientinfo",
		Params: map[string][]string{
			"clid": []string{clid},
		},
	}
}

// Checks client privileges
func CheckClientPrivileges(client *ts3.Client, data ts3.Response, clid string) bool {
	cluid := data.Params[0]["client_unique_identifier"]
	nickname := data.Params[0]["client_nickname"]
	log.Printf("Checking %s(%s) privileges", nickname, cluid)
	u, err := model.GetUserByTsId(cluid)
	if err != nil {
		log.Println(err)
	}

	if u.TsID != cluid {
		client.Exec(PokeMessageClient("TeamspeakID не найдет в базе данных", clid))
		log.Printf("Poke %s (Can`t find tsId in DB)", nickname)
		LockClient(client, data, clid)
		return false
	}

	cldbid := data.Params[0]["client_database_id"]
	client.Exec(AddClientServerGroup(strconv.Itoa(u.GroupID), cldbid))
	return true
}

// Lock client (deletes all groups, except guest group)
func LockClient(client *ts3.Client, data ts3.Response, clid string) {
	cldbid := data.Params[0]["client_database_id"]
	nickname := data.Params[0]["client_nickname"]
	log.Printf("Removing privileges from %s", nickname)
	groups := strings.Split(data.Params[0]["client_servergroups"], ",")

	client.Exec(AddClientServerGroup("8", cldbid))
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

// Checks client nickname
func CheckClientNickname(client *ts3.Client, data ts3.Response, clid string) {
	cluid := data.Params[0]["client_unique_identifier"]
	nickname := data.Params[0]["client_nickname"]
	log.Printf("Checking %s(%s) nickname", nickname, cluid)

	u, err := model.GetUserByTsId(cluid)
	if err != nil {
		log.Println(err)
	}

	databaseNickname := u.Nickname

	if u.Tag != "" {
		databaseNickname = "[" + u.Tag + "]" + u.Nickname
	}

	if nickname != databaseNickname {
		log.Printf("Poke %s (Nickname in TS doesn`t match nickname in DB: %s)", nickname, databaseNickname)
		client.Exec(PokeMessageClient("Никнейм в teamspeak не совпадает с никнеймом в базе данных. Измени никнейм на "+databaseNickname, clid))
	}
}
