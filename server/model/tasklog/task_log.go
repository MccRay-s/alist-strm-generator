package tasklog

import (
	"time"
)

// TaskLog 状态常量
const (
	TaskLogStatusRunning   = "running"
	TaskLogStatusCompleted = "completed"
	TaskLogStatusFailed    = "failed"
	TaskLogStatusCancelled = "cancelled"
)

// TaskLog 任务日志模型
type TaskLog struct {
	ID                 uint       `json:"id" gorm:"primaryKey"`
	CreatedAt          time.Time  `json:"createdAt"`
	UpdatedAt          time.Time  `json:"updatedAt"`
	TaskID             uint       `json:"taskId" gorm:"not null;index"`
	Status             string     `json:"status" gorm:"not null"`
	Message            string     `json:"message" gorm:"type:text"`
	StartTime          time.Time  `json:"startTime" gorm:"not null"`
	EndTime            *time.Time `json:"endTime" gorm:"default:null"`
	Duration           int64      `json:"duration" gorm:"not null;default:0"` // 持续时间，单位为秒
	TotalFile          int        `json:"totalFile" gorm:"not null;default:0"`
	GeneratedFile      int        `json:"generatedFile" gorm:"not null;default:0"`
	SkipFile           int        `json:"skipFile" gorm:"not null;default:0"`
	OverwriteFile      int        `json:"overwriteFile" gorm:"not null;default:0"`
	MetadataCount      int        `json:"metadataCount" gorm:"not null;default:0"`      // 处理的元数据文件总数
	SubtitleCount      int        `json:"subtitleCount" gorm:"not null;default:0"`      // 处理的字幕文件总数
	MetadataDownloaded int        `json:"metadataDownloaded" gorm:"not null;default:0"` // 下载的元数据文件数
	SubtitleDownloaded int        `json:"subtitleDownloaded" gorm:"not null;default:0"` // 下载的字幕文件数
	FailedCount        int        `json:"failedCount" gorm:"not null;default:0"`        // 处理失败的文件数
}

// TableName 表名
func (TaskLog) TableName() string {
	return "task_logs"
}
