package model

import (
	"github.com/Yol96/TSBot/internal/app/database"
)

type User struct {
	ID       int    `db:"user_id"`
	Nickname string `db:"username"`
	GroupID  int    `db:"group_id"`
	Tag      string `db:"tag"`
	TsID     string `db:"ts_id"`
}

func GetUserByTsId(teamspeakId string) (u User, err error) {
	query := `SELECT user_id, username, group_id, tag, ts_id FROM ts_users WHERE ts_id = ?`
	err = database.Db.Get(&u, query, teamspeakId)
	return
}
