package services

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/src/tools"
)

type FembedResp struct {
	Data []struct {
		Url     string `json:"file"`
		Quality string `json:"label"`
	} `json:"data"`
}

func Fembed(u string) string {
	u = "https://fembed-hd.com/api/source/" +
		u[(strings.LastIndex(u, "/")+1):]

	req := tools.Guard(http.NewRequest("POST", u, nil))
	req.Header.Set("x-requested-with", "XMLHttpRequest")

	client := &http.Client{}

	var l = &FembedResp{}

	re := tools.Guard(client.Do(req))
	defer re.Body.Close()

	json.NewDecoder(re.Body).Decode(&l)

	return l.Data[(len(l.Data) - 1)].Url
}
