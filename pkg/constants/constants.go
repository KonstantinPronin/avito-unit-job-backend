package constants

import "github.com/lib/pq"

//Postgres code error classes
const (
	// Class 23 - Integrity Constraint Violation
	IntegrityConstraintViolation = pq.ErrorClass("23")
)
