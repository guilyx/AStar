//args: -Eexhaustive
//config_path: testdata/configs/exhaustive.yml
package testdata

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func processDirectionDefault(d Direction) {
	switch d {
	case North, South:
	default:
	}
}
