package repository

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/bruceneco/party-invites/utils"
	_ "github.com/glebarez/go-sqlite"
)

type DAO interface {
	NewGuestQuery() GuestQuery
	NewInviteQuery() InviteQuery
}

var DB *sql.DB

type dao struct{}

func pgQb() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(DB)
}
func NewDAO() DAO {
	return &dao{}
}

func NewDB(filepath string) {
	var err error
	DB, err = sql.Open("sqlite", filepath)
	if err != nil {
		utils.ErrorLog.Fatalln(err)
	}
	if DB == nil {
		utils.ErrorLog.Fatalln("db is nil")
	}
}

func (d *dao) NewGuestQuery() GuestQuery {
	return &guestQuery{}
}

func (d *dao) NewInviteQuery() InviteQuery {
	return &inviteQuery{}
}
