package main

import (
	"big-battle-bits/cmd/combat"
	"big-battle-bits/cmd/gen"
	"big-battle-bits/cmd/place"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	teamsFile string
	rootCmd   = &cobra.Command{
		Use:   "blottle",
		Short: "Battlefield simulator for fun :)",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Welcome to blottle!")
			return cmd.Usage()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&teamsFile, "teams-file", "data/teams.json", "The team configuration data")
	rootCmd.AddCommand(gen.MapGen, place.PlaceTeam, combat.StepCombat)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
