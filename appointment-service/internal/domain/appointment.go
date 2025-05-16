package domain

type Appointment struct {
	ID        string
	UserID    string
	MasterID  string
	ServiceID string
	StartTime string
	Status    string
}
