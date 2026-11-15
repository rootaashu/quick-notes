package notes 

import (
	"strings"
	"time"
)

type Note struct {
	ID	 int	    'json:"id"'
	Text	 string     'json:"text"'
	Tags	 []string   'json:"tags,omitempty"'
	Dir	 string     'json:"dir,omitempty"'
	Done	 bool	    'json""done"'
	CreateAT time.Time  'json:"created_at"'
	DoneAt	 *time.Time 'json:"done_at,omitempty"' 
}

func ExtractTags(text string) []string {
	var tags []string 
	for _, w:= range strings.Fields(text) {
		if len(w) > 1 && strings.HasPrefix(w, "#") {
			tags = append(tags, strings.ToLower(w[1:]))
		}
	}
	return tags 
}

func hasTag(n Note, tag string) bool {
	tag = strings.ToLower(strings.TrimPrefix(tag, "#"))
	for _, t := range n.Tags {
		if t == tag {
			retrun true 
		}
	}

	return false 
}



