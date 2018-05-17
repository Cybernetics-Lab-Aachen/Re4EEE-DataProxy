package wikipedia

type wikiIndexWithoutTitle struct {
	ByteOffset int64
	ArticleID  int64
}

type wikiIndexWithoutArticleID struct {
	ByteOffset int64
	Title      string
}

type wikiAnswer4Handler struct {
	Error     string `json:"Error"`
	Title     string `json:"Title"`
	ArticleID string `json:"ArticleID"`
	Text      string `json:"Text"`
}

type wikiPage struct {
	Title string `xml:"title"`
	Text  string `xml:"revision>text"`
	ID    int64  `xml:"id"`
}
