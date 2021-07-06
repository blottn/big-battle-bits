package place

import (
	"big-battle-bits/bf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	PlaceTeam.PersistentFlags().StringVar(&file, "map-file", "map.png", "The file to read the map from")
	PlaceTeam.PersistentFlags().StringVar(&name, "team-name", "", "The team name to be placed")
	PlaceTeam.PersistentFlags().IntVar(&x, "x", 0, "The x coord for placing")
	PlaceTeam.PersistentFlags().IntVar(&y, "y", 0, "The y coord for placing")
	PlaceTeam.MarkPersistentFlagRequired("x")
	PlaceTeam.MarkPersistentFlagRequired("y")
	PlaceTeam.MarkPersistentFlagRequired("team-name")
}

var (
	file      string
	name      string
	x, y      int
	PlaceTeam = &cobra.Command{
		Use:   "place",
		Short: "Used to place teams initially",
		RunE: func(cmd *cobra.Command, args []string) error {
			// read teams
			teamsFile, err := cmd.Flags().GetString("teams-file")
			if err != nil {
				return err
			}
			f, err := os.Open(teamsFile)
			if err != nil {
				return err
			}
			b, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}
			var teams map[string]bf.Team
			err = json.Unmarshal(b, &teams)
			if err != nil {
				return err
			}
			bg, err := bf.ReadFromFile(file)
			if err != nil {
				return err
			}
			team, ok := teams[name]
			if !ok {
				return fmt.Errorf("Team not found: %s", name)
			}
			bg.Add(team, x, y)
			return bf.WriteTo(bg, file)
		},
	}
)
