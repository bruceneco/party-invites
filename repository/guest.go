package repository

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/bruceneco/party-invites/datastruct"
	"github.com/bruceneco/party-invites/utils"
)

type GuestQuery interface {
	NewGuest(guest *datastruct.Guest) (int, error)
	GetGuest(id int) (*datastruct.Guest, error)
	GetGuests() ([]datastruct.Guest, error)
}

type guestQuery struct{}

func (gq *guestQuery) NewGuest(guest *datastruct.Guest) (id int, err error) {
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

func (gq *guestQuery) GetGuest(id int) (*datastruct.Guest, error) {
	qb := pgQb().
		Select("id", "name", "invite_id").
		From(datastruct.GuestTableName).
		Where(squirrel.Eq{"id": id})
	var guest datastruct.Guest
	err := qb.QueryRow().Scan(&guest.ID, &guest.Name, &guest.InviteID)
	return &guest, err
}

func (gq *guestQuery) GetGuests() (guests []datastruct.Guest, err error) {
	qb := pgQb().
		Select("id", "name", "invite_id").
		From(datastruct.GuestTableName)

	var rows *sql.Rows
	rows, err = qb.Query()
	if err != nil {
		return
	}

	var guest datastruct.Guest
	for rows.Next() {
		err = rows.Scan(&guest.ID, &guest.Name, &guest.InviteID)
		if err != nil {
			return
		}
		guests = append(guests, guest)
	}
	return
}
