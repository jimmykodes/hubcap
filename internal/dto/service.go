package dto

type Service struct {
	ID            int64       `db:"id"`
	Date          int64       `db:"date"`
	Odometer      int64       `db:"odometer"`
	Data          ServiceData `db:"data"`
	UserID        int64       `db:"user_id"`
	VehicleID     int64       `db:"vehicle_id"`
	ServiceTypeID int64       `db:"service_type_id"`
}
