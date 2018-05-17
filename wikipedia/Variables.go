package wikipedia

var (
	lookupTableByTitles    map[string]wikiIndexWithoutTitle    = make(map[string]wikiIndexWithoutTitle)
	lookupTableByArticleID map[int64]wikiIndexWithoutArticleID = make(map[int64]wikiIndexWithoutArticleID)
	allTitles              []string                            = make([]string, 0)
)
