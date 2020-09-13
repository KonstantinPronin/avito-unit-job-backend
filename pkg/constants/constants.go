package constants

import "github.com/lib/pq"

//Postgres
const (
	// Class 23 - Integrity Constraint Violation
	IntegrityConstraintViolation = pq.ErrorClass("23")
)

//Internal
const (
	DefaultPageSize = 15
)
