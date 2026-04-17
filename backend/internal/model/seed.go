package model

import (
	"time"

	"gorm.io/gorm"
)

// SeedData 初始化种子数据
func SeedData(db *gorm.DB) error {
	// 检查是否已有学生数据
	var count int64
	db.Model(&Student{}).Count(&count)
	if count > 0 {
		return nil
	}

	students := []Student{
		{Name: "张三", Gender: "男", Phone: "13800138001", Email: "zhangsan@example.com", Major: "计算机科学与技术", Class: "计科2401", StudentNo: "2024001", Status: 1},
		{Name: "李四", Gender: "女", Phone: "13800138002", Email: "lisi@example.com", Major: "软件工程", Class: "软工2401", StudentNo: "2024002", Status: 1},
		{Name: "王五", Gender: "男", Phone: "13800138003", Email: "wangwu@example.com", Major: "数据科学", Class: "数据2401", StudentNo: "2024003", Status: 1},
		{Name: "赵六", Gender: "女", Phone: "13800138004", Email: "zhaoliu@example.com", Major: "人工智能", Class: "AI2401", StudentNo: "2024004", Status: 2},
		{Name: "孙七", Gender: "男", Phone: "13800138005", Email: "sunqi@example.com", Major: "计算机科学与技术", Class: "计科2402", StudentNo: "2024005", Status: 3},
		{Name: "周八", Gender: "男", Phone: "13800138006", Email: "zhouba@example.com", Major: "信息安全", Class: "安全2401", StudentNo: "2024006", Status: 1},
		{Name: "吴九", Gender: "女", Phone: "13800138007", Email: "wujiu@example.com", Major: "软件工程", Class: "软工2402", StudentNo: "2024007", Status: 1},
		{Name: "郑十", Gender: "男", Phone: "13800138008", Email: "zhengshi@example.com", Major: "物联网工程", Class: "物联2401", StudentNo: "2024008", Status: 1},
	}

	for i := range students {
		students[i].CreatedAt = time.Now()
		students[i].UpdatedAt = time.Now()
	}

	return db.Create(&students).Error
}
