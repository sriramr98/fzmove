package dependencies
type Dependency interface {

  IsInstalled() bool
  Install() bool
  Name() string

}

var ALL_DEPENDENCIES = []Dependency { 
  NewFzFDependency(),
}
