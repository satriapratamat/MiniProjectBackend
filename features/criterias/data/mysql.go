package data

import (
	"campusmanagement/features/criterias"

	"gorm.io/gorm"
)

type CriteriaData struct {
	db *gorm.DB
}

func NewMySqlCriteria(db *gorm.DB) criterias.Data {
	return &CriteriaData{db}
}

func (wlData *CriteriaData) InsertCriteria(criterias criterias.CriteriaCore) error {
	convData := toCriteriaRecord(criterias)

	if err := wlData.db.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}

func (wlData *CriteriaData) SelectCriteria(criterias criterias.CriteriaCore) ([]criterias.CriteriaCore, error) {
	var acriterias []Criteria

	err := wlData.db.Find(&acriterias).Error
	if err != nil {
		return nil, err
	}
	return toCriteriaCoreList(acriterias), nil
}
