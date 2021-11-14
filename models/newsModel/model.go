package newsModel

type NewsModel struct {
	NewsId int	`gorm:"column:id"`
	NewsTitle string	`gorm:"column:title"`
	NewsContent string	`gorm:"column:content"`
}

func(this *NewsModel) TableName() string {
	return "news"
}

func New(fs ...NewsModelAttrFunc) *NewsModel {
	news := &NewsModel{}
	NewsModelAttrFuncs(fs).apply(news)
	return news
}