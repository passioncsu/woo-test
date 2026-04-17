package repository

import (
	"student-admin/backend/internal/model"

	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) Create(student *model.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepository) FindByID(id uint) (*model.Student, error) {
	var student model.Student
	if err := r.db.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) Update(student *model.Student) error {
	return r.db.Save(student).Error
}

func (r *StudentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Student{}, id).Error
}

type StudentQuery struct {
	Keyword  string
	Page     int
	PageSize int
}

func (r *StudentRepository) List(q StudentQuery) ([]model.Student, int64, error) {
	var students []model.Student
	var total int64

	db := r.db.Model(&model.Student{})

	if q.Keyword != "" {
		like := "%" + q.Keyword + "%"
		db = db.Where("name LIKE ? OR student_no LIKE ? OR major LIKE ?", like, like, like)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (q.Page - 1) * q.PageSize
	if err := db.Order("id DESC").Offset(offset).Limit(q.PageSize).Find(&students).Error; err != nil {
		return nil, 0, err
	}

	return students, total, nil
}
