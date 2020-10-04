package records

import (
	"time"
)

//RecordProcessedList - List of records processeds
type RecordProcessedList []RecordProcessed

//RecordProcessed - Information about a record processed
type RecordProcessed struct {
	ID         string    `json:"id" bson:"id"`
	CustomerID string    `json:"customer_id" bson:"customer_id"`
	LoadAmount string    `json:"load_amount" bson:"load_amount"`
	Time       time.Time `json:"time" bson:"time"`
	Accepted   string    `json:"accepted" bson:"accepted"`
	CodError   string    `json:"cod_error" bson:"cod_error"`
	Message    string    `json:"message" bson:"message"`
}
