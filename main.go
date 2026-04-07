package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	vaultDir string
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home directory", err)
	}

	vaultDir = fmt.Sprintf("%s%s", homeDir, "/.totion")
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	textarea               textarea.Model
	list                   list.Model
	showingList            bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v-5)
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.createFileInputVisible {
				m.createFileInputVisible = false
			}

			if m.currentFile != nil {
				m.currentFile = nil
			}

			if m.showingList {
				if m.list.FilterState() == list.Filtering {
					break
				}
				m.showingList = false
			}
			return m, nil
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

		case "ctrl+l":
			noteList := listFiles()
			m.list.SetItems(noteList)
			m.showingList = true
			return m, nil
		case "enter":
			if m.currentFile != nil {
				break
			}

			if m.showingList {
				item, ok := m.list.SelectedItem().(item)
				if ok {
					filepath := fmt.Sprintf("%s/%s", vaultDir, item.title)
					content, err := os.ReadFile(filepath)
					if err != nil {
						log.Printf("error reading file: %v", err)
						return m, nil
					}

					m.textarea.SetValue(string(content))

					f, err := os.OpenFile(filepath, os.O_RDWR, 0644)
					if err != nil {
						log.Printf("error reading file: %v", err)
						return m, nil
					}

					m.currentFile = f
					m.showingList = false
				}

				return m, nil
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

	if m.showingList {
		m.list, cmd = m.list.Update(msg)
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

	help := "Ctrl+N: new file | Ctrl+L: list | Esc: back | Ctrl+S: save | Ctrl+C: quit"

	view := ""
	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.textarea.View()
	}

	if m.showingList {
		view = m.list.View()
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

	noteList := listFiles()
	finalist := list.New(noteList, list.NewDefaultDelegate(), 0, 0)
	finalist.Title = "All notes list"
	finalist.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("254")).
		Padding(0, 1)

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		textarea:               ta,
		list:                   finalist,
	}
}

func main() {
	p := tea.NewProgram(initializeMode(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error!!")
		os.Exit(1)
	}
}

func listFiles() []list.Item {
	items := make([]list.Item, 0)

	entries, err := os.ReadDir(vaultDir)
	if err != nil {
		log.Fatal("Error reading notes")
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			modTime := info.ModTime().Format("01-12-2006 15:04")

			items = append(items, item{
				title: entry.Name(),
				desc:  fmt.Sprintf("Modified: %s", modTime),
			})
		}
	}

	return items
}
