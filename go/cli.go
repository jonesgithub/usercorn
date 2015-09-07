package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fs := flag.NewFlagSet("cli", flag.ExitOnError)
	verbose := fs.Bool("v", false, "verbose output")
	strace := fs.Bool("strace", false, "trace syscalls")
	mtrace := fs.Bool("mtrace", false, "trace memory access")
	etrace := fs.Bool("etrace", false, "trace execution")
	rtrace := fs.Bool("rtrace", false, "trace register modification")
	prefix := fs.String("prefix", "", "library load prefix")
	fs.Usage = func() {
		fmt.Printf("Usage: %s [options] <exe> [args...]\n", os.Args[0])
		fs.PrintDefaults()
	}
	fs.Parse(os.Args[1:])
	args := fs.Args()
	if len(args) < 1 {
		fs.Usage()
		os.Exit(1)
	}
	absPrefix := ""
	var err error
	if *prefix != "" {
		absPrefix, err = filepath.Abs(*prefix)
		if err != nil {
			log.Fatal(err)
		}
	}
	corn, err := NewUsercorn(args[0], absPrefix)
	if err != nil {
		log.Fatal(err)
	}
	corn.Verbose = *verbose
	corn.TraceSys = *strace
	corn.TraceMem = *mtrace
	corn.TraceReg = *rtrace
	corn.TraceExec = *etrace
	err = corn.Run(args...)
	if err != nil {
		log.Fatal(err)
	}
}
