package dto

type Guest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	InviteID int    `json:"invite_id"`
}
