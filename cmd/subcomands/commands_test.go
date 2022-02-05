package subcomands

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestAdd(t *testing.T) {
	type args struct {
		command *cli.Command
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "empty command",
			args: args{
				command: &cli.Command{
					Name: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add(tt.args.command)

			assert.Len(t, subcommandSingleton, 1)
			assert.Equal(t, subcommandSingleton, []*cli.Command{{Name: "test"}})
		})
	}
}
