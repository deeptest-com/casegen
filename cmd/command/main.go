package main

import (
	"flag"
	"github.com/aaronchen2k/deeptest/pkg/consts"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"syscall"
)

var (
	AppVersion string
	BuildTime  string
	GoVersion  string
	GitHash    string

	Spec string

	flagSet *flag.FlagSet
)

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet("deeptest", flag.ContinueOnError)

	flagSet.StringVar(&Spec, "s", "", "")

	flagSet.BoolVar(&_consts.Verbose, "verbose", false, "")

	flagSet.Parse(os.Args[1:])

	switch os.Args[1] {
	case "help", "-h", "-help", "--help":

	default:

	}
}

func init() {
	cleanup()
}

func cleanup() {
	color.Unset()
}
