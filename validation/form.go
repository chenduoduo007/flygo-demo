package validation

type Id struct {
	Id int `form:"id" json:"id" binding:"required"`
}

