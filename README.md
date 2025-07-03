# Go Test Runner

A terminal UI application for finding and running Go tests with an intuitive interface.

## Features

- 🔍 **Test Discovery**: Automatically finds all Go test files in directories
- 📁 **Directory Navigation**: Browse through your project structure
- 🧪 **Interactive Test Running**: Select and run individual test files or all tests in a directory
- 🎨 **Color-coded Results**: Green for passing tests, red for failing tests
- ⚡ **Fast and Lightweight**: Built with Bubble Tea for smooth terminal UI

## Installation

Clone this repository and build the application:

```bash
git clone <your-repo>
cd gotest-runner
go build .
```

## Usage

Run the application in the current directory:
```bash
./gotest-runner
```

Or specify a directory:
```bash
./gotest-runner /path/to/your/go/project
```

## Controls

- **↑/↓ or k/j**: Navigate through the list
- **Enter**: Select a directory (to navigate) or test file (to run)
- **r**: Refresh the current directory
- **Backspace**: Go to parent directory
- **q or Ctrl+C**: Quit the application
- **Any key**: Return from test results view

## Example

The project includes example test files in the `examples/` directory:

- `examples/calculator/` - Simple calculator functions with passing tests
- `examples/stringutils/` - String utility functions with some failing tests

Try running the application and navigating to these directories to see the tool in action!

## How it Works

1. **Discovery**: The app scans directories for `*_test.go` files
2. **Navigation**: Use arrow keys to browse directories and test files
3. **Execution**: When you select a test file, it runs `go test -v` 
4. **Results**: View detailed output with color-coded pass/fail indicators

## Requirements

- Go 1.19 or later
- Terminal with color support

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style and layout library

## License

MIT License

