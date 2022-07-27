package main

import "reflect"

// program to print hello world
func main() {
	println(nillifyZeroValueBuildFilter(nil) == nil)
	println(nillifyZeroValueBuildFilter(&BuildFilter{}) == nil)
}

type BuildFilter struct {
	Paths        []string
	IgnoredPaths []string
}

func nillifyZeroValueBuildFilter(filter *BuildFilter) *BuildFilter {
	if filter == nil {
		return nil
	}
	if reflect.DeepEqual(*filter, BuildFilter{}) {
		return nil
	}
	return filter
	// side note: maybe we can rewrite this func while using filter.IsEmpty()
}
