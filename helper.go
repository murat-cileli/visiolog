package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type helperType struct {
	appDataDir  string
	capturesDir string
}

func (helper *helperType) catch(err error) {
	if err != nil {
		panic(err)
	}
}

func (helper *helperType) initFileSystem() {
	helper.initAppDataDir()
	helper.initCapturesDir()
}

func (helper *helperType) initAppDataDir() {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		helper.catch(errors.New("HOME environment variable not set."))
	}

	helper.appDataDir = filepath.Join(homeDir, ".local", "share", "visiolog")
	if _, err := os.Stat(helper.appDataDir); os.IsNotExist(err) {
		helper.catch(os.MkdirAll(helper.appDataDir, 0700))
	}
}

func (helper *helperType) initCapturesDir() {
	helper.capturesDir = filepath.Join(helper.appDataDir, "captures")
	if _, err := os.Stat(helper.capturesDir); os.IsNotExist(err) {
		helper.catch(os.Mkdir(helper.capturesDir, 0700))
	}
}

func (helper *helperType) getCaptureSubDirs(captureFileName string) string {
	subDirs := strings.Split(captureFileName, "-")
	subDirs = subDirs[:len(subDirs)-1]
	subDirsFullPath := filepath.Join(helper.appDataDir, "captures", strings.Join(subDirs, string(os.PathSeparator)))
	if _, err := os.Stat(subDirsFullPath); os.IsNotExist(err) {
		helper.catch(os.MkdirAll(subDirsFullPath, 0700))
	}

	return subDirsFullPath
}
