package records

//RecordsResponse - Information of the result about the processing of the record
type RecordsResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}
