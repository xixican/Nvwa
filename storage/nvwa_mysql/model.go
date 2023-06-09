package nvwa_mysql

import "gorm.io/gorm"

type Memory struct {
	gorm.Model
	ID            int64  // 唯一Id
	AgentId       int    // agentId
	Content       string // 记忆内容
	Importance    int    // 记录重要度
	RecordTime    int64  // 记录时间
	LastVisitTime int64  // 上次访问时间
}
