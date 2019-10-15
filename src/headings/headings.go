package headings

import (
	"fmt"
	"regexp"
	"sort"
)

func replaceHeadingValue(input string, headingValue int) string {
	replaceString := fmt.Sprintf("<h%[1]v>$2</h%[1]v>", headingValue)

	currentHeadingPattern := fmt.Sprintf("(?m)^(#{%v}) (.*)$", headingValue)

	currentHeadingRegexp := regexp.MustCompile(currentHeadingPattern)

	return currentHeadingRegexp.ReplaceAllString(input, replaceString)
}

func regexpMatchesToDistinctHeadingValues(matches [][]string) []int {
	headingValues := make([]int, 0)

	for _, match := range matches {
		if len(match) >= 2 && sort.SearchInts(headingValues, len(match[1])) == len(headingValues) {
			headingValues = append(headingValues, len(match[1]))
		}
	}

	return headingValues
}

func standardSyntax(mdInput string) string {
	headingRegexp := regexp.MustCompile("(?m)^(#{1,6}) (.*)$")

	matches := headingRegexp.FindAllStringSubmatch(mdInput, -1)

	htmlOutput := mdInput

	headingValues := regexpMatchesToDistinctHeadingValues(matches)

	for _, headingValue := range headingValues {
		htmlOutput = replaceHeadingValue(htmlOutput, headingValue)
	}

	return htmlOutput
}

var alternateSyntaxHeadingValue = map[int]rune{
	1: '=',
	2: '-',
}

func replaceAlternateSyntax(mdInput string, headingValue int) string {
	headingAlternateSyntax := alternateSyntaxHeadingValue[headingValue]

	regexpPattern := fmt.Sprintf("(?m)^(.*)\n%v+$", string(headingAlternateSyntax))

	headingRegexp := regexp.MustCompile(regexpPattern)

	replaceString := fmt.Sprintf("<h%[1]v>$1</h%[1]v>", headingValue)

	return headingRegexp.ReplaceAllString(mdInput, replaceString)
}

func alternateSyntax(mdInput string) string {
	htmlOutput := replaceAlternateSyntax(mdInput, 1)

	return replaceAlternateSyntax(htmlOutput, 2)
}

// Map Map the Markdown input's heading into HTML elements
func Map(mdInput string) string {
	htmlOutput := standardSyntax(mdInput)
	htmlOutput = alternateSyntax(htmlOutput)
	return htmlOutput
}
