package store

import (
		"time"
)

type PlanProduct struct {
	PlanProductCode string
	PlanProductName     string
	PlanProductDescription string
}
type BusinessUnit struct {
	BusinessUnitCode string
	BusinessUnitName     string
	BusinessUnitDescription string
	BusinessType string
}
type Member struct {
	MemberCode        string
	FirstName     string
	LastName      string
	DOB 		 time.Time
	Email        string
	PhoneNumber  string
	EffectiveDate time.Time
	EndDate      time.Time
	BusinessUnit BusinessUnit
	PlanDetails PlanProduct
	CreatedAt time.Time
	UpdatedAt time.Time
}

