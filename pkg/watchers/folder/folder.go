package folder

import (
	"os"
)

type FolderWatcher struct {
	f    string
	size int
}

func New(c *Config) *FolderWatcher {

	return &FolderWatcher{}

}

func (f *FolderWatcher) Run() error {

}
