# TrackGit

TrackGit is a Git commit visualizer that displays your contributions in a calendar-style heatmap right in your terminal—similar to GitHub’s contributions graph. It scans your local Git repositories, counts commits per day, and shows them as color-coded cells. Clicking on a cell provides additional commit details.

## Go Version
1.23.1

## Features
- **Calendar Heatmap**: Visualize commit activity by day over a week, month, or year.
- **Interactive UI**: Click on a day to see details (e.g. commit count and date) using an intuitive terminal interface built with tview.
- **Customizable Views**: Easily switch between different time frames (week, month, year).
- **Modular Design**: Separate components handle commit counting, database interactions, and UI rendering.
- **Automatic Dependency Management**: Uses Go modules for dependency tracking.

## Prerequisites
- Go 1.23.1
- Git
- A terminal that supports mouse events (for interactivity)
- SQLite (if using the built-in database for commit tracking)

## Installation

### Clone the Repository:
```bash
git clone https://github.com/yourusername/TrackGit.git
cd TrackGit
```

### Install Dependencies:
Use Go modules to install dependencies. Run:
```bash
go mod tidy
```

### Make the Initialization Script Executable:
The `init.sh` script sets up your environment, builds, and runs the main application. Make it executable and run it once:
```bash
chmod +x init.sh
./init.sh
```

The `init.sh` script will:
- Export necessary environment variables (e.g. `TRACKGIT_PATH`).
- Add an alias (if needed) to your shell configuration.
- Build and run your main application.

### Troubleshooting

If you encounter errors such as `TRACKGIT_PATH not found` or `command not found: trackgit`, you may need to source your shell configuration file to ensure the environment variables and aliases are loaded correctly. Run the following command in your terminal:

```bash
source ~/.bashrc
```

If you are using a different shell (e.g., zsh), source the appropriate configuration file:

```bash
source ~/.zshrc
```

For Windows users, you can set environment variables permanently using the `setx` command in Command Prompt or PowerShell:

```powershell
setx TRACKGIT_PATH "C:\path\to\TrackGit"
```

To add an alias in PowerShell, you can modify your PowerShell profile script. Open the profile script with:

```powershell
notepad $PROFILE
```

Add the following line to create an alias:

```powershell
Set-Alias trackgit "C:\path\to\TrackGit\trackgit.exe"
```

Replace `C:\path\to\TrackGit` with the actual path to your TrackGit directory. Save and close the file, then restart your terminal.

By adding these lines to your shell configuration file or PowerShell profile, the environment variables and aliases will be set up automatically in all future terminal sessions, eliminating the need to source the file manually each time.


#### Permanent Environment Variable Setup

To avoid sourcing your shell configuration file every time, you can add the necessary environment variables and aliases permanently. Follow these steps:

1. **Open your shell configuration file**:
    - For bash users, open `~/.bashrc`:
      ```bash
      nano ~/.bashrc
      ```
    - For zsh users, open `~/.zshrc`:
      ```bash
      nano ~/.zshrc
      ```

2. **Add the following lines to the file**:
    ```bash
    export TRACKGIT_PATH=/path/to/TrackGit
    alias trackgit='/path/to/TrackGit/trackgit'
    ```

    Replace `/path/to/TrackGit` with the actual path to your TrackGit directory.

3. **Save and close the file**.

4. **Source the configuration file to apply changes**:
    - For bash users:
      ```bash
      source ~/.bashrc
      ```
    - For zsh users:
      ```bash
      source ~/.zshrc
      ```

By adding these lines to your shell configuration file, the environment variables and aliases will be set up automatically in all future terminal sessions, eliminating the need to source the file manually each time.
## Usage
After installation, you can run TrackGit via its various options. For example, you can display your commit visualizations with different time frames:

### Week View:
```bash
./trackgit streak --week
```

### Month View:
```bash
./trackgit streak --month
```

### Year View:
```bash
./trackgit streak --year
```

The interactive UI will launch in your terminal. You can navigate using the arrow keys or the mouse and click on a day to see commit details.

## Environment Setup
The `init.sh` script handles your environment setup by:
- Exporting the current directory as `TRACKGIT_PATH`.
- Configuring your shell (e.g. ensuring that `~/.bashrc` is sourced from `~/.bash_profile`).
- Adding a convenient alias for launching the commit visualizer.
- Building the project and running the main application.

**Note**: Run `init.sh` only once during your initial setup.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Additional Information
- **Go Modules**: This project uses Go modules to manage dependencies. Use `go mod tidy` to download missing packages and remove unused ones.
