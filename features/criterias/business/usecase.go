package business

import "campusmanagement/features/criterias"

type CriteriaBusiness struct {
	criteriatData criterias.Data
}

func NewBusinessCriteria(criteriaData criterias.Data) criterias.Business {
	return &CriteriaBusiness{criteriaData}
}

func (wlBusiness *CriteriaBusiness) CreateCriteria(wlData criterias.CriteriaCore) error {
	if err := wlBusiness.criteriatData.InsertCriteria(wlData); err != nil {
		return err
	}
	return nil
}

func (wlBusiness *CriteriaBusiness) GetCriteria(wlData criterias.CriteriaCore) ([]criterias.CriteriaCore, error) {
	criterias, err := wlBusiness.criteriatData.SelectCriteria(wlData)
	if err != nil {
		return nil, err
	}
	return criterias, nil
}
