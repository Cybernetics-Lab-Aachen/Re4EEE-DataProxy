package wikipedia

import (
	"bufio"
	"compress/bzip2"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func HandlerGetArticle(response http.ResponseWriter, request *http.Request) {

	answer := wikiAnswer4Handler{}
	request.ParseForm()
	requestedTitle := request.FormValue("title")
	requestedTitle = strings.TrimSpace(requestedTitle)

	if requestedTitle == `` {
		answer.Error = "No title were defined"
		sendAnswer(answer, response)
		return
	}

	entry, entryOK := lookupTableByTitles[requestedTitle]
	if !entryOK {
		answer.Error = "Unknown title were defined"
		sendAnswer(answer, response)
		return
	}

	answer.ArticleID = fmt.Sprintf("%d", entry.ArticleID)

	fileDump, errFile := os.OpenFile("/home/data/enwiki-multistream.xml.bz2", os.O_RDONLY, 0)
	if errFile != nil {
		answer.Error = fmt.Sprintf("Was not possible to open the Wikipedia dump: %s", errFile)
		log.Println(answer.Error)
		sendAnswer(answer, response)
		return
	}

	defer fileDump.Close()
	fileDump.Seek(entry.ByteOffset, 0)
	rawReader := bufio.NewReader(fileDump)
	zipReader := bzip2.NewReader(rawReader)
	xmlReader := xml.NewDecoder(zipReader)
	numberReadArticles := 0
	found := false
	for {
		token, _ := xmlReader.Token()
		if token == nil {
			break
		}

		if numberReadArticles > 100 {
			break
		}

		switch section := token.(type) {
		case xml.StartElement:
			name := section.Name.Local
			if name == "page" {
				numberReadArticles++
				var page wikiPage
				xmlReader.DecodeElement(&page, &section)

				if page.ID == entry.ArticleID {
					found = true
					answer.Title = page.Title
					answer.Text = page.Text
					break
				}
			}
		}
	}

	if !found {
		answer.Error = "Was not able to find the desired article in the Wikipedia dump"
	}

	sendAnswer(answer, response)
}
