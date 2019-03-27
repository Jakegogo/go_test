package sexpr

import (
	"path/filepath"
	"strings"
)

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	// ...
}

// 判断targpath是否为basepath的子目录
func IsSubDirectory(basepath, targpath string) bool {
	if basepath == targpath {
		return false
	}

	rel, err := filepath.Rel(basepath, targpath)
	if err != nil {
		return false
	}

	if rel == "." {
		return false
	}

	if strings.HasPrefix(rel, "..") {
		return false
	}

	return true
}
