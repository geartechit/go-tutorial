package appmsg

const (
	// Success messages
	CreateEmployeeSuccess = "employee created successfully"
	UpdateEmployeeSuccess = "employee updated successfully"
	DeleteEmployeeSuccess = "employee deleted successfully"
	GetEmployeeSuccess    = "employee retrieved successfully"
	ListEmployeeSuccess   = "employee list retrieved successfully"

	// Error messages
	InvalidRequest      = "invalid request body"
	ValidationFailed    = "validation failed"
	EmployeeNotFound    = "employee not found"
	FailedToCreate      = "failed to create employee"
	FailedToUpdate      = "failed to update employee"
	FailedToDelete      = "failed to delete employee"
	FailedToGet         = "failed to get employee"
	FailedToList        = "failed to get employees"
	InternalServerError = "internal server error"
)
