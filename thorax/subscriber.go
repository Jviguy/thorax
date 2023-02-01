package thorax

import "strings"


type Subscriber struct {
	doc *strings.Builder
}

func (s Subscriber) Message(a ...any) {
	for _, i := range a {
		s.doc.WriteString(i.(string))
	}
}
