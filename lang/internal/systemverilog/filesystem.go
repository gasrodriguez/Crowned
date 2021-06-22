package systemverilog

import "github.com/gasrodriguez/crowned/internal/util"

const (
	extensionV   = ".v"
	extensionVH  = ".vh"
	extensionSV  = ".sv"
	extensionSVH = ".svh"
)

type files struct {
	compileExtensions *util.Set
	includeExtensions *util.Set
	compileFiles      *util.Set
	includeFiles      *util.Set
}

func NewFiles() *files {
	o := files{}
	o.compileExtensions = util.NewSet()
	o.compileExtensions.Add(extensionV)
	o.compileExtensions.Add(extensionSV)
	o.includeExtensions = util.NewSet()
	o.includeExtensions.Add(extensionVH)
	o.includeExtensions.Add(extensionSVH)
	return &o
}

func (o *files) ScanWorkspace(cwd string, filters []string) {
	o.compileFiles = util.FileSet(cwd, o.compileExtensions, filters)
	o.includeFiles = util.FileSet(cwd, o.includeExtensions, filters)
}

func (o *files) Compile() []string {
	return o.compileFiles.Keys()
}

func (o *files) Include() []string {
	return util.DirSet(o.includeFiles).Keys()
}
