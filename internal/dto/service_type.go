package dto

type ServiceType struct {
	ID        int64                `db:"id" json:"id"`
	Name      string               `db:"name" json:"name"`
	FreqMiles int64                `db:"freq_miles" json:"freq_miles"`
	FreqDays  int64                `db:"freq_days" json:"freq_days"`
	Questions ServiceTypeQuestions `db:"questions" json:"questions"`
	UserID    int64                `db:"user_id" json:"user_id"`
}
