package records

import (
	"time"
)

//RecordProcessedList - List of records processeds
type RecordProcessedList []RecordProcessed

//RecordProcessed - Information about a record processed
//Error Codes:
//10 - The Load ID more than once for a customer is not acceptable.
//20 - The record reached a maximum of 3 charges per day.
//30 - The record reached a maximum load of $5,000 per day.
//40 - The record reached a maximum load of $20,000 per week.
type RecordProcessed struct {
	ProcessID  string    `json:"process_id" bson:"process_id" example:"c01d7cf6-ec3f-47f0-9556-a5d6e9009a43"`
	ID         string    `json:"id" bson:"id" example:"1234"`
	CustomerID string    `json:"customer_id" bson:"customer_id" example:"1234"`
	LoadAmount string    `json:"load_amount" bson:"load_amount" example:"$3456.78"`
	Time       time.Time `json:"time" bson:"time" example:"2000-01-01T00:00:00Z"`
	Accepted   string    `json:"accepted" bson:"accepted" example:"true"`
	CodError   string    `json:"cod_error" bson:"cod_error" example:"10"`
	Message    string    `json:"message" bson:"message" example:"Message error!"`
}
