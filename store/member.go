package store

import (
	"context"
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
	BusinessType bool
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

type Store interface {
	CreateMember(ctx context.Context, member *Member) error
	GetMember(ctx context.Context, id string) (*Member, error)
	UpdateMember(ctx context.Context, member *Member) error
	DeleteMember(ctx context.Context, id string) error
}