package service

import (
	"testing"

	"student-admin/backend/internal/model"
	"student-admin/backend/internal/repository"
	pkgresponse "student-admin/backend/pkg/response"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&model.Student{})
	assert.NoError(t, err)
	return db
}

func setupService(t *testing.T) *StudentService {
	db := setupTestDB(t)
	repo := repository.NewStudentRepository(db)
	return NewStudentService(repo)
}

func strPtr(s string) *string { return &s }

// --- Create ---

func TestCreateStudent_Success(t *testing.T) {
	svc := setupService(t)
	err := svc.Create(CreateStudentRequest{
		Name:      "张三",
		StudentNo: "2024001",
		Gender:    "男",
		Phone:     "13800138001",
		Email:     "test@example.com",
		Birthday:  "2000-01-15",
		Status:    1,
	})
	assert.NoError(t, err)
}

func TestCreateStudent_DuplicateNo(t *testing.T) {
	svc := setupService(t)
	err := svc.Create(CreateStudentRequest{
		Name: "张三", StudentNo: "2024001",
	})
	assert.NoError(t, err)

	err = svc.Create(CreateStudentRequest{
		Name: "李四", StudentNo: "2024001",
	})
	assert.Error(t, err)
	assert.IsType(t, &pkgresponse.ValidationError{}, err)
	assert.Contains(t, err.Error(), "学号已存在")
}

func TestCreateStudent_InvalidPhone(t *testing.T) {
	svc := setupService(t)
	err := svc.Create(CreateStudentRequest{
		Name: "张三", StudentNo: "2024001", Phone: "1234",
	})
	assert.Error(t, err)
	assert.IsType(t, &pkgresponse.ValidationError{}, err)
	assert.Contains(t, err.Error(), "手机号")
}

func TestCreateStudent_InvalidEmail(t *testing.T) {
	svc := setupService(t)
	err := svc.Create(CreateStudentRequest{
		Name: "张三", StudentNo: "2024001", Email: "not-an-email",
	})
	assert.Error(t, err)
	assert.IsType(t, &pkgresponse.ValidationError{}, err)
	assert.Contains(t, err.Error(), "邮箱")
}

func TestCreateStudent_InvalidGender(t *testing.T) {
	svc := setupService(t)
	err := svc.Create(CreateStudentRequest{
		Name: "张三", StudentNo: "2024001", Gender: "未知",
	})
	assert.Error(t, err)
	assert.IsType(t, &pkgresponse.ValidationError{}, err)
	assert.Contains(t, err.Error(), "性别")
}

func TestCreateStudent_InvalidBirthday(t *testing.T) {
	svc := setupService(t)
	err := svc.Create(CreateStudentRequest{
		Name: "张三", StudentNo: "2024001", Birthday: "not-a-date",
	})
	assert.Error(t, err)
	assert.IsType(t, &pkgresponse.ValidationError{}, err)
	assert.Contains(t, err.Error(), "出生日期")
}

// --- Update ---

func TestUpdateStudent_ClearFields(t *testing.T) {
	svc := setupService(t)

	// 先创建
	err := svc.Create(CreateStudentRequest{
		Name: "张三", StudentNo: "2024001", Phone: "13800138001", Email: "test@example.com",
	})
	assert.NoError(t, err)

	// 清空 phone
	err = svc.Update(1, UpdateStudentRequest{
		Phone: strPtr(""),
	})
	assert.NoError(t, err)

	// 验证
	student, err := svc.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "", student.Phone)
	assert.Equal(t, "test@example.com", student.Email) // 未传的不变
}

func TestUpdateStudent_NotFound(t *testing.T) {
	svc := setupService(t)
	err := svc.Update(999, UpdateStudentRequest{
		Name: strPtr("新名字"),
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestUpdateStudent_SetBirthdayAndClear(t *testing.T) {
	svc := setupService(t)

	err := svc.Create(CreateStudentRequest{
		Name: "张三", StudentNo: "2024001",
	})
	assert.NoError(t, err)

	// 设置 birthday
	err = svc.Update(1, UpdateStudentRequest{
		Birthday: strPtr("2000-06-15"),
	})
	assert.NoError(t, err)

	student, _ := svc.GetByID(1)
	assert.NotNil(t, student.Birthday)

	// 清空 birthday
	err = svc.Update(1, UpdateStudentRequest{
		Birthday: strPtr(""),
	})
	assert.NoError(t, err)

	student, _ = svc.GetByID(1)
	assert.Nil(t, student.Birthday)
}

// --- List ---

func TestListStudent_DefaultPagination(t *testing.T) {
	svc := setupService(t)

	// 创建多条数据
	for i := 0; i < 15; i++ {
		svc.Create(CreateStudentRequest{
			Name: "学生", StudentNo: "202400" + string(rune('0'+i)),
		})
	}

	// 不传 page/pageSize → 默认 page=1, pageSize=10
	result, err := svc.List(ListStudentQuery{})
	assert.NoError(t, err)
	assert.Equal(t, 1, result.Page)
	assert.Equal(t, 10, result.PageSize)
	assert.Equal(t, int64(15), result.Total)
}

func TestListStudent_KeywordSearch(t *testing.T) {
	svc := setupService(t)
	svc.Create(CreateStudentRequest{Name: "张三", StudentNo: "2024001", Major: "计算机"})
	svc.Create(CreateStudentRequest{Name: "李四", StudentNo: "2024002", Major: "数学"})

	result, err := svc.List(ListStudentQuery{Keyword: "张", Page: 1, PageSize: 10})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result.Total)
}

// --- Delete ---

func TestDeleteStudent_Success(t *testing.T) {
	svc := setupService(t)
	svc.Create(CreateStudentRequest{Name: "张三", StudentNo: "2024001"})

	err := svc.Delete(1)
	assert.NoError(t, err)

	_, err = svc.GetByID(1)
	assert.Error(t, err) // 软删除后查不到
}
