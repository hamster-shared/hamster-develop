package service

import (
	db2 "github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
)

type ITemplateService interface {
	//GetTemplateTypeList get template type list
	GetTemplateTypeList(templateType int) (*[]vo.TemplateTypeVo, error)
	//GetTemplatesByTypeId get templates by template type id
	GetTemplatesByTypeId(templateTypeId int) (*[]vo.TemplateVo, error)
	//GetTemplateDetail get template detail by template id
	GetTemplateDetail(templateId int) (*vo.TemplateDetailVo, error)
}

type TemplateService struct {
	db *gorm.DB
}

func NewTemplateService() *TemplateService {
	return &TemplateService{}
}

func (t *TemplateService) Init(db *gorm.DB) {
	t.db = db
}

func (t *TemplateService) GetTemplateTypeList(templateType int) (*[]vo.TemplateTypeVo, error) {
	var list []db2.TemplateType
	var listVo []vo.TemplateTypeVo
	result := t.db.Model(db2.TemplateType{}).Where("type = ?", templateType).Find(&list)
	if result.Error != nil {
		return &listVo, result.Error
	}
	if len(list) > 0 {
		copier.Copy(&listVo, &list)
	}
	return &listVo, nil
}

func (t *TemplateService) GetTemplatesByTypeId(templateTypeId int) (*[]vo.TemplateVo, error) {
	var list []db2.Template
	var listVo []vo.TemplateVo
	log.Println(t.db)
	result := t.db.Model(db2.TemplateType{}).Where("template_type_id = ?", templateTypeId).Find(&list)
	if result.Error != nil {
		return &listVo, result.Error
	}
	if len(list) > 0 {
		copier.Copy(&listVo, &list)
	}
	return &listVo, nil
}

func (t *TemplateService) GetTemplateDetail(templateId int) (*vo.TemplateDetailVo, error) {
	var data db2.TemplateDetail
	var dataVo vo.TemplateDetailVo
	result := t.db.Where("template_id = ? ", templateId).First(&data)
	if result.Error != nil {
		return &dataVo, result.Error
	}
	copier.Copy(&dataVo, &data)
	return &dataVo, nil
}
