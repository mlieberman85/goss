package system

import (
	"errors"
	"strings"

	"github.com/aelsabbahy/goss/util"
)

type BrewPackage struct {
	name      string
	versions  []string
	loaded    bool
	installed bool
}

func NewBrewPackage(name string, system *System, config util.Config) Package {
	return &BrewPackage{name: name}
}

func (p *BrewPackage) setup() {
	if p.loaded {
		return
	}
	p.loaded = true
	cmd := util.NewCommand("brew", "list", "--versions", p.name)
	if err := cmd.Run(); err != nil {
		return
	}
	p.installed = true
	p.versions = strings.Split(strings.TrimSpace(cmd.Stdout.String()), "\n")
}

func (p *BrewPackage) Name() string {
	return p.name
}

func (p *BrewPackage) Exists() (bool, error) { return p.Installed() }

func (p *BrewPackage) Installed() (bool, error) {
	p.setup()

	return p.installed, nil
}

func (p *BrewPackage) Versions() ([]string, error) {
	p.setup()
	if len(p.versions) == 0 {
		return p.versions, errors.New("Package version not found")
	}
	return p.versions, nil
}
