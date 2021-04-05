package requests

type AddRecordRequest struct {
	Name   string
	Amount float64
}

type MoveToSavingsRequest struct {
	Amount float64
}
