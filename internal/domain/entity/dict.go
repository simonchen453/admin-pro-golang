package entity

import "time"

// DictType 对应 sys_dict_tbl
type DictType struct {
	ID          string    `json:"id" gorm:"column:col_id;primaryKey"`
	Name        string    `json:"name" gorm:"column:col_name"`
	Type        string    `json:"type" gorm:"column:col_key"` // col_key 存储类型代码
	Status      string    `json:"status" gorm:"column:col_status"`
	Remark      string    `json:"remark" gorm:"column:col_remark"`
	CreatedBy   string    `json:"createdBy" gorm:"column:col_created_by_user_id"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:col_created_date"`
	UpdatedBy   string    `json:"updatedBy" gorm:"column:col_updated_by_user_id"`
	UpdatedDate time.Time `json:"updatedDate" gorm:"column:col_updated_date"`
}

func (DictType) TableName() string {
	return "sys_dict_tbl"
}

// DictData 对应 sys_dict_data_tbl
type DictData struct {
	ID          string    `json:"id" gorm:"column:col_id;primaryKey"`
	DictType    string    `json:"dictType" gorm:"column:col_key"` // col_key 存储类型代码
	Sort        int       `json:"sort" gorm:"column:col_sort"`
	Label       string    `json:"label" gorm:"column:col_label"`
	Value       string    `json:"value" gorm:"column:col_value"`
	CSSClass    string    `json:"cssClass" gorm:"column:col_css_class"`
	ListClass   string    `json:"listClass" gorm:"column:col_list_class"`
	IsDefault   int       `json:"isDefault" gorm:"column:col_default"` // int(1)
	Status      string    `json:"status" gorm:"column:col_status"`
	Remark      string    `json:"remark" gorm:"column:col_remark"`
	CreatedBy   string    `json:"createdBy" gorm:"column:col_created_by_user_id"`
	CreatedDate time.Time `json:"createdDate" gorm:"column:col_created_date"`
	UpdatedBy   string    `json:"updatedBy" gorm:"column:col_updated_by_user_id"`
	UpdatedDate time.Time `json:"updatedDate" gorm:"column:col_updated_date"`
}

func (DictData) TableName() string {
	return "sys_dict_data_tbl"
}
