package ui

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"

	"go.uber.org/fx"
)

func main() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

type tickMsg struct{}
type errMsg error

type model struct {
	questionArea viewport.Model
	textarea     textarea.Model
	err          error
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	style := lipgloss.NewStyle()
	style.Underline(true)
	return fmt.Sprintf(
		"Type This Story.\n\n%s\n\n%s\n\n%s",
		m.questionArea.View(),
		style.Render(m.textarea.View()),
		"(ctrl+c to quit)",
	) + "\n\n"
}

func initialModel() model {
	ti := textarea.New()
	ti.ShowLineNumbers = false
	ti.FocusedStyle.Prompt.BorderLeft(false)
	ti.FocusedStyle.CursorLine = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ff0000"))
	// ti := textinput.New()
	// ti.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ti.Placeholder = "Once upon a time..."
	ti.Focus()

	t2 := viewport.New(80, 10)
	t2.SetContent("hi, this is a test model")

	return model{
		questionArea: t2,
		textarea:     ti,
		err:          nil,
	}
}

var Module = fx.Module(
	"ui",
	fx.Provide(
		func() *tea.Program {
			return tea.NewProgram(initialModel(), tea.WithAltScreen())
		},
	),
	fx.Invoke(
		func(lf fx.Lifecycle, program *tea.Program) {
			lf.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						go program.Start()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						program.Quit()
						return nil
					},
				},
			)
		},
	),
	// fx.Invoke(
	// 	func(d domain.Domain) {
	// 		ctx := context.TODO()
	// 		resp, err := d.GetUserById(ctx, "hi")
	//
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	//
	// 		fmt.Println(resp)
	// 	},
	// ),
)
