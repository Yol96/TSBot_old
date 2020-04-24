package tsbot

import "github.com/darfk/ts3"

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

// Sends message to given user
func SendMessageUser(message, clid string) ts3.Command {
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
func PokeMessageUser(message, clid string) ts3.Command {
	return ts3.Command{
		Command: "clientpoke",
		Params: map[string][]string{
			"msg":  []string{message},
			"clid": []string{clid},
		},
	}
}

// Get info of given user
func GetClientInfo(clid string) ts3.Command {
	return ts3.Command{
		Command: "clientinfo",
		Params: map[string][]string{
			"clid": []string{clid},
		},
	}
}


