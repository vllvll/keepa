package pages

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	key, title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func NewAuthModel() AuthModel {
	items := []list.Item{
		item{key: "Authorization", title: "Log in", desc: "If already registered"},
		item{key: "Registration", title: "Register", desc: "New user"},
	}

	m := AuthModel{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Keepa"

	return m
}

type AuthModel struct {
	list list.Model
}

func (m AuthModel) Init() tea.Cmd {
	return nil
}

func (m AuthModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		keypress := msg.String()
		if keypress == "ctrl+c" {
			return m, tea.Quit
		}

		if keypress == "enter" {
			selectedItem := m.list.SelectedItem().(item)
			if selectedItem.key == "Authorization" {
				loginModel := NewLoginModel()

				return loginModel, loginModel.Init()
			}

			if selectedItem.key == "Registration" {
				loginModel := NewLoginModel()

				return loginModel, loginModel.Init()
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m AuthModel) View() string {
	return docStyle.Render(m.list.View())
}
