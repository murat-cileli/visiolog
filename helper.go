package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func getAppDataDir() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		catch(errors.New("HOME environment variable not set."))
	}

	appDataDir := filepath.Join(homeDir, ".local", "share", "visiolog")
	if _, err := os.Stat(appDataDir); os.IsNotExist(err) {
		catch(os.MkdirAll(appDataDir, 0700))
	}

	capturesDir := filepath.Join(appDataDir, "captures")
	if _, err := os.Stat(capturesDir); os.IsNotExist(err) {
		catch(os.Mkdir(capturesDir, 0700))
	}

	return appDataDir
}

func getCaptureSubDirsFromCaptureFileName(captureFileName string) string {
	subDirs := strings.Split(captureFileName, "-")
	subDirs = subDirs[:len(subDirs)-1]
	fullPath := filepath.Join(appDataDir, "captures", strings.Join(subDirs, string(os.PathSeparator)))
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		catch(os.MkdirAll(fullPath, 0700))
	}

	return fullPath
}
