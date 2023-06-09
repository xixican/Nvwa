package nvwa_mysql

import (
	"Nvwa/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func InitMysqlClient() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/agent?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.NvwaLog.Errorf("gorm.open 出错了，请检查dsn")
		return
	}
	// 创建表
	gormDB.Migrator().CreateTable(&Memory{})
}

func StoreMemoryRecord(agentId int, content string, recordTime int64, lastVisitTime int64, importance int) int64 {
	memory := &Memory{
		AgentId:       agentId,
		Content:       content,
		RecordTime:    recordTime,
		LastVisitTime: lastVisitTime,
		Importance:    importance,
	}
	db := gormDB.Create(memory)
	if db.Error != nil {
		logger.NvwaLog.Errorf("MysqlUpdateMemoryVisitTime error:%v", db.Error)
		return 0
	}
	return memory.ID
}

func MysqlUpdateMemoryVisitTime(visitTime int64) {
	db := gormDB.Model(&Memory{}).Update("lastVisitTime", visitTime)
	if db.Error != nil {
		logger.NvwaLog.Errorf("MysqlUpdateMemoryVisitTime error:%v", db.Error)
	}
}
