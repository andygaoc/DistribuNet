package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Role 角色模型定义
type Role struct {
	ID         int       `gorm:"primaryKey;comment:'角色ID'"`
	Name       string    `gorm:"uniqueIndex;comment:'角色名称'"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

// Permission 权限模型定义
type Permission struct {
	ID         int       `gorm:"primaryKey;comment:'权限ID'"`
	Name       string    `gorm:"uniqueIndex;comment:'权限名称'"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

// User 用户模型定义
type User struct {
	ID         int       `gorm:"primaryKey;comment:'用户ID'"`
	Username   string    `gorm:"uniqueIndex;comment:'用户名'"`
	Password   string    `gorm:"comment:'密码'"`
	Email      string    `gorm:"uniqueIndex;comment:'电子邮件'"`
	Phone      string    `gorm:"uniqueIndex;comment:'手机号码'"`
	Status     int       `gorm:"comment:'状态0.正常'"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

// UserRole 用户角色模型定义
type UserRole struct {
	UserID int `gorm:"primaryKey;comment:'用户ID'"`
	RoleID int `gorm:"primaryKey;comment:'角色ID'"`
}

// RolePermission 角色权限模型定义
type RolePermission struct {
	RoleID       int `gorm:"primaryKey;comment:'角色ID'"`
	PermissionID int `gorm:"primaryKey;comment:'权限ID'"`
}

func main() {
	// 连接数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// 自动迁移数据库表结构
	db.AutoMigrate(&Role{}, &Permission{}, &User{}, &UserRole{}, &RolePermission{})

	// 创建角色
	createRole(db)

	// 创建权限
	createPermission(db)

	// 创建用户
	createUser(db)

	// 分配角色给用户
	assignRoleToUser(db)

	// 分配权限给角色
	assignPermissionToRole(db)

	// 查询用户角色和权限
	retrieveUserRolePermission(db)
}

// 创建角色
func createRole(db *gorm.DB) {
	role := Role{
		Name:       "admin",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	result := db.Create(&role)
	if result.Error != nil {
		log.Fatal("failed to create role:", result.Error)
	}
	fmt.Println("Create role:", role.ID)
}

// 创建权限
func createPermission(db *gorm.DB) {
	permission := Permission{
		Name:       "manage_users",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	result := db.Create(&permission)
	if result.Error != nil {
		log.Fatal("failed to create permission:", result.Error)
	}
	fmt.Println("Create permission:", permission.ID)
}

// 创建用户
func createUser(db *gorm.DB) {
	user := User{
		Username:   "admin",
		Password:   "123456",
		Email:      "admin@example.com",
		Phone:      "1234567890",
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("failed to create user:", result.Error)
	}
	fmt.Println("Create user:", user.ID)
}

// 分配角色给用户
func assignRoleToUser(db *gorm.DB) {
	userRole := UserRole{
		UserID: 1,
		RoleID: 1,
	}
	result := db.Create(&userRole)
	if result.Error != nil {
		log.Fatal("failed to assign role to user:", result.Error)
	}
	fmt.Println("Assign role to user:", userRole.UserID, userRole.RoleID)
}

// 分配权限给角色
func assignPermissionToRole(db *gorm.DB) {
	rolePermission := RolePermission{
		RoleID:       1,
		PermissionID: 1,
	}
	result := db.Create(&rolePermission)
	if result.Error != nil {
		log.Fatal("failed to assign permission to role:", result.Error)
	}
	fmt.Println("Assign permission to role:", rolePermission.RoleID, rolePermission.PermissionID)
}

// 查询用户角色和权限
func retrieveUserRolePermission(db *gorm.DB) {
	var user User
	result := db.Preload("Roles.Permissions").First(&user, 1)
	if result.Error != nil {
		log.Fatal("failed to retrieve user role permission:", result.Error)
	}
	fmt.Println("Retrieve user role permission:", user)
}
