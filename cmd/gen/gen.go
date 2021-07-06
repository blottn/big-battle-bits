package gen

import (
	"big-battle-bits/bf"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	MapGen.PersistentFlags().IntVar(&width, "width", 256, "The width of the map to generate")
	MapGen.PersistentFlags().IntVar(&height, "height", 256, "The height of the map to generate")
	MapGen.PersistentFlags().IntVar(&seed, "seed", int(time.Now().UnixNano()), "The noise seed")
	MapGen.PersistentFlags().StringVar(&output, "output", "map.png", "The output file name")
}

var (
	width, height, seed int
	output              string
	MapGen              = &cobra.Command{
		Use:   "gen",
		Short: "Used to generate battlefields",
		Run: func(cmd *cobra.Command, args []string) {
			bg := bf.NewBattleGround(width, height, seed)
			f, _ := os.OpenFile(output, os.O_RDWR|os.O_CREATE, 0755)
			bg.Output(f)
			f.Close()
		},
	}
)
