package model

import (
	"github.com/Yol96/TSBot/internal/app/database"
)

type User struct {
	ID       int    `db:"user_id"`
	Nickname string `db:"username"`
	GroupID  int    `db:"group_id"`
	Bantime  string `db:"srvbantime"`
	Tag      string `db:"tag"`
	TsID     string `db:"ts_id"`
}

func GetUserByTsId(teamspeakId string) (u User, err error) {
	query := `SELECT * FROM phpbb_users WHERE ts_id = ?`
	err = database.Db.Get(&u, query, teamspeakId)
	return
}
