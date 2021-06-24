package systemverilog

import (
	"github.com/gasrodriguez/crowned/internal/util"
	"path/filepath"
)

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
	modifiedFiles     *util.Set
}

func NewFiles() *files {
	o := files{}
	o.compileExtensions = util.NewSet(extensionV, extensionSV)
	o.includeExtensions = util.NewSet(extensionVH, extensionSVH)
	return &o
}

func (o *files) ScanWorkspace(cwd string, filters []string) {
	o.compileFiles = util.FileSet(cwd, o.compileExtensions, filters)
	o.includeFiles = util.FileSet(cwd, o.includeExtensions, filters)
}

func (o *files) CompileFiles() []string {
	return o.compileFiles.Keys()
}

func (o *files) IncludeDirs() []string {
	return util.DirSet(o.includeFiles).Keys()
}

func (o *files) AddFile(file string) {
	if o.isInclude(file) {
		o.includeFiles.Add(file)
		return
	}
	if o.isCompile(file) {
		o.compileFiles.Add(file)
	}
}

func (o *files) RemoveFile(file string) {
	if o.isInclude(file) {
		o.includeFiles.Delete(file)
		return
	}
	if o.isCompile(file) {
		o.compileFiles.Delete(file)
	}
}

func (o *files) ChangeFile(file string) {
	o.modifiedFiles.Add(file)
}

func (o *files) SavedFile(file string) {
	o.modifiedFiles.Delete(file)
}

func (o *files) isModified(file string) bool {
	return o.modifiedFiles.Contains(file)
}

func (o *files) isInclude(file string) bool {
	return o.includeExtensions.Contains(filepath.Ext(file))
}

func (o *files) isCompile(file string) bool {
	return o.includeExtensions.Contains(filepath.Ext(file))
}
