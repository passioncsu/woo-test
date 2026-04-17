package service

import (
	"errors"
	"regexp"
	"time"

	"student-admin/backend/internal/model"
	"student-admin/backend/internal/repository"
	pkgresponse "student-admin/backend/pkg/response"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)
)

func validateGender(gender string) error {
	if gender == "" {
		return nil
	}
	if gender != "男" && gender != "女" {
		return pkgresponse.NewValidationError("性别必须为 男 或 女")
	}
	return nil
}

func validateStatus(status int) error {
	if status != 0 && (status < 1 || status > 3) {
		return pkgresponse.NewValidationError("状态必须为 1(在读)、2(休学) 或 3(毕业)")
	}
	return nil
}

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
	Name      *string `json:"name"`
	Gender    *string `json:"gender"`
	Birthday  *string `json:"birthday"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
	Address   *string `json:"address"`
	Major     *string `json:"major"`
	Class     *string `json:"class"`
	StudentNo *string `json:"student_no"`
	Status    *int    `json:"status"`
}

type ListStudentQuery struct {
	Keyword  string `form:"keyword"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

func (s *StudentService) Create(req CreateStudentRequest) error {
	// 验证学号唯一性
	exists, err := s.repo.ExistsByStudentNo(req.StudentNo, 0)
	if err != nil {
		return err
	}
	if exists {
		return pkgresponse.NewValidationError("学号已存在")
	}

	// 验证字段
	if err := validateGender(req.Gender); err != nil {
		return err
	}
	if err := validateStatus(req.Status); err != nil {
		return err
	}
	if req.Phone != "" && !phoneRegex.MatchString(req.Phone) {
		return pkgresponse.NewValidationError("手机号格式不正确")
	}
	if req.Email != "" && !emailRegex.MatchString(req.Email) {
		return pkgresponse.NewValidationError("邮箱格式不正确")
	}

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

	if req.Birthday != "" {
		t, err := time.Parse("2006-01-02", req.Birthday)
		if err != nil {
			return pkgresponse.NewValidationError("出生日期格式不正确，应为 YYYY-MM-DD")
		}
		d := model.Date(t)
		student.Birthday = &d
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

	// 验证学号唯一性
	if req.StudentNo != nil && *req.StudentNo != "" {
		exists, err := s.repo.ExistsByStudentNo(*req.StudentNo, id)
		if err != nil {
			return err
		}
		if exists {
			return pkgresponse.NewValidationError("学号已存在")
		}
	}

	// 验证字段
	if req.Gender != nil {
		if err := validateGender(*req.Gender); err != nil {
			return err
		}
	}
	if req.Status != nil {
		if err := validateStatus(*req.Status); err != nil {
			return err
		}
	}
	if req.Phone != nil && *req.Phone != "" && !phoneRegex.MatchString(*req.Phone) {
		return pkgresponse.NewValidationError("手机号格式不正确")
	}
	if req.Email != nil && *req.Email != "" && !emailRegex.MatchString(*req.Email) {
		return pkgresponse.NewValidationError("邮箱格式不正确")
	}

	// 应用更新（nil = 未提供不修改，非 nil = 更新值含空串）
	if req.Name != nil {
		student.Name = *req.Name
	}
	if req.Gender != nil {
		student.Gender = *req.Gender
	}
	if req.Phone != nil {
		student.Phone = *req.Phone
	}
	if req.Email != nil {
		student.Email = *req.Email
	}
	if req.Address != nil {
		student.Address = *req.Address
	}
	if req.Major != nil {
		student.Major = *req.Major
	}
	if req.Class != nil {
		student.Class = *req.Class
	}
	if req.StudentNo != nil {
		student.StudentNo = *req.StudentNo
	}
	if req.Status != nil {
		student.Status = *req.Status
	}
	if req.Birthday != nil {
		if *req.Birthday == "" {
			student.Birthday = nil
		} else {
			t, err := time.Parse("2006-01-02", *req.Birthday)
			if err != nil {
				return pkgresponse.NewValidationError("出生日期格式不正确，应为 YYYY-MM-DD")
			}
			d := model.Date(t)
			student.Birthday = &d
		}
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
	if query.PageSize <= 0 || query.PageSize > 100 {
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
