package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	configFile   = flag.String("config", "", "Config file for Env For DEV.")
	printVersion = flag.Bool("version", false, "show version")
)

var (
	version  = "0.1"
	codename = "Env For Dev"
	intro    = "This is a tool for DevOps."
)

func showVersion() {
	fmt.Printf("%s %s (%s) \n", codename, version, intro)
}

func getConfig() *viper.Viper {
	config := viper.New()

	// Set custom path and name
	if *configFile != "" {
		configName := path.Base(*configFile)
		configFileExt := path.Ext(*configFile)
		configNameOnly := strings.TrimSuffix(configName, configFileExt)
		configPath := path.Dir(*configFile)
		config.SetConfigName(configNameOnly)
		config.SetConfigType(strings.TrimPrefix(configFileExt, "."))
		config.AddConfigPath(configPath)
		// Set ASSET Path and Config Path for Env For Dev
		os.Setenv("LOCATION_ASSET", configPath)
		os.Setenv("LOCATION_CONFIG", configPath)
	} else {
		// Set default config path
		config.SetConfigName("config")
		config.SetConfigType("yml")
		config.AddConfigPath(".")

	}

	if err := config.ReadInConfig(); err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}

	config.WatchConfig() // Watch the config

	return config
}

func main() {
	flag.Parse()
	showVersion()

	if *printVersion {
		return
	}

	config := getConfig()
	lastTime := time.Now()
	config.OnConfigChange(func(e fsnotify.Event) {
		// Discarding event received within a short period of time after receiving an event.
		if time.Now().After(lastTime.Add(3 * time.Second)) {
			// Hot reload function
			fmt.Println("Config file changed:", e.Name)
			// Delete old instance and trigger GC
			runtime.GC()
			// print Redis Info
			lastTime = time.Now()
		}
	})
}
