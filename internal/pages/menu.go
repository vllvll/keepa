package pages

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func NewMenuModel() AuthModel {
	items := []list.Item{
		item{key: "bank_cards", title: "Bank cards", desc: "See all bank cards"},
		item{key: "texts", title: "Texts", desc: "See all texts"},
		item{key: "binaries", title: "Binaries", desc: "See all binaries"},
	}

	m := AuthModel{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Keepa"

	return m
}

type MenuModel struct {
	list list.Model
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		keypress := msg.String()
		if keypress == "ctrl+c" {
			return m, tea.Quit
		}

		//if keypress == "enter" {
		//	selectedItem := m.list.SelectedItem().(item)
		//	if selectedItem.key == "Authorization" {
		//		loginModel := NewLoginModel()
		//
		//		return loginModel, loginModel.Init()
		//	}
		//
		//	if selectedItem.key == "Registration" {
		//		loginModel := NewLoginModel()
		//
		//		return loginModel, loginModel.Init()
		//	}
		//}
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
