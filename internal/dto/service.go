package dto

type Service struct {
	ID              int64       `db:"id" json:"id"`
	Date            int64       `db:"date" json:"date"`
	Odometer        int64       `db:"odometer" json:"odometer"`
	UserID          int64       `db:"user_id" json:"user_id"`
	VehicleID       int64       `db:"vehicle_id" json:"vehicle_id"`
	ServiceTypeID   int64       `db:"service_type_id" json:"service_type_id"`
	Data            ServiceData `db:"data" json:"data"`
	VehicleName     string      `db:"vehicle_name" json:"vehicle_name"`
	ServiceTypeName string      `db:"service_type_name" json:"service_type_name"`
}
