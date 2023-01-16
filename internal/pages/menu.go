package pages

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// WindowSize store the size of the terminal window
	WindowSize tea.WindowSizeMsg
)

// DocStyle styling for viewports
var DocStyle = lipgloss.NewStyle().Margin(0, 2)

func NewMenuModel() MenuModel {
	items := []list.Item{
		item{key: "bank_cards", title: "Bank cards", desc: "See all bank cards"},
		item{key: "texts", title: "Texts", desc: "See all texts"},
		item{key: "binaries", title: "Binaries", desc: "See all binaries"},
	}

	m := MenuModel{list: list.New(items, list.NewDefaultDelegate(), 20, 20)}
	m.list.Title = "Keepa"

	if WindowSize.Height != 0 {
		top, right, bottom, left := DocStyle.GetMargin()
		m.list.SetSize(WindowSize.Width-left-right, WindowSize.Height-top-bottom-1)
	}

	return m
}

type MenuModel struct {
	list list.Model
}

func (m MenuModel) Init() tea.Cmd {
	m.list.View()

	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		keypress := msg.String()
		if keypress == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m MenuModel) View() string {
	return docStyle.Render(m.list.View())
}
