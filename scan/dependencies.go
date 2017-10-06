package scan

import "github.com/wesovilabs/glide-graph/core"

type Dependency struct {
	repository 	string
	version 	string
}

func (prj *core.Project)GetDependencies() ([]Dependency,error){

}