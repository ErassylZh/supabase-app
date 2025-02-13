package schema

type Pagination struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}

func (p *Pagination) Validate() {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.Size == 0 {
		p.Size = 25
	}
}
