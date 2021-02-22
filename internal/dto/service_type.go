package dto

type ServiceType struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	FreqMiles int64  `db:"freq_miles"`
	FreqDays  int64  `db:"freq_days"`
	User      User   `db:"user"`
}
