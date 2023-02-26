package main

import (
	utils "example/project/mypackage"
	"regexp"
)

func main() {
	reStr := "Te ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr)
	utils.Pl(match)

	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")
	utils.Pl("Match String:", r.MatchString(reStr2))
	utils.Pl("Find String:", r.FindString(reStr2))
	utils.Pl("Index:", r.FindStringIndex(reStr2))
	utils.Pl("All Matches:", r.FindAllString(reStr2, -1))
	utils.Pl("First 2:", r.FindAllString(reStr2, 2))
	utils.Pl("All Submach Index:", r.FindAllStringSubmatchIndex(reStr2, -1))
	utils.Pl("Replace all maches with 'Dog':", r.ReplaceAllString(reStr2, "Dog"))

}
