package service

import (
	"github.com/hamster-shared/a-line/pkg/model"
	"github.com/stretchr/testify/assert"
	ass "gotest.tools/v3/assert"
	"testing"
)

func TestCreateProject(t *testing.T) {
	projectService := NewProjectService()
	var project model.Project
	project.Name = "sun"
	err := projectService.CreateProject(&project)
	ass.NilError(t, err)
}

func TestUpdateProject(t *testing.T) {
	projectService := NewProjectService()
	var project model.Project
	project.Name = "jian"
	err := projectService.UpdateProject("test", &project)
	ass.NilError(t, err)
}

func TestGetProjects(t *testing.T) {
	projectService := NewProjectService()
	data := projectService.GetProjects("a", 1, 10)
	assert.NotEmpty(t, data.Data)
}

func TestDeleteProject(t *testing.T) {
	projectService := NewProjectService()
	err := projectService.DeleteProject("tom")
	ass.NilError(t, err)
}
