
# Script Starter
![License](https://img.shields.io/github/license/sefatanam/script-starter)

Script Starter is a Go-based project template designed to help developers quickly set up and start working on new applications. This template includes essential configurations and basic structures to streamline the development process.

## Features

- **Basic Project Structure**: Includes a clean and organized project layout.
- **Configuration Management**: Utilizes `FyneApp.toml` for easy configuration.
- **Dependency Management**: Set up with `go.mod` and `go.sum` for dependency tracking.
- **Sample UI**: Provides a basic user interface setup with `Fyne`.


## Getting Started

### Prerequisites

- Go 1.16 or higher
- [FyneUI](https://fyne.io/)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/sefatanam/script-starter.git
    cd script-starter
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run main.go
    ```
## Live Demo

Check out this [30sec video from YouTube](https://youtu.be/Ny0XSw_kvCs)

## Project Structure

```
script-starter/
├── constant/
├── lib/
├── ui/
├── FyneApp.toml
├── go.mod
├── go.sum
├── icon.png
└── main.go
```

- **constant**: Contains constant values used throughout the project.
- **lib**: Includes library functions and utilities.
- **ui**: Holds user interface components.
- **FyneApp.toml**: Configuration file for the Fyne application.
- **main.go**: Entry point of the application.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
