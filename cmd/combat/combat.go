package combat

import (
	"big-battle-bits/bf"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	StepCombat.PersistentFlags().StringVar(&inputMapFile, "input-file", "map.png", "The file to read the map from")
	StepCombat.PersistentFlags().StringVar(&outputMapFile, "output-file", "map.png", "The file to write the map to")
	StepCombat.PersistentFlags().StringVar(&ordersFile, "orders-file", "data/orders.json", "The file to read prioritisations from")
	StepCombat.PersistentFlags().IntVar(&steps, "steps", 5, "Number of steps to run")

	StepCombat.MarkPersistentFlagRequired("input-file")
	StepCombat.MarkPersistentFlagRequired("output-file")
}

func readOrders(filename string) (bf.Orders, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var orders bf.Orders
	err = json.Unmarshal(b, &orders)
	return orders, err
}

func readTeams(teamsFile string) (bf.TeamColors, error) {
	f, err := os.Open(teamsFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var teams bf.TeamColors
	err = json.Unmarshal(b, &teams)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

var (
	inputMapFile, outputMapFile string
	ordersFile                  string
	steps                       int
	StepCombat                  = &cobra.Command{
		Use:   "fight",
		Short: "Steps combat forward on the map",
		RunE: func(cmd *cobra.Command, args []string) error {
			// read available teams
			teamsFile, err := cmd.Flags().GetString("teams-file")
			if err != nil {
				return err
			}
			teams, err := readTeams(teamsFile)
			orders, err := readOrders(ordersFile)
			if err != nil {
				return err
			}
			bg, err := bf.ReadFromFile(inputMapFile)
			for i := 0; i < steps; i++ {
				err := bf.StepCombat(bg, orders, teams)
				if err != nil {
					return err
				}
			}

			return bf.WriteTo(bg, "dummy.png")
		},
	}
)
