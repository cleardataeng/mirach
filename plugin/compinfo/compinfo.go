package compinfo

import (
	"encoding/json"

	"github.com/cleardataeng/mirach/plugin"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
)

// Exceptions is a list of strings containing error strings that are expected
// in conditions are are okay and for which operation should continue.
var Exceptions = []string{
	docker.ErrDockerNotAvailable.Error(),
}

// Docker contains information about Docker containers.
type Docker struct {
	IDs  []string `json:"container_ids"`
	Stat []docker.CgroupDockerStat
}

// Load contains information about load on a machine.
type Load struct {
	Avg  *load.AvgStat  `json:"average"`
	Misc *load.MiscStat `json:"misc"`
}

// Sys contains general information about a machine.
// This is intended for data that is mostly static. It will be checked more than
// once but not frequently.
type Sys struct {
	Host *host.InfoStat `json:"host"`
	CPUs []cpu.InfoStat `json:"cpus"`
}

var (
	d = new(Docker)
	l = new(Load)
	s = new(Sys)
)

// GetInfo retrieves information about Docker containers and populates the Docker
// struct with this data.
func (g *Docker) GetInfo() {
	var err error
	g.IDs, err = docker.GetDockerIDList()
	if err != nil {
		panic(plugin.ExceptionOrError(err, Exceptions))
	}
	g.Stat, err = docker.GetDockerStat()
	if err != nil {
		panic(plugin.ExceptionOrError(err, Exceptions))
	}
}

func (g *Docker) String() string {
	s, _ := json.Marshal(g)
	return string(s)
}

// GetDockerInfo will update docker information and return the object.
// This is a helper function that shortens:
//     d := new(compinfo.Docker)
//     d.GetInfo()
// to:
//     d := compinfo.GetDockerInfo()
func GetDockerInfo() *Docker {
	d.GetInfo()
	return d
}

// GetDockerString will update docker information and return the string.
// This is a helper function that shortens:
//     d := new(compinfo.Docker)
//     d.GetInfo()
//     json := d.String()
// to:
//     json := compinfo.GetDockerString()
func GetDockerString() string {
	d.GetInfo()
	return d.String()
}

// GetInfo retrieves information about system load and populates the Load struct
// with this data.
func (g *Load) GetInfo() {
	var err error
	g.Avg, err = load.Avg()
	if err != nil {
		panic(err)
	}
	g.Misc, err = load.Misc()
	if err != nil {
		panic(err)
	}
}

func (g *Load) String() string {
	s, _ := json.Marshal(g)
	return string(s)
}

// GetLoadInfo will update load information and return the object.
// This is a helper function that shortens:
//     l := new(compinfo.Load)
//     l.GetInfo()
// to:
//     l := compinfo.GetLoadInfo()
func GetLoadInfo() *Load {
	l.GetInfo()
	return l
}

// GetLoadString will update load information and return the string.
// This is a helper function that shortens:
//     l := new(compinfo.Load)
//     l.GetInfo()
//     json := l.String()
// to:
//     json := compinfo.GetLoadString()
func GetLoadString() string {
	l.GetInfo()
	return l.String()
}

// GetInfo retrieves general information about a system and populates the Load
// struct with this data.
func (g *Sys) GetInfo() {
	var err error
	g.Host, err = host.Info()
	if err != nil {
		panic(err)
	}
	g.CPUs, err = cpu.Info()
	if err != nil {
		panic(err)
	}
}

func (g *Sys) String() string {
	s, _ := json.Marshal(g)
	return string(s)
}

// GetSysInfo will update system information and return the object.
// This is a helper function that shortens:
//     sys := new(compinfo.Sys)
//     sys.GetInfo()
// to:
//     sys := compinfo.GetSysInfo()
func GetSysInfo() *Sys {
	s.GetInfo()
	return s
}

// GetSysString will update system information and return the string.
// This is a helper function that shortens:
//     sys := new(compinfo.Sys)
//     sys.GetInfo()
//     json := sys.String()
// to:
//     json := compinfo.GetSysString()
func GetSysString() string {
	s.GetInfo()
	return s.String()
}
