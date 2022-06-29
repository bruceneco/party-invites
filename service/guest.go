package service

import (
	"github.com/bruceneco/party-invites/datastruct"
	"github.com/bruceneco/party-invites/dto"
	"github.com/bruceneco/party-invites/repository"
	"github.com/bruceneco/party-invites/utils"
)

type GuestService interface {
	GetGuest(id int) (*dto.Guest, error)
	GetGuests() ([]dto.Guest, error)
}

type guestService struct {
	dao repository.DAO
}

func NewGuestService(dao repository.DAO) GuestService {
	return &guestService{dao: dao}
}

func (gs *guestService) GetGuest(id int) (g *dto.Guest, err error) {
	guestGot := new(datastruct.Guest)
	guestGot, err = gs.dao.NewGuestQuery().GetGuest(id)
	if err != nil {
		return
	}
	utils.InfoLog.Print(guestGot.InviteID.Valid)
	g = &dto.Guest{
		ID:       guestGot.ID,
		Name:     guestGot.Name,
		InviteID: int(guestGot.InviteID.Int64),
	}
	return
}

func (gs *guestService) GetGuests() (g []dto.Guest, err error) {
	var guests []datastruct.Guest
	guests, err = gs.dao.NewGuestQuery().GetGuests()
	if err != nil {
		return
	}
	for _, guest := range guests {
		g = append(g, dto.Guest{
			ID:       guest.ID,
			Name:     guest.Name,
			InviteID: int(guest.InviteID.Int64),
		})
	}
	return
}
