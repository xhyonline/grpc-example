// mod 是 ORM 模型
package mod

// User 用户表
type User struct {
	// 主键
	ID int
	// 姓名
	Name string
	// 年龄
	Age int
	// 手机
	Phone []*Phone
}

// Phone 手机模型
type Phone struct {
	// 自身主键
	ID int
	// 外键关联
	UserID int
	// 手机号
	Number string
	// 类型
	Type int `gorm:"type:tinyint(1);"`
}

