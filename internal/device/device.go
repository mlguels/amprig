package device

// Status type is set to int
type Status int

// Here we have an iota which will start a count at 0 +
const (
	Idle Status = iota
	Running
	Warning
	Error
)

// This is used to return a string for the const iota instead of returning numbers
func (s Status) String() string {
	return [...]string{"Idle", "Running", "Warning", "Error"}[s]
}

// This is the Device interface which will work with any deviced that is plugged in
type Device interface {
	GetID() string
	GetStatus() Status
	EmergencyStop() error
}
