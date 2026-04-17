package service

import (
	"errors"

	"student-admin/backend/internal/model"
	"student-admin/backend/internal/repository"
	pkgresponse "student-admin/backend/pkg/response"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

type CreateStudentRequest struct {
	Name      string `json:"name" binding:"required"`
	Gender    string `json:"gender"`
	Birthday  string `json:"birthday"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Major     string `json:"major"`
	Class     string `json:"class"`
	StudentNo string `json:"student_no" binding:"required"`
	Status    int    `json:"status"`
}

type UpdateStudentRequest struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Birthday  string `json:"birthday"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Major     string `json:"major"`
	Class     string `json:"class"`
	StudentNo string `json:"student_no"`
	Status    int    `json:"status"`
}

type ListStudentQuery struct {
	Keyword  string `form:"keyword"`
	Page     int    `form:"page" binding:"min=1"`
	PageSize int    `form:"pageSize" binding:"min=1,max=100"`
}

func (s *StudentService) Create(req CreateStudentRequest) error {
	student := model.Student{
		Name:      req.Name,
		Gender:    req.Gender,
		Phone:     req.Phone,
		Email:     req.Email,
		Address:   req.Address,
		Major:     req.Major,
		Class:     req.Class,
		StudentNo: req.StudentNo,
		Status:    req.Status,
	}

	return s.repo.Create(&student)
}

func (s *StudentService) GetByID(id uint) (*model.Student, error) {
	return s.repo.FindByID(id)
}

func (s *StudentService) Update(id uint, req UpdateStudentRequest) error {
	student, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("student not found")
	}

	if req.Name != "" {
		student.Name = req.Name
	}
	if req.Gender != "" {
		student.Gender = req.Gender
	}
	if req.Phone != "" {
		student.Phone = req.Phone
	}
	if req.Email != "" {
		student.Email = req.Email
	}
	if req.Address != "" {
		student.Address = req.Address
	}
	if req.Major != "" {
		student.Major = req.Major
	}
	if req.Class != "" {
		student.Class = req.Class
	}
	if req.StudentNo != "" {
		student.StudentNo = req.StudentNo
	}
	if req.Status != 0 {
		student.Status = req.Status
	}

	return s.repo.Update(student)
}

func (s *StudentService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *StudentService) List(query ListStudentQuery) (*pkgresponse.PageResult, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	students, total, err := s.repo.List(repository.StudentQuery{
		Keyword:  query.Keyword,
		Page:     query.Page,
		PageSize: query.PageSize,
	})
	if err != nil {
		return nil, err
	}

	return &pkgresponse.PageResult{
		List:     students,
		Total:    total,
		Page:     query.Page,
		PageSize: query.PageSize,
	}, nil
}
