package request

import "campusmanagement/features/criterias"

type ReqCriteria struct {
	FacultyID  int `json:"faculty_id"`
	CategoryID int `json:"category_id"`
}

func (reqData *ReqCriteria) ToCriteriaCore() criterias.CriteriaCore {
	return criterias.CriteriaCore{
		FacultyID:  reqData.FacultyID,
		CategoryID: reqData.CategoryID,
	}
}
