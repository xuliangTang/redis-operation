package newsModel

type NewsModelAttrFunc func(news *NewsModel)
type NewsModelAttrFuncs []NewsModelAttrFunc

func(this NewsModelAttrFuncs) apply(news *NewsModel)  {
	for _,f := range this {
		f(news)
	}
}

func WithNewsTitle(title string) NewsModelAttrFunc {
	return func(news *NewsModel) {
		news.NewsTitle = title
	}
}

func WithNewsContent(content string) NewsModelAttrFunc {
	return func(news *NewsModel) {
		news.NewsContent = content
	}
}