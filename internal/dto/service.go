package dto

type Service struct {
	ID          int64       `db:"id"`
	Date        int64       `db:"date"`
	Odometer    int64       `db:"odometer"`
	Data        ServiceData `db:"data"`
	User        User        `db:"user"`
	Vehicle     Vehicle     `db:"vehicle"`
	ServiceType ServiceType `db:"service_type"`
}
