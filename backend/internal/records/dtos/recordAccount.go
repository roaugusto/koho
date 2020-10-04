package records

import (
	"time"
)

//RecordAccountList - Describes a list of records
type RecordAccountList []RecordAccount

//RecordAccount - Describes a record in an account
type RecordAccount struct {
	ID         string    `json:"id" bson:"id" example:"1234"`
	CustomerID string    `json:"customer_id" bson:"customer_id" example:"4567"`
	LoadAmount string    `json:"load_amount" bson:"load_amount" example:"$3456.78"`
	Time       time.Time `json:"time" bson:"time" example:"2000-01-01T00:00:00Z"`
}
