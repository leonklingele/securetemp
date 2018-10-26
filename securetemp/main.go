package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.leonklingele.de/securetemp"
)

func run(size int, createDir bool, quiet bool) error {
	var cleanupFunc securetemp.CleanupFunc
	if createDir {
		tmpDir, f, err := securetemp.TempDir(size)
		if err != nil {
			return err
		}
		cleanupFunc = f
		fmt.Println(tmpDir)
	} else {
		tmpFile, f, err := securetemp.TempFile(size)
		if err != nil {
			return err
		}
		cleanupFunc = f
		fmt.Println(tmpFile.Name())
	}

	printf := func(msg string) {
		if !quiet {
			fmt.Printf(msg)
		}
	}
	println := func(msg string) {
		printf(msg + "\n")
	}

	println("The RAM disk will be cleaned up once this process terminates.")
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	printf("Cleaning up..")
	defer printf(" done.\n")
	cleanupFunc()
	return nil
}

func main() {
	size := flag.Int("size", securetemp.DefaultSize, "specifies the maximum RAM disk size in byte")
	createDir := flag.Bool("d", false, "create a directory, not a file")
	quiet := flag.Bool("q", false, "do not output status messages")
	showHelp := flag.Bool("help", false, "show help and exit")
	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if err := run(*size, *createDir, *quiet); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
