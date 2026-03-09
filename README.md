# golang-application

A simple Golang web application using the Gin framework that displays a random movie name on each browser refresh.

## Features
- Random movie selection from a predefined list
- Clean, responsive UI with gradient background
- Auto-refresh capability to get new movies
- Built with Gin web framework

## Prerequisites
- Go 1.21 or higher
- Git (optional)

## Project Structure
```
golang-application/
├── main.go              # Main application file
├── go.mod               # Go module dependencies
├── templates/
│   └── index.html       # HTML template
└── README.md            # Documentation
```

## Installation

1. Clone or navigate to the project directory:
```bash
cd golang-application
```

2. Install dependencies:
```bash
go mod tidy
```

## Running the Application

```bash
go run main.go
```

The application will start on `http://localhost:8080`

## Usage

1. Open your browser
2. Navigate to `http://localhost:8080`
3. Refresh the page to see a different random movie

## How It Works

### main.go
- Imports Gin framework and required packages
- Defines a slice of 10 popular movies
- Seeds random number generator for true randomness
- Creates a GET route (`/`) that selects a random movie
- Renders the HTML template with the selected movie
- Starts server on port 8080

### templates/index.html
- Displays the random movie name using Gin's template syntax
- Styled with CSS for a modern, centered layout
- Uses gradient background and card-style container

## Technologies Used
- **Go** - Programming language
- **Gin** - Web framework
- **HTML/CSS** - Frontend presentation

## Customization

To add more movies, edit the `movies` slice in `main.go`:
```go
var movies = []string{
    "Your Movie 1",
    "Your Movie 2",
    // Add more...
}
```

To change the port, modify the `Run()` parameter in `main.go`:
```go
r.Run(":3000")  // Change to port 3000
```

## License
MIT
