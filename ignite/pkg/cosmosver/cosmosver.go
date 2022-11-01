package cosmosver

import (
	"fmt"
	"strings"

	"github.com/blang/semver/v4"
)

const prefix = "v"

// Version represents a range of Cosmos SDK versions.
type Version struct {
	// Version is the exact sdk version string.
	Version string

	// Semantic is the parsed version.
	Semantic semver.Version
}

var (
	MaxLaunchpadVersion           = newVersion("0.39.99")
	StargateFortyVersion          = newVersion("0.40.0")
	StargateFortyFourVersion      = newVersion("0.44.0-alpha")
	StargateFortyFiveThreeVersion = newVersion("0.45.3")
)

var (
	// Versions is a list of known, sorted Cosmos-SDK versions.
	Versions = []Version{
		MaxLaunchpadVersion,
		StargateFortyVersion,
		StargateFortyFourVersion,
	}

	// Latest is the latest known version of the Cosmos-SDK.
	Latest = Versions[len(Versions)-1]
)

func newVersion(version string) Version {
	return Version{
		Version:  "v" + version,
		Semantic: semver.MustParse(version),
	}
}

// Parse parses a Cosmos-SDK version.
func Parse(version string) (v Version, err error) {
	v.Version = version

	if v.Semantic, err = semver.Parse(strings.TrimPrefix(version, prefix)); err != nil {
		return v, err
	}

	v.Family = Stargate
	if v.LTE(MaxLaunchpadVersion) {
		v.Family = Launchpad
	}

	return
}

// GTE checks if v is greater than or equal to version.
func (v Version) GTE(version Version) bool {
	return v.Semantic.GTE(version.Semantic)
}

// LT checks if v is less than version.
func (v Version) LT(version Version) bool {
	return v.Semantic.LT(version.Semantic)
}

// LTE checks if v is less than or equal to version.
func (v Version) LTE(version Version) bool {
	return v.Semantic.LTE(version.Semantic)
}

// Is checks if v is equal to version.
func (v Version) Is(version Version) bool {
	return v.Semantic.EQ(version.Semantic)
}

func (v Version) String() string {
	return fmt.Sprintf("%s - %s", v.Family, v.Version)
}

func (v Version) IsFamily(family Family) bool {
	return v.Family == family
}
