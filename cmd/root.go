package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type item struct {
	title       string
	description string
}

type model struct {
	list        list.Model
	choice      string
	description string
}

type editorFinishedMessage struct{ err error }

func (i item) Title() string {
	return i.title
}

func (i item) Description() string {
	return i.description
}

func (i item) FilterValue() string {
	return i.title
}

func editor(path string) tea.Cmd {
	fmt.Println("Editor")
	return nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		height, width := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-width, msg.Height-height)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "escape":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return "\n" + m.list.View()
}

func (m model) init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().
		StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/flamingo/config.yaml)")
	rootCmd.PersistentFlags().
		StringP("author", "a", "Pheon-Dev", "author name for copyright attribution")
	rootCmd.PersistentFlags().
		StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().
		Bool("viper", true, "use Viper for configuration")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Pheon-Dev <devpheon@gmail.com>")
	viper.SetDefault("license", "MIT")

	rootCmd.AddCommand(versionCmd)
}

var (
	cfgFile       string
	userLicense   string
	docStyle      = lipgloss.NewStyle().Padding(1, 2)
	quitTextStyle = lipgloss.NewStyle().Padding(1, 2)
	itemStyle     = lipgloss.NewStyle().PaddingLeft(1)
	titleStyle    = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#c0caf5")).
			Background(lipgloss.Color("#ff79c6")).
			Padding(0, 1)
	rootCmd = &cobra.Command{
		Use:   "flamingo",
		Short: "Swift Configurations & Projects File Navigator",
		Long:  `Switch smoothly between different file configurations and projects without ever needing to cd into each individual file location`,
		Run: func(cmd *cobra.Command, args []string) {
			projects := []list.Item{
				item{title: "nvim", description: "$HOME/.config/nvim"},
				item{
					title:       "flamingo",
					description: "$HOME/Documents/go/src/github.com/Pheon-Dev/flamingo",
				},
				item{title: "dwm", description: "$HOME/.config/arco-dwm"},
				item{title: "zsh", description: "$HOME/.config/zsh"},
				item{title: "dmenu", description: "$HOME/.config/dmenu"},
				item{title: "btop", description: "$HOME/.config/btop"},
				item{title: "tmux", description: "$HOME/.tmux"},
				item{
					title:       "st Simple Terminal",
					description: "$HOME/.config/arco-st",
				},
				item{
					title:       "lazygit",
					description: "$HOME/.config/lazygit",
				},
				item{
					title:       "ranger",
					description: "$HOME/.config/ranger",
				},
				item{
					title:       "fm file manager",
					description: "$HOME/.config/fm",
				},
				item{title: "moc", description: ".moc"},
				item{
					title:       "p app",
					description: "$HOME/Documents/go/src/github.com/Pheon-Dev/p",
				},
				item{
					title:       "neovim",
					description: "$HOME/Documents/Neovim",
				},
				item{
					title:       "class",
					description: "$HOME/Documents/CMT",
				},
				item{
					title:       "go",
					description: "$HOME/Documents/go/src/github.com/Pheon-Dev",
				},
				item{
					title:       "bubbletea",
					description: "$HOME/Documents/go/git/bubbletea/examples",
				},
				item{
					title:       "go apps",
					description: "$HOME/Documents/go/git",
				},
				item{
					title:       "destiny",
					description: "$HOME/Documents/NextJS/App/destiny-credit",
				},
				item{
					title:       "devlen",
					description: "$HOME/Documents/NextJS/App/devlen",
				},
				item{
					title:       "typescript",
					description: "$HOME/Documents/NextJS/App",
				},
			}

			l := list.New(projects, list.NewDefaultDelegate(), 0, 0)
			m := model{list: l}

			if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
				fmt.Println("Error Running Program: ", err)
				os.Exit(1)
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.config/flamingo")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

		viper.AutomaticEnv()

		if verr := viper.ReadInConfig(); verr != nil {
			fmt.Println("Using config file: ", viper.ConfigFileUsed())
		}

	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print flamingo version",
	Long:  `This is the latest version of flamingo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Flamingo version: v0.0.1")
	},
}
