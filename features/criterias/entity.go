package criterias

import (
	"time"
)

type CriteriaCore struct {
	ID         int
	FacultyID  int
	CategoryID int
	YearInput  int
	IndicatorA int
	IndicatorB int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Data interface {
	InsertCriteria(wlData CriteriaCore) (err error)
	SelectCriteria(wlData CriteriaCore) ([]CriteriaCore, error)
}

type Business interface {
	CreateCriteria(wlData CriteriaCore) error
	GetCriteria(wlData CriteriaCore) ([]CriteriaCore, error)
}
