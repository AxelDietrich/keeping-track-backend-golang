package requests

type AddRecordRequest struct {
	Name   string
	Amount float64
}

type AmountRequest struct {
	Amount float64
}

type Login struct {
	Email    string
	Password string
}

type SignUp struct {
	Username	string
	Email		string
	Password	string
}