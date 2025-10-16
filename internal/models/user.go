package models

type UserStatus string

const (
	UserPending  UserStatus = "pending"
	UserApproved UserStatus = "approved"
	UserBanned   UserStatus = "banned"
)

type User struct {
	ID       int64      `db:"id"`       // Telegram chat ID
	Username string     `db:"username"` // Telegram username
	Status   UserStatus `db:"status"`
	JoinedAt string     `db:"joined_af"` // ISO timestamp
}
