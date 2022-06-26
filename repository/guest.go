package repository

import (
	"github.com/bruceneco/party-invites/datastruct"
	"github.com/bruceneco/party-invites/dto"
	"github.com/bruceneco/party-invites/utils"
)

type GuestQuery interface {
	NewGuest(guest *dto.Guest) (id int, err error)
}

type guestQuery struct{}

func (gq *guestQuery) NewGuest(guest *dto.Guest) (id int, err error) {
	qb := pgQb().
		Insert(datastruct.GuestTableName).
		Columns("name", "invite_id").
		Values(guest.Name, guest.InviteID).
		Suffix("RETURNING id")

	err = qb.QueryRow().Scan(&id)
	if err != nil {
		utils.ErrorLog.Print(err)
	}
	return
}
