package dependencies
type Dependency interface {

  IsAlreadyInstalled() bool
  Install() bool
  Name() string

}

var ALL_DEPENDENCIES = []Dependency { 
  NewFzFDependency(),
}
