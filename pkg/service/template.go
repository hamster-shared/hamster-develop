package service

import (
	"fmt"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
)

type ITemplateService interface {
	//GetTemplateTypeList get template type list
	GetTemplateTypeList(templateType int) (*[]vo.TemplateTypeVo, error)
	//GetTemplatesByTypeId get templates by template type id
	GetTemplatesByTypeId(templateTypeId, languageType, deployType int) (*[]vo.TemplateVo, error)
	//GetTemplateDetail get template detail by template id
	GetTemplateDetail(templateId int) (*vo.TemplateDetailVo, error)
	GetFrontendTemplateDetail(templateId int) (*vo.TemplateDetailVo, error)
	GetChainTemplateDetail(templateId int) (*vo.ChainTemplateVo, error)
	TemplateShow(templateType, languageType, deploymentType int) (*[]vo.TemplateVo, error)
	TemplateDownload(id int, repoName string) string
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

func (t *TemplateService) GetTemplatesByTypeId(templateTypeId, languageType, deployType int) (*[]vo.TemplateVo, error) {
	var list []db2.Template
	var listVo []vo.TemplateVo
	if deployType == 1 {
		result := t.db.Model(db2.Template{}).Where("template_type_id = ? and language_type = ? and deploy_type = ? ", templateTypeId, languageType, deployType).Find(&list)
		if result.Error != nil {
			return &listVo, result.Error
		}
		if len(list) > 0 {
			copier.Copy(&listVo, &list)
		}
	} else {
		result := t.db.Model(db2.Template{}).Where("template_type_id = ? and language_type = ?", templateTypeId, languageType).Find(&list)
		if result.Error != nil {
			return &listVo, result.Error
		}
		if len(list) > 0 {
			copier.Copy(&listVo, &list)
		}
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

func (t *TemplateService) GetFrontendTemplateDetail(templateId int) (*vo.TemplateDetailVo, error) {
	var data db2.TemplateDetail
	var dataVo vo.TemplateDetailVo
	result := t.db.Table("t_frontend_template_detail").Where("template_id = ? ", templateId).First(&data)
	if result.Error != nil {
		return &dataVo, result.Error
	}
	copier.Copy(&dataVo, &data)
	return &dataVo, nil
}

func (t *TemplateService) GetChainTemplateDetail(templateId int) (*vo.ChainTemplateVo, error) {
	var data db2.ChainTemplateDetail
	var dataVo vo.ChainTemplateVo
	result := t.db.Table("t_chain_template_detail").Where("template_id = ? ", templateId).First(&data)
	if result.Error != nil {
		return &dataVo, result.Error
	}
	copier.Copy(&dataVo, &data)
	return &dataVo, nil
}

func (t *TemplateService) TemplateShow(templateType, languageType, deploymentType int) (*[]vo.TemplateVo, error) {
	var list []db2.Template
	var listVo []vo.TemplateVo
	sql := ""
	if deploymentType == 1 || deploymentType == 3 {
		sql = "select  t.*  from t_template t left join t_template_type ttt on t.template_type_id = ttt.id where ttt.type = ? and t.whether_display = 1 and t.language_type = ? and t.deploy_type = ?"
		res := t.db.Raw(sql, templateType, languageType, deploymentType).Scan(&list)
		if res.Error != nil {
			return &listVo, res.Error
		}
	} else if deploymentType == 2 {
		log.Println(languageType)
		sql = "select  t.*  from t_template t left join t_template_type ttt on t.template_type_id = ttt.id where ttt.type = ? and t.whether_display = 1 and t.deploy_type is null or t.deploy_type != 3 and t.language_type = ?"
		res := t.db.Raw(sql, templateType, languageType).Scan(&list)
		if res.Error != nil {
			return &listVo, res.Error
		}
	} else {
		sql = "select  t.*  from t_template t left join t_template_type ttt on t.template_type_id = ttt.id where ttt.type = ? and t.whether_display = 1 and t.language_type = ?"
		res := t.db.Raw(sql, templateType, languageType).Scan(&list)
		if res.Error != nil {
			return &listVo, res.Error
		}
	}
	copier.Copy(&listVo, &list)
	return &listVo, nil
}

func (t *TemplateService) TemplateDownload(id int, repoName string) string {
	fmt.Printf("download template id is:%d", id)
	downloadUrl := fmt.Sprintf("https://github.com/hamster-template/%s/archive/refs/heads/master.zip", repoName)
	return downloadUrl
}
