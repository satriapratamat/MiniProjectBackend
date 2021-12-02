package factory

import (
	"campusmanagement/config"
	bussinesscr "campusmanagement/features/criterias/business"
	datacr "campusmanagement/features/criterias/data"
	pptcr "campusmanagement/features/criterias/presentation"
)

type presenter struct {
	CriteriaPresentation pptcr.CriteriaHandler
}

func Init() presenter {
	criteriaData := datacr.NewMySqlCriteria(config.DB)
	criteriaBusiness := bussinesscr.NewBusinessCriteria(criteriaData)
	return presenter{
		CriteriaPresentation: *pptcr.NewHandlerCriteria(criteriaBusiness),
	}
}
