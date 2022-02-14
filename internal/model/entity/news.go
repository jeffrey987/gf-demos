package entity

// News is the golang structure for table news.
type News struct {
	Id         uint   `json:"id"       description:"新闻ID"`
	Title      string `json:"title" description:"标题"`
	Content    string `json:"content" description:"内容"`
	CreateTime string `json:"createTime" description:"创建时间"`
}
