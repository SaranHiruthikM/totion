package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	vaultDir string
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home directory", err)
	}

	vaultDir = fmt.Sprintf("%s%s", homeDir, "/.totion")
}

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	textarea               textarea.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "ctrl+n":
			m.createFileInputVisible = true
			return m, nil

		case "ctrl+s":
			if m.currentFile == nil {
				break
			}

			if err := m.currentFile.Truncate(0); err != nil {
				fmt.Println("cannot save the file")
				return m, nil
			}

			if _, err := m.currentFile.Seek(0, 0); err != nil {
				fmt.Println("cannot save the file")
				return m, nil
			}

			if _, err := m.currentFile.WriteString(m.textarea.Value()); err != nil {
				fmt.Println("cannot save the file")
				return m, nil
			}

			if err := m.currentFile.Close(); err != nil {
				fmt.Println("cannot save the file")
			}

			m.currentFile = nil
			m.textarea.SetValue("")
			return m, nil
		case "enter":
			if m.currentFile != nil {
				break
			}
			fileName := m.newFileInput.Value()
			if fileName != "" {
				filepath := fmt.Sprintf("%s/%s.md", vaultDir, fileName)

				if _, err := os.Stat(filepath); err == nil {
					return m, nil
				}

				f, err := os.Create(filepath)
				if err != nil {
					log.Fatalf("%v", err)
				}

				m.currentFile = f
				m.createFileInputVisible = false
				m.newFileInput.SetValue("")

			}
			return m, nil
		}
	}

	if m.createFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
	}

	if m.currentFile != nil {
		m.textarea, cmd = m.textarea.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).
		PaddingLeft(2).
		PaddingRight(2)

	welcomeMsg := style.Render("Welcome to Totion")

	help := "Ctrl+N: new file | Ctrl+L: list | Esc: back/save | Ctrl+S: save | Ctrl+C: quit"

	view := ""
	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.textarea.View()
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcomeMsg, view, help)
}

func initializeMode() model {

	err := os.MkdirAll(vaultDir, 0750)
	if err != nil {
		log.Fatal(err)
	}

	ti := textinput.New()
	ti.Placeholder = "What would like to name the file?"
	ti.Focus()
	ti.CharLimit = 156

	ta := textarea.New()
	ta.Placeholder = "Write your notes here..."
	ta.Focus()

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		textarea:               ta,
	}
}

func main() {
	p := tea.NewProgram(initializeMode(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error!!")
		os.Exit(1)
	}
}
