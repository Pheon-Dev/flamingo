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
			fmt.Println("Flamingoes to town")
		},
	}
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

func Execute() error {
	return rootCmd.Execute()
}

func init() {
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

		if err := viper.ReadInConfig(); err != nil {
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
