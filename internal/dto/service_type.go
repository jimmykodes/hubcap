package dto

type ServiceType struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	FreqMile int64  `db:"freq_mile"`
	FreqDays int64  `db:"freq_days"`
	User     User   `db:"user"`
}
