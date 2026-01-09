package entity

import "time"

type TableInfo struct {
	TableName    string    `json:"tableName" gorm:"column:TABLE_NAME"`
	TableComment string    `json:"tableComment" gorm:"column:TABLE_COMMENT"`
	CreateTime   time.Time `json:"createTime" gorm:"column:CREATE_TIME"`
}

type ColumnInfo struct {
	ColumnName    string `json:"columnName" gorm:"column:COLUMN_NAME"`
	DataType      string `json:"dataType" gorm:"column:DATA_TYPE"`
	ColumnComment string `json:"columnComment" gorm:"column:COLUMN_COMMENT"`
	ColumnKey     string `json:"columnKey" gorm:"column:COLUMN_KEY"` // PRI, UNI
	Extra         string `json:"extra" gorm:"column:EXTRA"`           // auto_increment
}
