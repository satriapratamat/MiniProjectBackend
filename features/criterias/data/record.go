package data

import (
	"campusmanagement/features/criterias"

	"gorm.io/gorm"
)

type Criteria struct {
	gorm.Model
	FacultyID  int
	CategoryID int
}

func toCriteriaRecord(criterias criterias.CriteriaCore) Criteria {
	return Criteria{
		FacultyID:  criterias.FacultyID,
		CategoryID: criterias.CategoryID,
	}
}

func toCriteriaCore(wl Criteria) criterias.CriteriaCore {
	return criterias.CriteriaCore{
		FacultyID:  wl.FacultyID,
		CategoryID: wl.CategoryID,
	}
}

func toCriteriaCoreList(wlList []Criteria) []criterias.CriteriaCore {
	criteriaConv := []criterias.CriteriaCore{}

	for _, acriterias := range wlList {
		criteriaConv = append(criteriaConv, toCriteriaCore(acriterias))
	}
	return criteriaConv
}
