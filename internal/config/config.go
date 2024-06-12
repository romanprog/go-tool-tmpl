package config

import (
	"os"

	"github.com/apex/log"
	"github.com/shalb/cluster.dev/pkg/colors"
	"github.com/shalb/cluster.dev/pkg/logging"
)

// Version - git tag from compiller
var Version string

// BuildTimestamp - build date from compiller
var BuildTimestamp string

// ConfSpec type for global config.
type ConfSpec struct {
	WorkingDir string
	LogLevel   string
	Version    string
	TraceLog   bool
	NoColor    bool
}

// Global config for executor.
var Global ConfSpec

// InitConfig set global config values.
func InitConfig() {
	curPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %s", err.Error())
	}
	Global.WorkingDir = curPath
	Global.Version = Version
	if Global.NoColor {
		// Turn off colored output.
		colors.SetColored(false)
	}
	logging.InitLogLevel(Global.LogLevel, Global.TraceLog)
}
