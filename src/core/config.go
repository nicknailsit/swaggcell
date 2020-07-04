package core

import (
	"net"
	"os"
	"path/filepath"
)

type config struct {
	debug bool
	ConfigFile string
	DataDir string
	ZoneDir string
	ZoneIndex bool
	LogDir string
	torLookup func(string) ([]net.IP, error)
	torDial func(string, string) (net.Conn, error)
	TorProxy string
	TorProxyUser string
	TorProxyPass string
	TorEnabled bool
	Proxy string
	ProxyUser string
	ProxyPass string
	CellLookup func(string) ([]net.IP, error)
	CellConnection func(string, string) (net.Conn, error)
	CellSec string
	UseTestNet bool
	TLSCertificate string
	AvailCellsOnStart uint32
	UseCellBreeder bool
	AddPeers []string
	ConnectPeers []string
	Listeners []string
	MaxPeersPerCell int
	dialer net.Dialer
	client func(string,string) (net.Conn, error)
	Tld string
	breedTarget string
	breedDifficulty float32
	maxTimeToBreed float32
	breedDict string
	indexGenerations bool
	targetFitnessScore float32
	defaultEncryption string
	bannedIPS []net.IP
	bannedRegions []string
	bootstrapIPS []net.IP
	commonAncestors map[string]string
	addNeurons []string
	connectNeurons map[string]string
}

func removeDuplicatedAddresses(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	seen := map[string]struct{}{}
	for _, val := range addrs {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = struct{}{}
		}
	}
	return result

}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

var (
	homeDir, _ = os.UserHomeDir()
	defaultConfigFile = filepath.Join(homeDir, "config.json")
	defaultBreedDict = filepath.Join(homeDir,"swg/dict/en_dict.json")
	defaultDataDir = filepath.Join(homeDir, "swg/data")
	defaultZoneDir = filepath.Join(homeDir, "swg/zones")
	defaultLogDir = filepath.Join(homeDir, "swg/logs")
	defaultAncestorA = filepath.Join(homeDir, "swg/.init/cella.swg")
	defaultAncestorB = filepath.Join(homeDir, "swg/.init/cellb.swg")
)

const (
	defaultDebug = true
	defaultTld = "swg"

)

func loadConfig() (*config, []string, error) {
	AncestorsMap := make(map[string]string)
	AncestorsMap["A"] = defaultAncestorA
	AncestorsMap["B"] = defaultAncestorB
	AncestorsMap["G"] = ""
	AncestorsMap["NumG"] = ""

	cfg := config{
		debug:              defaultDebug,
		DataDir:            defaultDataDir,
		ConfigFile:         defaultConfigFile,
		ZoneDir:            defaultZoneDir,
		Tld:                defaultTld,
		LogDir:             defaultLogDir,
		breedDict:          defaultBreedDict,
		commonAncestors:    AncestorsMap,
		targetFitnessScore: 0.1,
		breedDifficulty:    1,
		maxTimeToBreed:     3500,

	}
	return &cfg, []string{}, nil
}

