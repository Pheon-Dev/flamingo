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

func vip() *viper.Viper {

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

	return vp
}

func editor(path string) tea.Cmd {
	vp := vip()
	flags := vp.GetString("flags")
	editor := vp.GetString("editor")
	// editor := os.Getenv("EDITOR")
	pre_run := vp.GetString("pre-run")

	if editor == "" {
		editor = "vim"
	}

	c := exec.Command("bash", "-c", pre_run+" cd "+path+" 2>/dev/null && "+editor+" "+flags+" || "+editor+" "+flags+" "+path)
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
    vp := vip()
    quit_key := vp.GetString("quit-key")
    select_key := vp.GetString("select-key")
		switch msg.String() {
		case "q", "escape", quit_key:
			// tea.ClearScreen()
			return m, tea.Quit
		case " ", "enter", select_key:
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

			vp := vip()

			title := vp.GetString("title")
			statusbar := vp.GetBool("status-bar")
			filtering := vp.GetBool("filtering")
			ps := vp.Get("projects")

			projects := []list.Item{}
			for _, p := range ps.([]interface{}) {
				projects = append(
					projects,
					item{
						title:       string(p.(map[string]interface{})["title"].(string)),
						description: string(p.(map[string]interface{})["description"].(string)),
					},
				)
			}

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
