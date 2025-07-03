package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TestItem struct {
	Path     string
	Name     string
	IsDir    bool
	HasTests bool
}

type TestResult struct {
	Path    string
	Success bool
	Output  string
	Error   string
}

type model struct {
	tests         []TestItem
	cursor        int
	currentDir    string
	loading       bool
	testResult    *TestResult
	showingResult bool
	errorMsg      string
}

type testDiscoveredMsg []TestItem
type testResultMsg TestResult
type errorMsg string

var (
	titleStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
	itemStyle      = lipgloss.NewStyle().PaddingLeft(1)
	selectedStyle  = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
	successStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Bold(true)
	errorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	directoryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("33"))
	testFileStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("76"))
)

func initialModel() model {
	return model{
		tests:      []TestItem{},
		cursor:     0,
		currentDir: ".",
		loading:    true,
	}
}

func (m model) Init() tea.Cmd {
	return discoverTests(m.currentDir)
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.showingResult {
			switch msg.String() {
			case "q", "ctrl+c", "esc":
				m.showingResult = false
				m.testResult = nil
				return m, nil
			}
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.tests)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.tests) > 0 {
				selected := m.tests[m.cursor]
				if selected.IsDir {
					m.currentDir = selected.Path
					m.loading = true
					m.cursor = 0
					return m, discoverTests(selected.Path)
				} else if selected.HasTests {
					m.loading = true
					return m, runTest(selected.Path)
				}
			}
		case "r":
			m.loading = true
			m.cursor = 0
			return m, discoverTests(m.currentDir)
		case "backspace":
			if m.currentDir != "." {
				parent := filepath.Dir(m.currentDir)
				if parent == "." || parent == "/" {
					parent = "."
				}
				m.currentDir = parent
				m.loading = true
				m.cursor = 0
				return m, discoverTests(parent)
			}
		}

	case testDiscoveredMsg:
		m.tests = []TestItem(msg)
		m.loading = false
		m.errorMsg = ""

	case testResultMsg:
		result := TestResult(msg)
		m.testResult = &result
		m.showingResult = true
		m.loading = false

	case errorMsg:
		m.errorMsg = string(msg)
		m.loading = false
	}

	return m, nil
}

func (m model) View() string {
	if m.showingResult && m.testResult != nil {
		return m.renderTestResult()
	}

	s := titleStyle.Render("Go Test Runner") + "\n\n"
	s += fmt.Sprintf("Directory: %s\n\n", m.currentDir)

	if m.loading {
		s += "Loading tests...\n"
		return s
	}

	if m.errorMsg != "" {
		s += errorStyle.Render("Error: "+m.errorMsg) + "\n\n"
	}

	if len(m.tests) == 0 {
		s += "No tests found in this directory.\n"
	} else {
		for i, test := range m.tests {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			style := itemStyle
			if m.cursor == i {
				style = selectedStyle
			}

			icon := "📄"
			nameStyle := style
			if test.IsDir {
				icon = "📁"
				nameStyle = style.Copy().Inherit(directoryStyle)
			} else if test.HasTests {
				icon = "🧪"
				nameStyle = style.Copy().Inherit(testFileStyle)
			}

			s += fmt.Sprintf("%s %s %s\n", cursor, icon, nameStyle.Render(test.Name))
		}
	}

	s += "\n" + lipgloss.NewStyle().Faint(true).Render(
		"↑/↓: navigate • enter: select • r: refresh • backspace: parent • q: quit")

	return s
}

func (m model) renderTestResult() string {
	result := m.testResult
	s := titleStyle.Render("Test Result") + "\n\n"
	s += fmt.Sprintf("Path: %s\n\n", result.Path)

	if result.Success {
		s += successStyle.Render("✓ PASS") + "\n\n"
	} else {
		s += errorStyle.Render("✗ FAIL") + "\n\n"
	}

	if result.Output != "" {
		s += "Output:\n"
		s += lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(1).
			Render(result.Output) + "\n\n"
	}

	if result.Error != "" {
		s += errorStyle.Render("Error:") + "\n"
		s += lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("196")).
			Padding(1).
			Render(result.Error) + "\n\n"
	}

	s += lipgloss.NewStyle().Faint(true).Render("Press any key to return")
	return s
}

func discoverTests(dir string) tea.Cmd {
	return func() tea.Msg {
		tests, err := findTests(dir)
		if err != nil {
			return errorMsg(err.Error())
		}
		return testDiscoveredMsg(tests)
	}
}

func runTest(testPath string) tea.Cmd {
	return func() tea.Msg {
		result := TestResult{
			Path: testPath,
		}

		var cmd *exec.Cmd
		if strings.HasSuffix(testPath, "_test.go") {
			dir := filepath.Dir(testPath)
			cmd = exec.Command("go", "test", "-v", "./")
			cmd.Dir = dir
		} else {
			cmd = exec.Command("go", "test", "-v")
			cmd.Dir = testPath
		}

		output, err := cmd.CombinedOutput()
		result.Output = string(output)

		if err != nil {
			result.Success = false
			result.Error = err.Error()
		} else {
			result.Success = true
		}

		return testResultMsg(result)
	}
}

func findTests(rootDir string) ([]TestItem, error) {
	var items []TestItem

	if rootDir != "." {
		items = append(items, TestItem{
			Path:  filepath.Dir(rootDir),
			Name:  "..",
			IsDir: true,
		})
	}

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == rootDir {
			return nil
		}

		if strings.HasPrefix(filepath.Base(path), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.IsDir() && filepath.Base(path) == "vendor" {
			return filepath.SkipDir
		}

		rel, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}

		depth := strings.Count(rel, string(filepath.Separator))
		if depth > 0 {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		item := TestItem{
			Path:  path,
			Name:  filepath.Base(path),
			IsDir: info.IsDir(),
		}

		if info.IsDir() {
			item.HasTests = hasTestFiles(path)
		} else if strings.HasSuffix(path, "_test.go") {
			item.HasTests = true
		}

		if item.IsDir || item.HasTests {
			items = append(items, item)
		}

		return nil
	})

	return items, err
}

func hasTestFiles(dir string) bool {
	files, err := os.ReadDir(dir)
	if err != nil {
		return false
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), "_test.go") {
			return true
		}
	}
	return false
}

func main() {
	startDir := "."
	if len(os.Args) > 1 {
		startDir = os.Args[1]
	}

	if _, err := os.Stat(startDir); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist: %s\n", startDir)
		os.Exit(1)
	}

	m := initialModel()
	m.currentDir = startDir

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
