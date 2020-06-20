package domain

type Mail struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Subject string `json:"subject"`
	Text string `json:"text"`
}