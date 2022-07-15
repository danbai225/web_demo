package core

type Pagination struct {
	Page      int64 `json:"page" form:"page"  binding:"required"`           //第x页
	PageSize  int64 `json:"page_size" form:"page_size"  binding:"required"` //一页多少条
	CountPage int64 `json:"count_page"`                                     //一共多少页
	Count     int64 `json:"count"`                                          //一共多少条
}

func (p *Pagination) GetOffset() int {
	if p.Page <= 1 {
		return 0
	}
	return int((p.PageSize * p.Page) - p.PageSize)
}
func (p *Pagination) GetLimit() int {
	return int(p.PageSize)
}
func (p *Pagination) SetCount(count int64) {
	p.Count = count
	if p.Count <= p.PageSize {
		p.CountPage = 1
	} else {
		p.CountPage = p.Count / p.PageSize
		if p.Count%p.PageSize > 0 {
			p.CountPage++
		}
	}
}
