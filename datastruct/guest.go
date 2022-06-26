package datastruct

const GuestTableName = "guests"

type Guest struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	InviteID int    `db:"invite_id"`
}
