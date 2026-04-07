# Totion

Inspired from CodersGyan YT video

## Overview

Totion is a command-line note-taking tool that helps you quickly capture, organize, and edit your thoughts right from your terminal. It's built for developers who want a seamless, keyboard-driven way to manage notes without ever leaving their command-line environment, keeping all your important information neatly stored in a dedicated directory on your system.

## Features

- **Create New Notes**: Easily start a new note file with a simple keyboard shortcut and name it on the fly.
- **List All Notes**: View an interactive, filterable list of all your saved notes, making it easy to find what you need.
- **Edit Existing Notes**: Open any note from the list and modify its content directly within the terminal interface.
- **Save Notes**: Quickly save your changes to the current note with a dedicated command.
- **Persistent Storage**: All your notes are stored in a hidden directory (`~/.totion`) in your home folder for easy access and organization.

## Getting Started

To get Totion up and running on your local machine, follow these steps.

### Installation

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/SaranHiruthikM/totion.git
    ```
2.  **Navigate into the project directory**:
    ```bash
    cd totion
    ```
3.  **Build the executable**:
    This command compiles the Go source code and creates a `totion` executable in your current directory.
    ```bash
    go build -o totion .
    ```
    _Alternatively, you can use the provided `makefile` to build:_
    ```bash
    make build
    ```
4.  **(Optional) Install to your PATH**:
    To run `totion` from any directory, move the executable to a directory included in your system's PATH, like `/usr/local/bin`.
    ```bash
    sudo mv totion /usr/local/bin/
    ```

### Environment Variables

This project does not require any specific environment variables to run.

## Usage

Once installed, you can start Totion by running the executable from your terminal.

If you installed it to your PATH:

```bash
totion
```

Otherwise, if you built it in the project directory:

```bash
./totion
```

You'll be greeted with the Totion interface. Here are the key commands to navigate and manage your notes:

- `Ctrl+N`: Start creating a **new file**. You'll be prompted to enter a file name. Type the name and press `Enter` to confirm and start writing.
- `Ctrl+L`: **List all your notes**. This will bring up an interactive list where you can use arrow keys to navigate and type to filter. Press `Enter` to open a selected note.
- `Esc`: Use this to **go back** from a prompt (like new file name input), close the notes list, or clear the currently open note.
- `Ctrl+S`: **Save** the changes to the currently open note.
- `Ctrl+C` or `q`: **Quit** the Totion application.

## Technologies Used

| Technology     | Description                                                                                  | Link                                                                             |
| :------------- | :------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------- |
| **Go**         | The primary programming language used for building the application.                          | [go.dev](https://go.dev/)                                                        |
| **Bubble Tea** | A powerful framework for creating beautiful terminal user interfaces.                        | [github.com/charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) |
| **Bubbles**    | A collection of reusable components for Bubble Tea (e.g., list, textarea, textinput models). | [github.com/charmbracelet/bubbles](https://github.com/charmbracelet/bubbles)     |
| **Lipgloss**   | A style definition and rendering package for elegant terminal layouts and rich text.         | [github.com/charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)   |

## Contributing

We welcome contributions to Totion! If you have an idea for a new feature, find a bug, or want to improve the codebase, please consider contributing. Here's a general guideline:

1.  Fork the repository to your GitHub account.
2.  Create a new branch for your specific feature or bugfix.
3.  Implement your changes, ensuring they align with the project's coding style and pass any existing tests.
4.  Write clear and concise commit messages.
5.  Push your branch to your forked repository.
6.  Open a pull request to the main repository, describing your changes in detail.

Your contributions are greatly appreciated!

## License

This project is open-source. For detailed licensing information, please refer to the project's repository.

## Badges

[![Go](https://img.shields.io/badge/Go-1.25.8-00ADD8?logo=go)](https://go.dev/)
[![Bubble Tea](https://img.shields.io/badge/Bubble%20Tea-1.3.10-purple?logo=go)](https://github.com/charmbracelet/bubbletea)
[![Lipgloss](https://img.shields.io/badge/Lipgloss-1.1.0-blueviolet?logo=go)](https://github.com/charmbracelet/lipgloss)
