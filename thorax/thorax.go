package thorax

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sirupsen/logrus"
)

type MainMenu struct {
	options        []string
	cursor         int
	selectedOption int
	optionToMenu   []tea.Model
}

func (m MainMenu) View() string {
	// The header
	menu := make([]string, 0)
	title := `███        ▄█    █▄     ▄██████▄     ▄████████    ▄████████ ▀████    ▐████▀ 
▀█████████▄   ███    ███   ███    ███   ███    ███   ███    ███   ███▌   ████▀  
▀███▀▀██   ███    ███   ███    ███   ███    ███   ███    ███    ███  ▐███    
███   ▀  ▄███▄▄▄▄███▄▄ ███    ███  ▄███▄▄▄▄██▀   ███    ███    ▀███▄███▀    
███     ▀▀███▀▀▀▀███▀  ███    ███ ▀▀███▀▀▀▀▀   ▀███████████    ████▀██▄     
███       ███    ███   ███    ███ ▀███████████   ███    ███   ▐███  ▀███    
███       ███    ███   ███    ███   ███    ███   ███    ███  ▄███     ███▄  
▄████▀     ███    █▀     ▀██████▀    ███    ███   ███    █▀  ████       ███▄` + "\n\n\n"
	menu = append(menu, title)
	// Iterate over our choices
	for i, choice := range m.options {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		// Render the row
		menu = append(menu, fmt.Sprintf("%s %s", lipgloss.NewStyle().Foreground(lipgloss.Color("#d1ff33")).Blink(true).Render(cursor), lipgloss.NewStyle().Width(24).Foreground(lipgloss.Color("#33c9dc")).Render(choice)))
	}
	// Send the UI for rendering
	return lipgloss.JoinVertical(lipgloss.Center, menu...)
}

func (m MainMenu) Init() tea.Cmd {
	return nil
}

func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor == 0 {
				m.cursor = len(m.options)
			}
			m.cursor--
		// The "down" and "j" keys move the cursor down
		case "down", "j":
			m.cursor = (m.cursor + 1) % len(m.options)
		case "enter", " ":

		}
	}
	return m, nil
}

func Start(log *logrus.Logger) {
	cdoc := strings.Builder{}
	chat.Global.Subscribe(Subscriber{doc: &cdoc})
	ldoc := strings.Builder{}
	log.SetOutput(&ldoc)
	model := MainMenu{
		options: []string{"Log", "Command", "Chat", "Player List"},
	}
	p := tea.NewProgram(&model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}
