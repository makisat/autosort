# autosort

A lightweight file watcher written in Go that automatically organizes Downloads folder by moving newly added files into categorized subdirectories based on the file extension.

## Purpose

Managing a cluttered Downloads folder is annoying. autosort solves this by running in the background and sorting any new file the moment it appears and no manual organization needed.

## Technologies Used

- [`fsnotify`](https://github.com/fsnotify/fsnotify) — cross-platform filesystem event watcher
- [`godotenv`](https://github.com/joho/godotenv) — loads environment variables from a `.env` file

## Setup & Installation

### Prerequisites

- Go 1.18+

### Clone and Install Dependencies

```bash
git clone https://github.com/yourusername/autosort.git
cd autosort
go mod tidy
```

### Configuration

Create a `.env` file in the project root and set your Downloads directory path:

```env
DOWNLOAD_DIRECTORY=C:\Users\YourName\Downloads
```

### Run

```bash
go run main.go
```

Or build a binary:

```bash
go build -o autosort.exe
./autosort.exe
```

The program runs in the background and watches the configured directory until manually stopped.

## Key Features

- **Real-time sorting** — uses filesystem events to detect and move files the instant they're created
- **Auto-creates subdirectories** — if a category folder doesn't exist yet, it's created automatically
- **Configurable watch directory** — set via `.env`, no hardcoded paths
- **Six built-in categories:**

| Category  | Extensions                        |
|-----------|-----------------------------------|
| images    | `.jpeg`, `.jpg`, `.png`, `.gif`   |
| videos    | `.mp4`, `.avi`, `.mov`            |
| sounds    | `.mp3`, `.wav`, `.ogg`            |
| documents | `.pdf`, `.docx`, `.html`, `.txt`  |
| zip       | `.zip`                            |
| exe       | `.exe`                            |

## Sample Output

When a file is downloaded, autosort silently moves it:

```
Downloads/
├── images/
│   └── screenshot.png
├── documents/
│   └── notes.pdf
├── zip/
│   └── archive.zip
└── videos/
    └── clip.mp4
```

## My Contribution

I designed and built the entire project independently — the file watcher event loop, the extension-matching logic, the `.env`-based configuration, and the automatic subdirectory creation.

## Reflection

The trickiest part was handling the `fsnotify` event loop correctly inside a goroutine while keeping the main thread alive with a blocking channel receive. I also learned that filesystem events can fire before a file is fully written, which is worth handling more robustly in the future. A planned improvement is making the file categories fully configurable via a JSON or TOML config file rather than being hardcoded.
