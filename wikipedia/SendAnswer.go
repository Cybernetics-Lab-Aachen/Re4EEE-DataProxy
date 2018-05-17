package wikipedia

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendAnswer(answer wikiAnswer4Handler, response http.ResponseWriter) {
	jsonBytes, _ := json.Marshal(answer)
	fmt.Fprintln(response, string(jsonBytes))
}
