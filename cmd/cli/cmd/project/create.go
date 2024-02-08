package project

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/do"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"kimcha/config"
	"kimcha/immu"
	"os"
	"strings"
)

type createProjectModel struct {
	masterKeyInput textinput.Model

	name string

	inj *do.Injector
}

func initialModel(inj *do.Injector) createProjectModel {
	ti := textinput.New()
	ti.Placeholder = "New project name"
	ti.Focus()
	ti.EchoMode = textinput.EchoNormal
	return createProjectModel{
		masterKeyInput: ti,
		inj:            inj,
	}
}

func (m createProjectModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m createProjectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, m.handleSubmit()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	m.masterKeyInput, cmd = m.masterKeyInput.Update(msg)

	return m, cmd
}

func (m createProjectModel) View() string {
	var b strings.Builder

	b.WriteRune('\n')
	b.WriteString("Type new project name")
	b.WriteRune('\n')
	b.WriteString(m.masterKeyInput.View())

	return b.String()
}

func (m createProjectModel) handleSubmit() tea.Cmd {
	m.name = m.masterKeyInput.Value()

	im := do.MustInvoke[immu.Manager](m.inj)

	_, err := im.CreateProject(context.Background(), m.name)
	if err != nil {
		fmt.Println(err.Error())
		return tea.Quit
	}

	return tea.Quit
}

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		config.ReadConfigFromHomeDirToViper()
		injector := do.New()
		do.Provide(injector, immu.NewDatabase)
		do.Provide(injector, immu.NewManager)

		p := tea.NewProgram(initialModel(injector))
		if _, err := p.Run(); err != nil {
			fmt.Printf("error %s", err)
			os.Exit(1)
		}
	},
	GroupID: "project",
}
