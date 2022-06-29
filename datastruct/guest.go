package datastruct

import "database/sql"

const GuestTableName = "guests"

type Guest struct {
	ID       int           `db:"id"`
	Name     string        `db:"name"`
	InviteID sql.NullInt64 `db:"invite_id"`
}
