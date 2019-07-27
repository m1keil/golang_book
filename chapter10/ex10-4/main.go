package main

/*
Construct a tool that reports the set of all packages in the workspace that transitively depend on the packages
specified by the arguments. Hint: you will need to run go list twice, once for the initial packages and once for all
packages. You may want to parse its JSON output using the encoding/json package.
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	var imports []string
	for _, v := range os.Args[1:] {
		pkgs, err := golist(v)
		if err != nil {
			fmt.Println("package not found", v)
			os.Exit(1)
		}

		for _, p := range pkgs {
			imports = append(imports, p.ImportPath)
		}
	}

	all, err := golist("...")
	if err != nil {
		fmt.Println("unable to list all packages")
		os.Exit(1)
	}

	for _, im := range imports {
		for _, a := range all {
			for _, ad := range a.Deps {
				if ad == im {
					fmt.Printf("%v -> %v\n", a.ImportPath, im)
					break
				}
			}
		}
	}
}

type Package struct {
	Name, ImportPath string
	Deps []string
}

func golist(name string) ([]Package, error) {
	var packages []Package
	cmd := exec.Command("go", "list", "-json", name)

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(bytes.NewReader(output))

	for {
		var p Package

		err := dec.Decode(&p)
		if err == io.EOF {
			break
		}

		if err != nil  {
			return nil, err
		}

		packages = append(packages, p)
	}


	return packages, nil
}