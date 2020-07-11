package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/imgi/imgi"
	log "github.com/sirupsen/logrus"
)

func main() {
	opts, err := parseCmdArgs()
	if err != nil {
		exitWithError(err)
	}
	config, err := imgi.LoadConfig(opts.config)
	if err != nil {
		exitWithError(err)
	}
	if err := configLogging(config); err != nil {
		exitWithError(err)
	}

	srv := imgi.NewServer(config)
	srv.Start()
}

const usage = `Usage: imgi [options]

Options:
`

func init() {
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), usage)
		flag.PrintDefaults()
	}
}

type options struct {
	config string
}

func exitWithError(err error) {
	fmt.Fprint(os.Stderr, err.Error())
	os.Exit(1)
}

func parseCmdArgs() (*options, error) {
	config := flag.String("c", "imgi.conf", "Configuration file")
	help := flag.Bool("h", false, "Show help")
	version := flag.Bool("v", false, "Show version")
	flag.Parse()

	opts := &options{
		config: *config,
	}

	if *help {
		flag.Usage()
		os.Exit(2)
	}
	if *version {
		fmt.Println("imgi version ", imgi.Version)
		os.Exit(2)
	}
	return opts, nil
}

func configLogging(config imgi.Config) error {
	level, err := log.ParseLevel(config.Log.Level)
	if err != nil {
		return err
	}
	log.SetLevel(level)

	file := config.Log.File
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	logfile, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(logfile)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
	})
	return nil
}
