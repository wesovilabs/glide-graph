package core

import "github.com/wesovilabs/glide-graph/scan"

type Project struct {
	Repository 	string
	Dependencies 	[]*scan.Dependency
}
