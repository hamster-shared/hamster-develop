package service

import (
	"errors"
	"fmt"
	"github.com/google/go-github/v48/github"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
	"net/url"
	"strings"
)

type ReportService struct {
	db *gorm.DB
}

func NewReportService() *ReportService {
	return &ReportService{
		db: application.GetBean[*gorm.DB]("db"),
	}
}

func (c *ReportService) QueryReports(projectId, reportType string, Type string, page int, size int) (vo.Page[db2.Report], error) {
	var total int64
	var reports []db2.Report
	tx := c.db.Model(db2.Report{}).Where("project_id = ?", projectId)
	if Type != "" && reportType != "" {
		tx = tx.Where("check_tool = ? and name=?", Type, reportType)
	} else if Type == "" && reportType == "" {
		tx = tx.Where("1 = 1")
	} else if Type != "" {
		tx = tx.Where("check_tool = ?", Type)
	} else {
		tx = tx.Where("name = ?", reportType)
	}
	tx = tx.Order("create_time desc")
	result := tx.Offset((page - 1) * size).Limit(size).Find(&reports).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return vo.NewEmptyPage[db2.Report](), result.Error
	}

	return vo.NewPage[db2.Report](reports, int(total), page, size), nil
}

func (c *ReportService) QueryFrontendReports(projectId string, page int, size int) (vo.Page[db2.Report], error) {
	var total int64
	var reports []db2.Report
	tx := c.db.Model(db2.Report{}).Where("project_id = ?", projectId)
	result := tx.Offset((page - 1) * size).Limit(size).Find(&reports).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return vo.NewEmptyPage[db2.Report](), result.Error
	}
	return vo.NewPage[db2.Report](reports, int(total), page, size), nil
}

func (c *ReportService) QueryReportsByWorkflow(workflowId, workflowDetailId int) ([]vo.ReportVo, error) {
	var reports []db2.Report
	var result []vo.ReportVo
	res := c.db.Model(db2.Report{}).Where("workflow_id = ? and workflow_detail_id = ?", workflowId, workflowDetailId).Find(&reports)
	if res.Error != nil {
		return result, res.Error
	}
	if len(reports) > 0 {
		copier.Copy(&result, &reports)
	}
	return result, nil
}

func (c *ReportService) ReportOverview(workflowId, workflowDetailId int) ([]vo.ReportOverView, error) {
	var reports []db2.Report
	var result []vo.ReportVo
	var data []vo.ReportOverView
	res := c.db.Model(db2.Report{}).Where("workflow_id = ? and workflow_detail_id = ?", workflowId, workflowDetailId).Find(&reports)
	if res.Error != nil {
		return data, res.Error
	}
	if len(reports) > 0 {
		copier.Copy(&result, &reports)
	}
	groupedReports := make(map[int][]vo.ReportVo)
	for _, p := range result {
		groupedReports[p.ToolType] = append(groupedReports[p.ToolType], p)
	}
	secutityAnalysisData, ok := groupedReports[1]
	if ok {
		reportRes := vo.ReportOverView{
			Title:   "Secutity Analysis",
			Content: secutityAnalysisData,
		}
		data = append(data, reportRes)
	}
	OpenSourceAnalysisData, ok := groupedReports[2]
	if ok {
		reportRes := vo.ReportOverView{
			Title:   "Open Source Analysis",
			Content: OpenSourceAnalysisData,
		}
		data = append(data, reportRes)
	}
	CodeQualityAnalysisData, ok := groupedReports[3]
	if ok {
		reportRes := vo.ReportOverView{
			Title:   "Code Quality Analysisi",
			Content: CodeQualityAnalysisData,
		}
		data = append(data, reportRes)
	}
	GasUsageAnalysisData, ok := groupedReports[4]
	if ok {
		reportRes := vo.ReportOverView{
			Title:   "Gas Usage Analysis",
			Content: GasUsageAnalysisData,
		}
		data = append(data, reportRes)
	}
	OtherAnalysisData, ok := groupedReports[5]
	if ok {
		reportRes := vo.ReportOverView{
			Title:   "AI Analysis",
			Content: OtherAnalysisData,
		}
		data = append(data, reportRes)
	}

	return data, nil
}

func (c *ReportService) ReportDetail(reportId int) (vo.ReportVo, error) {
	var reports db2.Report
	var result vo.ReportVo
	res := c.db.Model(db2.Report{}).Where("id = ?", reportId).Find(&reports)
	if res.Error != nil {
		return result, res.Error
	}
	copier.Copy(&result, &reports)
	return result, nil
}

