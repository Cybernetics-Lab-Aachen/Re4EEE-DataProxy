package wikipedia

import (
	"bufio"
	"compress/bzip2"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	log.Println("Init the Wikipedia package")

	//
	// Parse the index file and build the lookup table
	//

	lookupTableByTitles = make(map[string]wikiIndexWithoutTitle)
	lookupTableByArticleID = make(map[int64]wikiIndexWithoutArticleID)
	allTitles = make([]string, 0)

	fileIndex, errFile := os.OpenFile("/home/thorsten/enwiki-multistream-index.txt.bz2", os.O_RDONLY, 0)
	if errFile != nil {
		log.Printf("Was not possible to open the Wikipedia index file: %s", errFile)
	}

	defer fileIndex.Close()
	rawReader := bufio.NewReader(fileIndex)
	zipReader := bzip2.NewReader(rawReader)
	txtReader := bufio.NewReader(zipReader)
	lineScanner := bufio.NewScanner(txtReader)

	log.Println("Building index: Start")
	for lineScanner.Scan() {
		line := lineScanner.Text()
		elements := strings.Split(line, ":")
		if len(elements) == 3 {
			title := strings.TrimSpace(elements[2])
			offset, errOffset := strconv.Atoi(elements[0])
			if errOffset != nil {
				log.Printf("Was not able to parse the byte offset for a Wikipedia index entry: %s. Line: '%s'", errOffset, line)
				continue
			}
			articleID, errID := strconv.Atoi(elements[1])
			if errID != nil {
				log.Printf("Was not able to parse the articleId for a Wikipedia index entry: %s. Line: '%s'", errID, line)
				continue
			}

			lookupTableByTitles[title] = wikiIndexWithoutTitle{
				ByteOffset: int64(offset),
				ArticleID:  int64(articleID),
			}

			lookupTableByArticleID[int64(articleID)] = wikiIndexWithoutArticleID{
				ByteOffset: int64(offset),
				Title:      title,
			}

			allTitles = append(allTitles, title)
		}
	}

	log.Println("Building index: Done")
}
