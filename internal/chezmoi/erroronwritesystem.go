package chezmoi

import (
	"io/fs"
	"os/exec"

	vfs "github.com/twpayne/go-vfs/v4"
)

// An ErrorOnWriteSystem is an System that passes reads to the wrapped System
// and returns an error if it is written to.
type ErrorOnWriteSystem struct {
	system System
	err    error
}

// NewErrorOnWriteSystem returns a new ErrorOnWriteSystem that wraps fs.
func NewErrorOnWriteSystem(system System, err error) *ErrorOnWriteSystem {
	return &ErrorOnWriteSystem{
		system: system,
		err:    err,
	}
}

// Chmod implements System.Chmod.
func (s *ErrorOnWriteSystem) Chmod(name AbsPath, mode fs.FileMode) error {
	return s.err
}

// Glob implements System.Glob.
func (s *ErrorOnWriteSystem) Glob(pattern string) ([]string, error) {
	return s.system.Glob(pattern)
}

// IdempotentCmdCombinedOutput implements System.IdempotentCmdCombinedOutput.
func (s *ErrorOnWriteSystem) IdempotentCmdCombinedOutput(cmd *exec.Cmd) ([]byte, error) {
	return s.system.IdempotentCmdCombinedOutput(cmd)
}

// IdempotentCmdOutput implements System.IdempotentCmdOutput.
func (s *ErrorOnWriteSystem) IdempotentCmdOutput(cmd *exec.Cmd) ([]byte, error) {
	return s.system.IdempotentCmdOutput(cmd)
}

// Link implements System.Link.
func (s *ErrorOnWriteSystem) Link(oldname, newname AbsPath) error {
	return s.err
}

// Lstat implements System.Lstat.
func (s *ErrorOnWriteSystem) Lstat(name AbsPath) (fs.FileInfo, error) {
	return s.system.Lstat(name)
}

// Mkdir implements System.Mkdir.
func (s *ErrorOnWriteSystem) Mkdir(name AbsPath, perm fs.FileMode) error {
	return s.err
}

// RawPath implements System.RawPath.
func (s *ErrorOnWriteSystem) RawPath(path AbsPath) (AbsPath, error) {
	return s.system.RawPath(path)
}

// ReadDir implements System.ReadDir.
func (s *ErrorOnWriteSystem) ReadDir(name AbsPath) ([]fs.DirEntry, error) {
	return s.system.ReadDir(name)
}

// ReadFile implements System.ReadFile.
func (s *ErrorOnWriteSystem) ReadFile(name AbsPath) ([]byte, error) {
	return s.system.ReadFile(name)
}

// Readlink implements System.Readlink.
func (s *ErrorOnWriteSystem) Readlink(name AbsPath) (string, error) {
	return s.system.Readlink(name)
}

// RemoveAll implements System.RemoveAll.
func (s *ErrorOnWriteSystem) RemoveAll(AbsPath) error {
	return s.err
}

// Rename implements System.Rename.
func (s *ErrorOnWriteSystem) Rename(oldpath, newpath AbsPath) error {
	return s.err
}

// RunCmd implements System.RunCmd.
func (s *ErrorOnWriteSystem) RunCmd(cmd *exec.Cmd) error {
	return s.err
}

// RunIdempotentCmd implements System.RunIdempotentCmd.
func (s *ErrorOnWriteSystem) RunIdempotentCmd(cmd *exec.Cmd) error {
	return s.system.RunIdempotentCmd(cmd)
}

// RunScript implements System.RunScript.
func (s *ErrorOnWriteSystem) RunScript(scriptname RelPath, dir AbsPath, data []byte, interpreter *Interpreter) error {
	return s.err
}

// Stat implements System.Stat.
func (s *ErrorOnWriteSystem) Stat(name AbsPath) (fs.FileInfo, error) {
	return s.system.Stat(name)
}

// UnderlyingFS implements System.UnderlyingFS.
func (s *ErrorOnWriteSystem) UnderlyingFS() vfs.FS {
	return s.system.UnderlyingFS()
}

// WriteFile implements System.WriteFile.
func (s *ErrorOnWriteSystem) WriteFile(AbsPath, []byte, fs.FileMode) error {
	return s.err
}

// WriteSymlink implements System.WriteSymlink.
func (s *ErrorOnWriteSystem) WriteSymlink(string, AbsPath) error {
	return s.err
}
