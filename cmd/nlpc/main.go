package main

import (
	"flag"
	"log"
)

var (
	fmod, fin, ftrg string

	mods = map[string]module{
		"languages": modLanguages{},
	}
)

func init() {
	rf := func(v *string, names []string, value, usage string) {
		for i := range names {
			flag.StringVar(v, names[i], value, usage)
		}
	}
	rf(&fmod, []string{"module", "mod", "m"}, "", "Module to compile: [languages]")
	rf(&fin, []string{"input", "in", "i"}, "", "Path to source data file")
	rf(&ftrg, []string{"target", "t"}, "", "Target file or directory")
	flag.Parse()

	if len(fmod) == 0 {
		log.Fatalln("param -module is required")
	}
	var (
		mod module
		ok  bool
		err error
	)
	if mod, ok = mods[fmod]; !ok {
		log.Fatalf("unknown module: %s\n", fmod)
	}
	if err = mod.Validate(fin, ftrg); err != nil {
		log.Fatalf("module validation failed: %s\n", err.Error())
	}
}

func main() {
	var (
		err error
		mod module
	)
	mod = mods[fmod]
	log.Printf("%s compilation started\n", fmod)
	if err = mod.Compile(fin, ftrg); err != nil {
		log.Fatalf("%s compilation failed: %s\n", fmod, err.Error())
	}
	log.Printf("%s compilation done\n", fmod)
}
