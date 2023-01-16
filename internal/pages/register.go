package pages

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	pb "github.com/vllvll/keepa/gen"
	"github.com/vllvll/keepa/internal/services"
)

func NewRegisterModel() LoginModel {
	m := LoginModel{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Login"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		m.inputs[i] = t
	}

	return m
}

type RegisterModel struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode textinput.CursorMode
}

func (r RegisterModel) Init() tea.Cmd {
	return textinput.Blink
}

func (r RegisterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return r, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && r.focusIndex == len(r.inputs) {
				for i := 0; i <= len(r.inputs)-1; i++ {
					if r.inputs[i].Value() == "" {
						r.inputs[i].PromptStyle = errorStyle

						return r, tea.Batch(make([]tea.Cmd, len(r.inputs))...)
					}
				}

				login := r.inputs[0].Value()
				password := r.inputs[1].Value()

				grpc, _ := services.NewGRPCSendClient()

				ctx := context.Background()
				request := pb.AuthRequest{
					Login:    login,
					Password: password,
				}

				response, err := grpc.Client.Register(ctx, &request)
				if err != nil {
					return r, tea.Quit
				}

				grpc.SetToken(response.GetToken())

				return NewMenuModel(), nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				r.focusIndex--
			} else {
				r.focusIndex++
			}

			if r.focusIndex > len(r.inputs) {
				r.focusIndex = 0
			} else if r.focusIndex < 0 {
				r.focusIndex = len(r.inputs)
			}

			cmds := make([]tea.Cmd, len(r.inputs))
			for i := 0; i <= len(r.inputs)-1; i++ {
				if i == r.focusIndex {
					// Set focused state
					cmds[i] = r.inputs[i].Focus()
					r.inputs[i].PromptStyle = focusedStyle
					r.inputs[i].TextStyle = focusedStyle
					continue
				}

				// Remove focused state
				r.inputs[i].Blur()
				r.inputs[i].PromptStyle = noStyle
				r.inputs[i].TextStyle = noStyle
			}

			return r, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := r.updateInputs(msg)

	return r, cmd
}

func (r RegisterModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(r.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range r.inputs {
		r.inputs[i], cmds[i] = r.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (r RegisterModel) View() string {
	var b strings.Builder

	for i := range r.inputs {
		b.WriteString(r.inputs[i].View())
		if i < len(r.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if r.focusIndex == len(r.inputs) {
		button = &focusedButton
	}

	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
