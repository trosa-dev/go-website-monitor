# Website Monitor

Website Monitor is a simple Go application that allows users to monitor the status of websites.

## Getting Started

These instructions will help you set up and run the Website Monitor application on your local machine.

### Prerequisites

- Go (at least version 1.21.4)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/website-monitor.git
   cd website-monitor
   ```

2. Build the application:

   ```bash
   go build
   ```

### Usage

1. Run the application:

   ```bash
   ./website-monitor
   ```

2. Follow the on-screen menu to start monitoring, display logs, or exit.

## File Structure

- **`main.go`**: Entry point of the application.
- **`menu/menu.go`**: Provides a menu for user interaction.
- **`readInput/readInput.go`**: Reads user input.
- **`readFile/readFile.go`**: Reads the list of sites from a file.
- **`saveLog/saveLog.go`**: Saves logs to a file.
- **`showLogs/showLogs.go`**: Displays logs.
- **`startMonitoring/startMonitoring.go`**: Initiates the website monitoring process.
- **`testSite/testSite.go`**: Tests the status of a website.

## Configuration

- **`logs.txt`**: File to store monitoring logs.
- **`sites.txt`**: File containing the list of websites to monitor.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
