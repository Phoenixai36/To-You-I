// T.U.I — To Unify Imagination
// cmd/tuicortex/main.go — Binary entrypoint
// by Phoenixai36 · Dark Phoenix · Barcelona
package main

import (
	"fmt"
	"os"

	"github.com/Phoenixai36/To-You-I/internal/ui"
	"github.com/Phoenixai36/To-You-I/internal/server"
	"github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const version = "0.1.0-dev"

func main() {
	_ = godotenv.Load()

	root := &cobra.Command{
		Use:     "tuicortex",
		Short:   "T.U.I — To Unify Imagination",
		Long:    `To:You&I — a glitch-professional HUD for AI agents, workspaces & shells.`,
		Version: version,
		RunE:    run,
	}

	root.Flags().Int("port", 7331, "Local event server port")
	root.Flags().String("host", "localhost", "Local event server host")
	root.Flags().Bool("no-server", false, "Disable local event server")

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, _ []string) error {
	port, _ := cmd.Flags().GetInt("port")
	host, _ := cmd.Flags().GetString("host")
	noServer, _ := cmd.Flags().GetBool("no-server")

	// Start local event server in background
	if !noServer {
		srv := server.New(host, port)
		go func() {
			if err := srv.Start(); err != nil {
				fmt.Fprintf(os.Stderr, "server error: %v\n", err)
			}
		}()
	}

	// Boot Bubble Tea app
	app := ui.NewApp()
	p := tea.NewProgram(
		app,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("TUI error: %w", err)
	}
	return nil
}
