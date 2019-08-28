package snr

import (
	"sort"
)

const maxListSize = 10

func SuggestSchool(pref string) (ret []string) {
	var res []string
	for i := 0; i < 10; i++ {
		if i >= len(pref) {
			break
		}

		res = gSchools.PrefixMembersList(pref[i:])
		if len(res) != 0 {
			break
		}
	}

	var nRes []*SchoolInfo
	for _, r := range res {
		if lst, suc := gPostInd[r]; suc {
			nRes = append(nRes, lst...)
		}
	}

	sort.Slice(nRes, func(i, j int) bool {
		s1, s2 := nRes[i], nRes[j]
		if s1.Level > s2.Level {
			return true
		} else if s1.Level < s2.Level {
			return false
		}

		return len(s1.Name) < len(s2.Name)
	})

	ret = make([]string, 0, len(nRes))
	dedup := make(map[string]bool)
	for _, r := range nRes {
		if _, suc := dedup[r.Name]; !suc {
			ret = append(ret, r.Name)
			dedup[r.Name] = true
		}

		if len(ret) >= 10 {
			break
		}
	}

	return
}
