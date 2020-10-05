package records

//LoadFundsResponse - uuid generated after process of a file with load funds.
type LoadFundsResponse struct {
	ProcessID string `json:"process_id" bson:"process_id" example:"c01d7cf6-ec3f-47f0-9556-a5d6e9009a43"`
}
