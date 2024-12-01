package version

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"runtime"
)

const versionNumber = "1.0.0"

// Filled on build time
var gitRef string
var buildDate string

// Singleton
var version *Version

type Version struct {
	Number    string
	BuildDate string
	Os        string
	Arch      string
	Sha       string
	GitRef    string
}

func (v *Version) String() string {
	return v.Number
}

func GetVersion() *Version {
	if version != nil {
		return version
	}
	executable, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("Could not get executable: %s", err.Error()))
	}
	sha, err := computeSHA256(executable)
	if err != nil {
		panic(fmt.Sprintf("Could not compute sha: %s", err.Error()))
	}
	version = &Version{
		Number:    versionNumber,
		BuildDate: buildDate,
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		Sha:       sha,
		GitRef:    gitRef,
	}
	return version
}

func computeSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
