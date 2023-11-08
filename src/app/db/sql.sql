

CREATE TABLE InvitationCode (
                                ID INT PRIMARY KEY AUTO_INCREMENT COMMENT '邀请码ID',
                                Code VARCHAR(255) NOT NULL COMMENT '邀请码',
                                UserID INT COMMENT '用户ID',
                                GenerateTime DATETIME COMMENT '生成时间',
                                ExpiryTime DATETIME COMMENT '过期时间',
                                Used BOOLEAN COMMENT '是否已使用',
                                Status INT COMMENT '状态0.正常1时效 ',
                                CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
)COMMENT='邀请码';

CREATE TABLE WithdrawalRecord (
                                  ID INT PRIMARY KEY COMMENT '提现记录ID',
                                  UserID INT COMMENT '用户ID',
                                  Amount DECIMAL(10, 2) COMMENT '提现金额',
                                  Status INT DEFAULT 0 COMMENT '状态（0.待处理，1.审核通过，2.完成，3.失败，4.取消，5审核未通过）',
                                  CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='提现记录表';

CREATE TABLE WithdrawalAuditRecord (
                                       ID INT PRIMARY KEY COMMENT '提现审核记录ID',
                                       WithdrawalID INT COMMENT '提现记录ID',
                                       AuditorID INT COMMENT '审核人员ID',
                                       AuditStatus INT DEFAULT 0 COMMENT '审核状态（0.待审核，1.审核通过，2.审核未通过）',
                                       AuditComment VARCHAR(255) COMMENT '审核备注',
                                       AuditTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '审核时间',
                                       CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                       UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='提现审核记录表';

CREATE TABLE User (
                      ID INT PRIMARY KEY COMMENT '自增ID',
                      UserID INT PRIMARY KEY COMMENT '用户ID',
                      Addr VARCHAR(255) COMMENT '用户名',
                      Status INT DEFAULT 0 COMMENT '状态0.正常 ',
                      CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                      UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'

) COMMENT='用户表';


CREATE TABLE CommissionRate (
                                ID INT PRIMARY KEY COMMENT '返佣比例ID',
                                ObjectType VARCHAR(255) COMMENT '返佣对象类型',
                                ObjectID INT COMMENT '返佣对象ID',
                                Rate DECIMAL(5,2) COMMENT '返佣比例',
                                CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='返佣比例表';

CREATE TABLE CommissionRecord (
                                  ID INT PRIMARY KEY COMMENT '返佣记录ID',
                                  UserID INT COMMENT '返佣对象ID',
                                  Amount DECIMAL(10,2) COMMENT '返佣金额',
                                  CommissionTime DATETIME COMMENT '返佣时间',
                                  Status INT COMMENT '状态1.未处理 2已经结算',
                                  CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='返佣记录表';

CREATE TABLE CommissionSettlement (
                                      ID INT PRIMARY KEY COMMENT '结算ID',
                                      SettlementPeriod VARCHAR(255) COMMENT '结算周期',
                                      UserID INT COMMENT '返佣对象ID',
                                      Amount DECIMAL(10,2) COMMENT '结算金额',
                                      SettlementTime DATETIME COMMENT '结算时间',
                                      CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                      UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='返佣结算表';

CREATE TABLE CommissionAudit (
                                 ID INT PRIMARY KEY COMMENT '审核ID',
                                 RecordID INT COMMENT '返佣记录ID',
                                 Status INT COMMENT '审核状态1未审核，2审核通过3审核未通过',
                                 AuditTime DATETIME COMMENT '审核时间',
                                 CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='返佣审核表';

CREATE TABLE CommissionReport (
                                  ID INT PRIMARY KEY COMMENT '报表ID',
                                  UserID INT COMMENT '返佣对象ID',
                                  Amount DECIMAL(10,2) COMMENT '返佣金额',
                                  ReportTime DATETIME COMMENT '报表时间',
                                  CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='返佣报表表';

CREATE TABLE Team (
                      ID INT PRIMARY KEY COMMENT '团队ID',
                      TeamName VARCHAR(255) COMMENT '团队名称',
                      LeaderID INT COMMENT '团队负责人ID',
                      CreateTime DATETIME COMMENT '创建时间',
                      CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                      UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='团队表';

CREATE TABLE Member (
                        ID INT PRIMARY KEY COMMENT '成员ID',
                        MemberName VARCHAR(255) COMMENT '成员姓名',
                        TeamID INT COMMENT '团队ID',
                        MemberRole INT COMMENT '成员角色，0普通成员，1主管负责人',
                        CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='成员表';

CREATE TABLE CustomerRelation (
                                  ID INT PRIMARY KEY COMMENT '关系ID',
                                  TeamID INT COMMENT '团队ID',
                                  CustomerID INT COMMENT '客户ID',
                                  RelationType int COMMENT '关系类型 0上下线',
                                  CreateTime DATETIME COMMENT '创建时间',
                                  CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                  UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='客户关系表';


CREATE TABLE OrderRecord (
                             ID INT PRIMARY KEY COMMENT '订单ID',
                             CustomerID INT COMMENT '客户ID',
                             OrderSN VARCHAR(255) COMMENT '订单编号',
                             Amount DECIMAL(10,2) COMMENT '订单金额',
                             OrderTime DATETIME COMMENT '下单时间',
                             CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             UpdateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) COMMENT='订单记录表';
