package constants

// Application constants
const (
	// Server
	DefaultPort = "8080"

	// Database
	DefaultMongoDBURI = "mongodb+srv://preetheshvechoor4263:HuZRwxMk7nOWfmg7@cluster0.y1km5.mongodb.net/EmployeePortal"
	DefaultDBName     = "EmployeePortal"

	// Timeouts
	DBConnectionTimeout = 10 // seconds
	DBOperationTimeout  = 5  // seconds

	// Collections
	UsersCollection = "users"

	// Email
	WelcomeEmailSubject  = "Your account details"
	WelcomeEmailTemplate = "welcome_email.html"

	// Status codes
	StatusSuccess = "success"
	StatusError   = "error"

	// Messages
	MsgUserRegistered = "User registered successfully"
	MsgUserNotFound   = "User not found"
	MsgInvalidInput   = "Invalid input data"
	MsgDBError        = "Database error"
	MsgEmailError     = "Error sending email"
)
