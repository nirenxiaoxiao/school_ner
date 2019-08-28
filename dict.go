package snr

import (
	"bufio"
	"os"
	"strings"

	r "github.com/armon/go-radix"
	t "github.com/fvbock/trie"
)

type SchoolInfo struct {
	Name  string
	Level int32
}

func loadSchoolDict(path string) (ret *t.Trie, postInd map[string][]*SchoolInfo, reterr error) {
	ret = t.NewTrie()
	postInd = make(map[string][]*SchoolInfo)
	f, reterr := os.Open(path)
	if reterr != nil {
		return
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		eles := strings.Split(line, "\t")
		if eles[0] == "name_cn" {
			continue
		}

		school := eles[0]
		var post string
		if len(eles) > 1 {
			post = eles[1]
		}

		inf := &SchoolInfo{
			Name: school,
		}

		if strings.HasSuffix(school, "学") {
			inf.Level = 1
		}

		ret.Add(school + "||1")
		ret.Add(post + "||2")

		lst := postInd[post+"||2"]
		lst = append(lst, inf)
		postInd[post+"||2"] = lst

		lst = postInd[school+"||1"]
		lst = append(lst, inf)
		postInd[school+"||1"] = lst
	}

	return
}

func loadSchoolAliasDict(path string) (ret *r.Tree, reterr error) {
	ret = r.New()
	f, reterr := os.Open(path)
	if reterr != nil {
		return
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		eles := strings.Split(line, "\t")
		if len(eles) < 3 {
			continue
		}

		if eles[0] == "alias" {
			continue
		}

		alias := eles[0]
		if len(alias) == 0 {
			continue
		}
		name := eles[1]

		val, suc := ret.Get(alias)
		var lst []string
		if suc {
			lst = val.([]string)
		}
		lst = append(lst, name)
		ret.Insert(alias, lst)
		ret.Insert(name, lst)
	}

	return
}

var gSchools *t.Trie
var gSchoolAliases *r.Tree
var gPostInd map[string][]*SchoolInfo

func Initialize(path, pathAlias string) (reterr error) {
	if gSchools, gPostInd, reterr = loadSchoolDict(path); reterr != nil {
		return
	}

	if gSchoolAliases, reterr = loadSchoolAliasDict(pathAlias); reterr != nil {
		return
	}

	return
}
