package templating

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func openTemplate(templateName string) []byte {
	dat, err := ioutil.ReadFile(templateName)
    check(err)
	return dat
}

func ReplaceOccurances(templateName string, items map[string]string) []byte {
	re := regexp.MustCompile(`{{\s?(([A-Za-z_][A-Za-z0-9_]*))\s?}}`)
	file := openTemplate(templateName)
	queries := re.FindAll(file, -1)

	keywords := [][]byte{}
	for  i := 0; i < len(queries); i++ { 
		singleQuery := queries[i]
		keyword := singleQuery[ 2 : len(singleQuery) - 2 ]
		trimmedKeyword := strings.Trim(string(keyword), " ")
		keywords = append(keywords, []byte(trimmedKeyword))

		file = []byte(strings.Replace(string(file), string(singleQuery), string(items[trimmedKeyword]), -1))
	}
	return file
}