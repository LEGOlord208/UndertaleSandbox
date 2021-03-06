package main

import (
	"os/user"
	"path/filepath"
)

func defaultConfigDir(current *user.User) string {
	return filepath.Join(current.HomeDir, ".config", "UNDERTALE")
}

func defaultSteamDir(current *user.User) string {
	return filepath.Join(current.HomeDir, ".local", "share", "Steam")
}
