// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/sh"
	"github.com/princjef/mageutil/bintool"
	"github.com/princjef/mageutil/shellcmd"
)

var Default = All

var linter = bintool.Must(bintool.New(
	"golangci-lint{{.BinExt}}",
	"1.51.1",
	"https://github.com/golangci/golangci-lint/releases/download/v{{.Version}}/golangci-lint-{{.Version}}-{{.GOOS}}-{{.GOARCH}}{{.ArchiveExt}}",
))

func Lint() error {
	if err := linter.Ensure(); err != nil {
		return err
	}

	return linter.Command(`run`).Run()
}

func Test() error {
	return shellcmd.Command(`go test -count 1 -coverprofile=coverage.txt ./...`).Run()
}

func Build() error {
	return shellcmd.Command(`xk6 build --with github.com/szkiba/xk6-mock=.`).Run()
}

func It() error {
	all, err := filepath.Glob("scripts/*.js")
	if err != nil {
		return err
	}

	for _, script := range all {
		err := shellcmd.Command("./k6 run --quiet --no-summary --no-usage-report " + script).Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func Coverage() error {
	return shellcmd.Command(`go tool cover -html=coverage.txt`).Run()
}

func glob(patterns ...string) (string, error) {
	buff := new(strings.Builder)

	for _, p := range patterns {
		m, err := filepath.Glob(p)
		if err != nil {
			return "", err
		}

		_, err = buff.WriteString(strings.Join(m, " ") + " ")
		if err != nil {
			return "", err
		}
	}

	return buff.String(), nil
}

func License() error {
	all, err := glob("*.go", "*/*.go", ".*.yml", ".gitignore", "*/.gitignore", "*.ts", "*/*ts", ".github/workflows/*")
	if err != nil {
		return err
	}

	return shellcmd.Command(
		`reuse annotate --copyright "Iván Szkiba" --merge-copyrights --license MIT --skip-unrecognised ` + all,
	).Run()
}

func Clean() error {
	sh.Rm("magefiles/dist")
	sh.Rm("magefiles/bin")
	sh.Rm("magefiles/node_modules")
	sh.Rm("magefiles/yarn.lock")
	sh.Rm("coverage.txt")
	sh.Rm("bin")

	return nil
}

func All() error {
	if err := Lint(); err != nil {
		return err
	}

	if err := Test(); err != nil {
		return err
	}

	if err := Build(); err != nil {
		return err
	}

	return It()
}

func yarn(arg string) shellcmd.Command {
	return shellcmd.Command("yarn --silent --cwd magefiles " + arg)
}

func Prepare() error {
	if err := yarn("install").Run(); err != nil {
		return err
	}

	return shellcmd.Command("pipx install reuse").Run()
}

func Doc() error {
	if err := yarn("typedoc").Run(); err != nil {
		return err
	}

	tsmd, err := yarn("concat-md  --decrease-title-levels --start-title-level-at 2 dist/docs").Output()
	if err != nil {
		fmt.Fprint(os.Stderr, string(tsmd))

		return err
	}

	hdr, err := os.ReadFile("magefiles/README-header.md")
	if err != nil {
		return err
	}

	api, err := os.ReadFile("magefiles/README-api.md")

	api = append(api, tsmd...)

	if err := os.WriteFile("docs/README.md", api, 0o644); err != nil {
		return err
	}

	out := append(hdr, tsmd...)

	if err := os.WriteFile("README.md", out, 0o644); err != nil {
		return err
	}

	return nil
}
