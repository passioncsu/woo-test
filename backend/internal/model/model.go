package model

import (
	"time"

	"gorm.io/gorm"
)

// Admin 管理员
type Admin struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Username  string         `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string         `gorm:"size:255;not null" json:"-"` // json:"-" 不返回密码
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Admin) TableName() string {
	return "admins"
}

// Student 学生
type Student struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Gender    string         `gorm:"size:10" json:"gender"`
	Birthday  *time.Time     `json:"birthday"`
	Phone     string         `gorm:"size:20" json:"phone"`
	Email     string         `gorm:"size:100" json:"email"`
	Address   string         `gorm:"size:255" json:"address"`
	Major     string         `gorm:"size:100" json:"major"`
	Class     string         `gorm:"size:100" json:"class"`
	StudentNo string         `gorm:"uniqueIndex;size:50;not null" json:"student_no"`
	Status    int            `gorm:"default:1" json:"status"` // 1:在读 2:休学 3:毕业
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Student) TableName() string {
	return "students"
}
