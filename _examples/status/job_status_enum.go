// Code generated by enum generator; DO NOT EDIT.
package status

import (
	"fmt"

	"database/sql/driver"
)

// JobStatus is the exported type for the enum
type JobStatus struct {
	name  string
	value int
}

func (e JobStatus) String() string { return e.name }

// MarshalText implements encoding.TextMarshaler
func (e JobStatus) MarshalText() ([]byte, error) {
	return []byte(e.name), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (e *JobStatus) UnmarshalText(text []byte) error {
	var err error
	*e, err = ParseJobStatus(string(text))
	return err
}

// Value implements the driver.Valuer interface
func (e JobStatus) Value() (driver.Value, error) {
	return e.name, nil
}

// Scan implements the sql.Scanner interface
func (e *JobStatus) Scan(value interface{}) error {
	if value == nil {
		*e = JobStatusValues()[0]
		return nil
	}

	str, ok := value.(string)
	if !ok {
		if b, ok := value.([]byte); ok {
			str = string(b)
		} else {
			return fmt.Errorf("invalid jobStatus value: %v", value)
		}
	}

	val, err := ParseJobStatus(str)
	if err != nil {
		return err
	}

	*e = val
	return nil
}

// ParseJobStatus converts string to jobStatus enum value
func ParseJobStatus(v string) (JobStatus, error) {

	switch v {
	case "active":
		return JobStatusActive, nil
	case "blocked":
		return JobStatusBlocked, nil
	case "inactive":
		return JobStatusInactive, nil
	case "unknown":
		return JobStatusUnknown, nil

	}

	return JobStatus{}, fmt.Errorf("invalid jobStatus: %s", v)
}

// MustJobStatus is like ParseJobStatus but panics if string is invalid
func MustJobStatus(v string) JobStatus {
	r, err := ParseJobStatus(v)
	if err != nil {
		panic(err)
	}
	return r
}

// GetJobStatusByID gets the correspondent jobStatus enum value by its ID (raw integer value)
func GetJobStatusByID(v int) (JobStatus, error) {
	switch v {
	case 1:
		return JobStatusActive, nil
	case 3:
		return JobStatusBlocked, nil
	case 2:
		return JobStatusInactive, nil
	case 0:
		return JobStatusUnknown, nil
	}
	return JobStatus{}, fmt.Errorf("invalid jobStatus value: %d", v)
}

// Public constants for jobStatus values
var (
	JobStatusActive   = JobStatus{name: "active", value: 1}
	JobStatusBlocked  = JobStatus{name: "blocked", value: 3}
	JobStatusInactive = JobStatus{name: "inactive", value: 2}
	JobStatusUnknown  = JobStatus{name: "unknown", value: 0}
)

// JobStatusValues returns all possible enum values
func JobStatusValues() []JobStatus {
	return []JobStatus{
		JobStatusActive,
		JobStatusBlocked,
		JobStatusInactive,
		JobStatusUnknown,
	}
}

// JobStatusNames returns all possible enum names
func JobStatusNames() []string {
	return []string{
		"active",
		"blocked",
		"inactive",
		"unknown",
	}
}
