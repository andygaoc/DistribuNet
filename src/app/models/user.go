package models

import (
	"gorm.io/gorm"
	"time"
)

type InvitationCode struct {
	ID           int       `gorm:"primaryKey;autoIncrement;comment:'邀请码ID'"`
	Code         string    `gorm:"not null;comment:'邀请码'"`
	UserID       int       `gorm:"comment:'用户ID'"`
	GenerateTime time.Time `gorm:"comment:'生成时间'"`
	ExpiryTime   time.Time `gorm:"comment:'过期时间'"`
	Used         bool      `gorm:"comment:'是否已使用'"`
	UsedByUserID int       `gorm:"comment:'使用者用户ID'"`
	UsedTime     time.Time `gorm:"comment:'使用时间'"`
	Status       int       `gorm:"comment:'状态0.正常1时效'"`
	CreateTime   time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime   time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

type WithdrawalRecord struct {
	ID         int       `gorm:"primaryKey;comment:'提现记录ID'"`
	UserID     int       `gorm:"comment:'用户ID'"`
	Amount     float64   `gorm:"type:decimal(10,2);comment:'提现金额'"`
	Status     int       `gorm:"default:0;comment:'状态（0.待处理，1.处理中，2.完成，3.失败，4.取消）'"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

type Admin struct {
	ID         int       `gorm:"primaryKey;comment:'用户ID'"`
	Username   string    `gorm:"comment:'用户名'"`
	Password   string    `gorm:"comment:'密码'"`
	Email      string    `gorm:"comment:'电子邮件'"`
	Phone      string    `gorm:"comment:'手机号码'"`
	Status     int       `gorm:"default:0;comment:'状态0.正常'"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

type CommissionRate struct {
	ID         int       `gorm:"primaryKey;comment:'返佣比例ID'"`
	ObjectType int       `gorm:"comment:'返佣对象类型'"`
	ObjectID   int       `gorm:"comment:'返佣对象ID'"`
	ProductID  int       `gorm:"comment:'产品ID'"`
	Rate       float64   `gorm:"type:decimal(5,2);comment:'返佣比例'"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

type CommissionRecord struct {
	ID             int       `gorm:"primaryKey;comment:'返佣记录ID'"`
	UserID         int       `gorm:"comment:'返佣对象ID'"`
	Amount         float64   `gorm:"type:decimal(10,2);comment:'返佣金额'"`
	CommissionTime time.Time `gorm:"comment:'返佣时间'"`
	Status         int       `gorm:"comment:'状态1.未处理 2已经结算'"`
	CreateTime     time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime     time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

type CommissionSettlement struct {
	ID               int       `gorm:"primaryKey;comment:'结算ID'"`
	SettlementPeriod string    `gorm:"comment:'结算周期'"`
	UserID           int       `gorm:"comment:'返佣对象ID'"`
	Amount           float64   `gorm:"type:decimal(10,2);comment:'结算金额'"`
	SettlementTime   time.Time `gorm:"comment:'结算时间'"`
	CreateTime       time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime       time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

type CommissionAudit struct {
	ID         int       `gorm:"primaryKey;comment:'审核ID'"`
	RecordID   int       `gorm:"comment:'返佣记录ID'"`
	Status     int       `gorm:"comment:'审核状态1未审核，2审核通过3审核未通过'"`
	AuditTime  time.Time `gorm:"comment:'审核时间'"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

type CommissionReport struct {
	ID         int       `gorm:"primaryKey;comment:'报表ID'"`
	UserID     int       `gorm:"comment:'返佣对象ID'"`
	Amount     float64   `gorm:"type:decimal(10,2);comment:'返佣金额'"`
	ReportTime time.Time `gorm:"comment:'报表时间'"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

type Team struct {
	ID         int       `gorm:"primaryKey;comment:'团队ID'"`
	TeamName   string    `gorm:"comment:'团队名称'"`
	LeaderID   int       `gorm:"comment:'团队负责人ID'"`
	CreateTime time.Time `gorm:"comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
}

type Member struct {
	ID         int       `gorm:"primaryKey;comment:'成员ID'"`
	MemberName string    `gorm:"comment:'成员姓名'"`
	TeamID     int       `gorm:"comment:'团队ID'"`
	MemberRole int       `gorm:"comment:'成员角色，0普通成员，1主管负责人'"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

type CustomerRelation struct {
	ID           int       `gorm:"primaryKey;comment:'关系ID'"`
	TeamID       int       `gorm:"comment:'团队ID'"`
	CustomerID   int       `gorm:"comment:'客户ID'"`
	RelationType int       `gorm:"comment:'关系类型 0上下线'"`
	CreateTime   time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime   time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

type OrderRecord struct {
	ID         int       `gorm:"primaryKey;comment:'订单ID'"`
	CustomerID int       `gorm:"comment:'客户ID'"`
	OrderSN    string    `gorm:"comment:'订单编号'"`
	Amount     float64   `gorm:"type:decimal(10,2);comment:'订单金额'"`
	OrderTime  time.Time `gorm:"comment:'下单时间'"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&InvitationCode{}, &WithdrawalRecord{}, &Admin{}, &CommissionRate{}, &CommissionRecord{}, &CommissionSettlement{}, &CommissionAudit{}, &CommissionReport{}, &Team{}, &Member{}, &CustomerRelation{}, &OrderRecord{})
}
