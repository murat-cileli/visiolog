package main

import (
	"errors"
	"os"
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

	appDataDir := homeDir + "/.local/share/visiolog"
	if _, err := os.Stat(appDataDir); os.IsNotExist(err) {
		catch(os.Mkdir(appDataDir, 0755))
	}

	capturesDir := appDataDir + "/captures"
	if _, err := os.Stat(capturesDir); os.IsNotExist(err) {
		catch(os.Mkdir(capturesDir, 0755))
	}

	return appDataDir
}
