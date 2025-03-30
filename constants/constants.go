package constants

const (
	// App configuration
	AppName    = "My Go Application"
	AppVersion = "1.0.0"

	// API constants
	APIPrefix    = "/api/v1"
	DefaultLimit = 25
	MaxFileSize  = 10 << 20 // 10 MB

	// Status messages
	StatusSuccess = "success"
	StatusError   = "error"

	// Common strings
	EmptyString = ""
)

// Numeric constants
const (
	SecondsPerMinute = 60
	MinutesPerHour   = 60
	HoursPerDay      = 24
)

// Error messages
const (
	ErrInvalidInput = "invalid input"
	ErrNotFound     = "resource not found"
	ErrUnauthorized = "unauthorized access"
)