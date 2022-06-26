package repository

import (
	"github.com/bruceneco/party-invites/datastruct"
	"github.com/bruceneco/party-invites/utils"
)

type GuestQuery interface {
	NewGuest() (id int, err error)
}

type guestQuery struct{}

func (gq *guestQuery) NewGuest() (id int, err error) {
	qb := pgQb().
		Insert(datastruct.GuestTableName).
		Columns("name").
		Values("Bruce lindo").
		Suffix("RETURNING id")

	err = qb.QueryRow().Scan(&id)
	if err != nil {
		utils.ErrorLog.Print(err)
	}
	return
}
