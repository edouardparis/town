package constants

// Status strings
const (
	ArticleStatusDraftStr     = "draft"
	ArticleStatusPublishedStr = "published"
)

// Statuses
const (
	ArticleStatusDraft int64 = iota
	ArticleStatusPublished
)

var ArticleStatusIntToStr = map[int64]string{
	ArticleStatusDraft:     ArticleStatusDraftStr,
	ArticleStatusPublished: ArticleStatusPublishedStr,
}
