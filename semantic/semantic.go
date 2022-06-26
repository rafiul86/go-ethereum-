package expressions

import "fmt"

type Semantic struct {
	major, minor, patch int
}

func NewSemantic(major, minor, patch int) Semantic {
	return Semantic{
		major, minor, patch,
	}
}

func SemanticString(sv Semantic) string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

func (sv *Semantic) IncremenMajor() {
	sv.major += 1
}
func (sv *Semantic) IncremenMinor() {
	sv.minor += 1
}
func (sv *Semantic) IncremenPatch() {
	sv.patch += 1
}