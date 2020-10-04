package records

//RecordsResponse - Information of the result about the processing of the record
type RecordsResponse struct {
	ID         string `json:"id" example:"1234"`
	CustomerID string `json:"customer_id" example:"1234"`
	Accepted   bool   `json:"accepted"  example:true`
}
