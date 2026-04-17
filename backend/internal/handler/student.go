package handler

import (
	"strconv"

	"student-admin/backend/internal/service"

	pkgresponse "student-admin/backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	svc *service.StudentService
}

func NewStudentHandler(svc *service.StudentService) *StudentHandler {
	return &StudentHandler{svc: svc}
}

func (h *StudentHandler) Create(c *gin.Context) {
	var req service.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkgresponse.Fail(c, 400, 400, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Create(req); err != nil {
		pkgresponse.Fail(c, 500, 500, "create failed: "+err.Error())
		return
	}

	pkgresponse.OKWithMessage(c, "created", nil)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		pkgresponse.Fail(c, 400, 400, "invalid id")
		return
	}

	student, err := h.svc.GetByID(uint(id))
	if err != nil {
		pkgresponse.Fail(c, 404, 404, "student not found")
		return
	}

	pkgresponse.OK(c, student)
}

func (h *StudentHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		pkgresponse.Fail(c, 400, 400, "invalid id")
		return
	}

	var req service.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkgresponse.Fail(c, 400, 400, "invalid request: "+err.Error())
		return
	}

	if err := h.svc.Update(uint(id), req); err != nil {
		pkgresponse.Fail(c, 500, 500, "update failed: "+err.Error())
		return
	}

	pkgresponse.OKWithMessage(c, "updated", nil)
}

func (h *StudentHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		pkgresponse.Fail(c, 400, 400, "invalid id")
		return
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		pkgresponse.Fail(c, 500, 500, "delete failed: "+err.Error())
		return
	}

	pkgresponse.OKWithMessage(c, "deleted", nil)
}

func (h *StudentHandler) List(c *gin.Context) {
	var query service.ListStudentQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		pkgresponse.Fail(c, 400, 400, "invalid query: "+err.Error())
		return
	}

	result, err := h.svc.List(query)
	if err != nil {
		pkgresponse.Fail(c, 500, 500, "query failed: "+err.Error())
		return
	}

	pkgresponse.OK(c, result)
}
