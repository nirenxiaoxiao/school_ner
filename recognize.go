package snr

import (
	"strings"
)

const maxAllowedPrefix = 6

func suffixMode(s string) (mode int32) {
	if strings.HasSuffix(s, "大学") {
		mode = 1
	} else if strings.HasSuffix(s, "学院") {
		mode = 2
	}

	return
}

func RecognizeSchool(s string) (ret []string) {
	var rt []string
	for i := 0; i < maxAllowedPrefix; i++ {
		if i >= len(s) {
			break
		}

		_, match, suc := gSchoolAliases.LongestPrefix(s[i:])
		if suc {
			rt = match.([]string)
			break
		}
	}

	sufMode := suffixMode(s)
	for _, item := range rt {
		sm := suffixMode(item)
		if sm == sufMode {
			ret = append(ret, item)
		}
	}

	return
}
