package regexp

import (
	"fmt"
	"regexp"
)

// RegularExpressionMustMatch ...
func RegularExpressionMustMatch(source, pattern string) bool {
	regPattern := regexp.MustCompile(pattern)
	results := regPattern.FindAllString(source, 1)

	regNumberPattern := regexp.MustCompile("[0-9]+")
	numberResults := regNumberPattern.FindAllString(results[0], 1)

	fmt.Println("results: ", results)
	fmt.Println("numberResult: ", numberResults)

	return true
}
