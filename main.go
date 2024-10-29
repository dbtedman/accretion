package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/dbtedman/accretion/cmd"
	"github.com/dbtedman/accretion/config"
	"github.com/dbtedman/accretion/interceptor"
	"os"
	"os/signal"
	"syscall"
)

const ErrorResult = 1
const SuccessResult = 0

func main() {
	bootstrap(run, performCleanup)
}

func run(errorCh *chan error) {
	log.Info(config.Name)

	var proxyServer interceptor.Proxy = &interceptor.ServerProxy{}

	err := cmd.RootCommand(errorCh, &proxyServer).Execute()

	if err != nil {
		*errorCh <- err
	} else {
		*errorCh <- nil
	}
}

func performCleanup(err error) {
	// Cleanup and resources used by this application on close.
}

func bootstrap(run func(errorCh *chan error), cleanup func(err error)) {
	signalsCh := make(chan os.Signal, 1)
	errorCh := make(chan error)
	var resultErr error

	signal.Notify(signalsCh, os.Interrupt, syscall.SIGTERM)

	defer func() {
		cleanup(resultErr)

		if resultErr != nil {
			fmt.Println(resultErr)
			os.Exit(ErrorResult)
		}

		os.Exit(SuccessResult)
	}()

	go func() {
		run(&errorCh)
	}()

	select {
	case <-signalsCh:
	case resultErr = <-errorCh:
	}
}
