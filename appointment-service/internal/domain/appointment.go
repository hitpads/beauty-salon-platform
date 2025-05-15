package domain

type Appointment struct {
	ID        string
	UserID    string
	MasterID  string
	ServiceID string
	StartTime string // Можно time.Time, но для простоты string
	Status    string // Например: "scheduled", "cancelled"
}
