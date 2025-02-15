package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/pimmytrousers/malpedia_cli/types"
	"github.com/pimmytrousers/malpedia_cli/util"

	"github.com/spf13/cobra"
)

// getActorCmd represents the getActor command
var getActorCmd = &cobra.Command{
	Use:   "actor",
	Short: "Will return metadata about a specific actor",
	Long: `This will make 2 requests, one to check that the actor ID is valid and another
to request the metadata if the actor ID is valid

Example Usage:
- malpedia_cli actor apt28
- malpedia_cli actor apt28 --json
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("actor takes in only a single argument")
		}
		return runGetActor(apiKey, args[0])
	},
}

func init() {
	rootCmd.AddCommand(getActorCmd)
}

func runGetActor(key, actorID string) error {
	name, err := util.GetActorName(actorID, key)
	if err != nil {
		return err
	}

	endpoint := types.EndpointGetActor
	endpoint = fmt.Sprintf(endpoint, name)
	res, err := util.HttpGetQuery(types.Endpoint(endpoint), key)
	if err != nil {
		if err == types.ErrResourceNotFound {
			return err
		}
		// BUG: res is now nil
	}

	actor := &types.Actor{}
	err = json.Unmarshal(res, actor)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetAllowedColumnLengths([]int{20, 128})

	t.Style().Options.SeparateColumns = true
	t.Style().Options.DrawBorder = true
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateColumns = true
	t.Style().Options.DrawBorder = true
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Field", "Value"})

	t.AppendRow(table.Row{"Common Name", actor.Value + "\n"})

	if actor.Meta.Country != "" {
		t.AppendRow(table.Row{"Suspected Country", actor.Meta.Country + "\n"})
	}

	if actor.Meta.CfrTypeOfIncident != "" {
		t.AppendRow(table.Row{"Operation Type", actor.Meta.CfrTypeOfIncident + "\n"})
	}

	if len(actor.Meta.Synonyms) > 0 {
		for i, syn := range actor.Meta.Synonyms {
			key := fmt.Sprintf("Alias %d", i+1)
			val := syn

			if i == len(actor.Meta.Synonyms)-1 {
				val += "\n"
			}

			t.AppendRow(table.Row{key, val})
		}
	}

	actor.Description = strings.ReplaceAll(actor.Description, "\t", " ")
	actor.Description = strings.ReplaceAll(actor.Description, "\n", " ")
	actor.Description = strings.ReplaceAll(actor.Description, "\r", "")
	if actor.Description != "" {
		t.AppendRow(table.Row{"Description", actor.Description + "\n"})
	}

	if len(actor.Families) > 0 {
		i := 0
		for family := range actor.Families {
			key := fmt.Sprintf("Family %d", i+1)
			val := family
			if i == len(actor.Families)-1 {
				val += "\n"
			}
			t.AppendRow(table.Row{key, val})

			i++
		}
	}

	if len(actor.Meta.Refs) > 0 {
		for i, reference := range actor.Meta.Refs {
			key := fmt.Sprintf("Reference %d", i+1)
			val := reference
			if i == len(actor.Meta.Refs)-1 {
				val += "\n"
			}
			t.AppendRow(table.Row{key, val})
		}
	}

	t.Render()
	return nil
}
