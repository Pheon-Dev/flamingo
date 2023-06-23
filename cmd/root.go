package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	docStyle      = lipgloss.NewStyle().Padding(1, 2)
	quitTextStyle = lipgloss.NewStyle().Padding(1, 2)
	itemStyle     = lipgloss.NewStyle().PaddingLeft(1)
	titleStyle    = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#21222c")).
			Background(lipgloss.Color("#ff79c6")).
			Padding(0, 1)
)

type item struct {
	title string

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

func terminalPopup() tea.Cmd {
	c := exec.Command("bash", "-c", "tmux popup")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return tea.ExecProcess(c, func(err error) tea.Msg { return editorFinishedMessage{err} })
}

func editor(path string) tea.Cmd {
	ed := os.Getenv("EDITOR")

	if ed == "" {
		ed = "vim"
	}

	c := exec.Command("bash", "-c", "clear && cd "+path+" && "+ed)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return tea.ExecProcess(c, func(err error) tea.Msg { return editorFinishedMessage{err} })
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
		case "q", "escape", "h":
			tea.ClearScreen()
			return m, tea.Quit
		case "`":
			return m, terminalPopup()
		case " ", "enter", "l":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i.title)
				m.description = string(i.description)
			}
			return m, editor(m.description)
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

	// rootCmd.AddCommand(versionCmd)
}

var (
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:   "flamingo",
		Short: "Swift Configurations & Projects File Navigator",
		Long:  `Switch smoothly between different file configurations and projects without ever needing to cd into each individual file location`,
		Run: func(cmd *cobra.Command, args []string) {
			projects := []list.Item{
				item{title: "  nvim", description: "$HOME/.config/nvim"},
				item{title: "  zellij", description: "$HOME/.config/zellij"},
				// item{title: "pigeon", description: "$HOME/Documents/Neovim/pigeon"},
				// item{title: "manta-api", description: "$HOME/Documents/Rust/manta-api"},
				// item{title: "manta-wallet", description: "$HOME/Documents/Rust/manta-wallet"},
				// item{title: "helix", description: "$HOME/.config/helix"},
				item{title: "  alacritty", description: "$HOME/.config/alacritty"},
				item{title: "  joshuto", description: "$HOME/.config/joshuto"},
				item{title: "  flamingo", description: "$HOME/go/src/github.com/Pheon-Dev/flamingo"},
				// item{title:       "hms", description: "$HOME/Documents/NextJS/App/devlen/apps/hms",},
				// item{title:       "devlen", description: "$HOME/Documents/NextJS/App/devlen/apps/devlen",},
				// item{title:       "hornet", description: "$HOME/Documents/go/src/github.com/Pheon-Dev/hornet",},
				// item{title:       "zap", description: "$HOME/Documents/go/src/github.com/Pheon-Dev/zap",},
				item{title: "  dwm", description: "$HOME/.config/dwm"},
				// item{title: "st Simple Terminal", description: "$HOME/.config/arco-st"},
				// item{title: "dwmbar", description: "$HOME/.config/dwmbar"},
				item{title: "  zsh", description: "$HOME/.config/zsh"},
				// item{title: "dmenu", description: "$HOME/.config/dmenu"},
				item{title: "  btop", description: "$HOME/.config/btop"},
				item{title: "  picom", description: "$HOME/.config/picom"},
				item{title: "  rofi", description: "$HOME/.config/rofi"},
				// item{title: "tmux", description: "$HOME/.tmux"},
				item{title: "  lazygit", description: "$HOME/.config/lazygit"},
				// item{
				// 	title:       "ranger",
				// 	description: "$HOME/.config/ranger",
				// },
				// item{
				// 	title:       "fm file manager",
				// 	description: "$HOME/.config/fm",
				// },
				// item{title: "moc", description: ".moc"},
				// item{
				// 	title:       "p app",
				// 	description: "$HOME/Documents/go/src/github.com/Pheon-Dev/p",
				// },
				// item{
				// 	title:       "neovim",
				// 	description: "$HOME/Documents/Neovim",
				// },
				// item{
				// 	title:       "class",
				// 	description: "$HOME/Documents/CMT",
				// },
				// item{
				// 	title:       "go",
				// 	description: "$HOME/Documents/go/src/github.com/Pheon-Dev",
				// },
				// item{
				// 	title:       "bubbletea",
				// 	description: "$HOME/Documents/go/git/bubbletea/examples",
				// },
				// item{
				// 	title:       "go apps",
				// 	description: "$HOME/Documents/go/git",
				// },
				item{title: "  starship", description: "$HOME/.config/starship"},
				// item{title: "rust", description: "$HOME/Documents/Rust/book"},
				// item{title: "m-pesa", description: "$HOME/Documents/NextJS/App/m-pesa"},
				// item{title: "destiny", description: "$HOME/Documents/NextJS/App/destiny-credit"},
				// item{title:       "typescript", description: "$HOME/Documents/NextJS/App",},
			}

			vp := viper.New()
			vp.SetConfigName("config")
			vp.SetConfigType("yaml")
			home, homeErr := os.UserHomeDir()
			cobra.CheckErr(homeErr)

			vp.AddConfigPath(home + "/.config/flamingo")

			err := vp.ReadInConfig()
			if err != nil {
				fmt.Println(err)
			}

			title := vp.GetString("title")
			statusbar := vp.GetBool("status-bar")
			filtering := vp.GetBool("filtering")
			// workspace := vp.Get("projects")

			l := list.New(projects, list.NewDefaultDelegate(), 0, 0)
			l.SetShowStatusBar(statusbar)
			l.SetFilteringEnabled(filtering)
			l.Styles.Title = titleStyle
			l.Title = title

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
