package pixiv

// Illust is pixiv illust data,
type Illust struct {
	IllustID       string   `json:"illustId"`
	IllustTitle    string   `json:"illustTitle"`
	IllustType     string   `json:"illustType"`
	URL            string   `json:"url"`
	Tags           []string `json:"tags"`
	UserID         string   `json:"userId"`
	UserName       string   `json:"userName"`
	UserImage      string   `json:"userImage"`
	IsBookmarkable string   `json:"isBookmarkable"`
	IsBookmarked   string   `json:"isBookmarked"`
	Width          int      `json:"width"`
	Height         int      `json:"height"`
	PageCount      int      `json:"pageCount"`
	BookmarkCount  int      `json:"bookmarkCount"`
	ResponseCount  int      `json:"responseCount"`
	Score          *IllustScore
	IllustURL      string
}

// IllustList is list of pixiv illust.
type IllustList []Illust

// SearchResult is used to unmarshal json data about pixiv illust.
type SearchResult struct {
	IllustList []Illust `json:"illustList"`
}

// IllustScore is "view" and "nice" score of pixiv illust.
type IllustScore struct {
	View int
	Nice int
}
