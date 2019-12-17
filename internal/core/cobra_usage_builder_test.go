package core

import (
	"reflect"
	"testing"

	"github.com/alecthomas/assert"
)

type ServerColor string

const (
	ServerColorBlue  = "blue"
	ServerColorRed   = "red"
	ServerColorPink  = "pink"
	ServerColorGreen = "green"
)

type instanceListServerArgs struct {
	Name              *string
	Required          *string
	Age               *int32
	Color             *ServerColor
	MultiWordArg      *string
	RootVolume        *instanceListServerArgsVolume
	AdditionalVolumes *[]instanceListServerArgsVolume
}

type instanceListServerArgsVolume struct {
	Name *string
}

func Test_buildUsageArgs(t *testing.T) {
	want := `  [name]                            Filter all servers who contains this name
  required                          Useless but required
  [age=1]                           Filter by age
  [color=red]                       Filter by color (blue | red | pink | green)
  [multi-word-arg]                  Useless multi word argument
  [root-volume.name]                Root volume name
  [additional-volumes.{idx}.name]   Additional volume name`

	got := buildUsageArgs(&Command{
		ArgsType: reflect.TypeOf(instanceListServerArgs{}),
		ArgSpecs: ArgSpecs{
			{
				Name:  "name",
				Short: "Filter all servers who contains this name",
			},
			{
				Name:     "required",
				Short:    "Useless but required",
				Required: true,
			},
			{
				Name:    "age",
				Short:   "Filter by age",
				Default: DefaultValueSetter("1"),
			},
			{
				Name:       "color",
				Short:      "Filter by color",
				Default:    DefaultValueSetter(ServerColorRed),
				EnumValues: []string{ServerColorBlue, ServerColorRed, ServerColorPink, ServerColorGreen},
			},
			{
				Name:  "multi-word-arg",
				Short: "Useless multi word argument",
			},
			{
				Name:  "root-volume.name",
				Short: "Root volume name",
			},
			{
				Name:  "additional-volumes.{idx}.name",
				Short: "Additional volume name",
			},
		},
	})

	assert.Equal(t, want, got)
}