func (c *ReportService) QueryFrontendReportsByWorkflow(workflowId, workflowDetailId int) ([]vo.ReportVo, error) {
	var reports []db2.Report
	var data []vo.ReportVo
	res := c.db.Model(db2.Report{}).Where("workflow_id = ? and workflow_detail_id = ?", workflowId, workflowDetailId).Find(&reports)
	if res.Error != nil {
		return data, res.Error
	}
	_ = copier.Copy(&data, &reports)
	return data, nil
}

func (c *ReportService) QueryReportCheckTools(projectId string) ([]string, error) {
	var data []string
	res := c.db.Model(db2.Report{}).Distinct("check_tool").Select("check_tool").Where("project_id = ?", projectId).Find(&data)
	if res.Error != nil {
		return data, res.Error
	}
	return data, nil
}

func (c *ReportService) GetFile(key string) (string, error) {
	token := utils.MetaScanHttpRequestToken()
	url := fmt.Sprintf("https://app.metatrust.io/api/scan/history/vulnerability-files/%s", key)
	res, err := utils.NewHttp().NewRequest().SetHeaders(map[string]string{
		"Authorization":  token,
		"X-MetaScan-Org": "1098616244203945984",
	}).Get(url)
	if err != nil {
		log.Println("获取失败")
		return "", err
	}
	if res.StatusCode() == 401 {
		log.Println("get file no auth")
		return "", errors.New("get file filed")
	}
	if res.StatusCode() != 200 {
		log.Println("meta scan file failed")
		return "", errors.New(fmt.Sprintf("%v", res.Error()))
	}
	return string(res.Body()), nil
}

func ParsingRepoType(repoType int, repoFiles []*github.RepositoryContent, userName, token string) (vo.RepoFrameType, error) {
	if repoType == int(consts.CONTRACT) {
		return parsingContractType(repoFiles, userName, token)
	} else if repoType == int(consts.FRONTEND) {
		return parsingFrontEndType(repoFiles, userName, token)
	}
	return vo.RepoFrameType{}, fmt.Errorf("repo type invalid")
}

func parsingContractType(repoFiles []*github.RepositoryContent, userName, token string) (vo.RepoFrameType, error) {
	var repoFrameType vo.RepoFrameType
	for _, v := range repoFiles {
		if strings.Contains(v.GetName(), "cairo") {
			repoFrameType.Type = consts.StarkWare
			return repoFrameType, nil
		} else if strings.Contains(v.GetName(), "Move.toml") {
			tomlType, err := parsingToml(v, v.GetName(), userName, token)
			if err != nil {
				return repoFrameType, err
			}
			repoFrameType.Type = tomlType
			return repoFrameType, nil
		} else if strings.Contains(v.GetName(), "truffle-config.js") {
			repoFrameType.Type = consts.Evm
			repoFrameType.EvmFrame = uint(consts.Truffle)
			return repoFrameType, nil
		} else if strings.Contains(v.GetName(), "foundry.toml") {
			repoFrameType.Type = consts.Evm
			repoFrameType.EvmFrame = uint(consts.Foundry)
			return repoFrameType, nil
		} else if strings.Contains(v.GetName(), "hardhat.config.js") {
			repoFrameType.Type = consts.Evm
			repoFrameType.EvmFrame = uint(consts.Hardhat)
			return repoFrameType, nil
		}
	}
	return repoFrameType, fmt.Errorf("parsing contract frame error")
}

func parsingFrontEndType(repoFiles []*github.RepositoryContent, userName, token string) (vo.RepoFrameType, error) {
	for _, v := range repoFiles {
		if strings.Contains(v.GetName(), "package.json") {
			frontEndType, err := parsingPackageJson(v, v.GetName(), userName, token)
			if err != nil {
				return vo.RepoFrameType{}, err
			}
			return vo.RepoFrameType{Type: frontEndType}, nil
		}
	}
	return vo.RepoFrameType{}, fmt.Errorf("parsing front end type err: package.json not exit")
}

func ParsingGitHubURL(urlStr string) (owner, repo string, err error) {
	// 移除 .git 后缀
	urlStr = strings.TrimSuffix(urlStr, ".git")

	u, err := url.Parse(urlStr)
	if err != nil {
		return "", "", err
	}

	if u.Host != "github.com" {
		return "", "", fmt.Errorf("invalid GitHub URL")
	}

	path := strings.TrimPrefix(u.Path, "/")
	segments := strings.Split(path, "/")
	if len(segments) < 2 {
		return "", "", fmt.Errorf("invalid GitHub URL")
	}

	owner = segments[0]
	repo = segments[1]

	return owner, repo, nil
}
