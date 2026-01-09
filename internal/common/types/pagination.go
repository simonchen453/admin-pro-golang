package types

// PageRequest 分页请求参数
type PageRequest struct {
	PageNo   int    `form:"pageNo" json:"pageNo" binding:"min=1"`
	PageSize int    `form:"pageSize" json:"pageSize" binding:"min=1,max=100"`
	Keyword  string `form:"keyword" json:"keyword"`
}

// PageResult 分页结果
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	PageNo   int         `json:"pageNo"`
	PageSize int         `json:"pageSize"`
}

// GetOffset 计算偏移量
func (p *PageRequest) GetOffset() int {
	return (p.PageNo - 1) * p.PageSize
}

// SetDefaults 设置默认值
func (p *PageRequest) SetDefaults() {
	if p.PageNo < 1 {
		p.PageNo = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
}
