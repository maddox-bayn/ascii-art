# Ascii-Art

## Overview
A robust, modular Command Line Interface (CLI) utility built in Go that processes standard string inputs and transforms them into large graphic representations using ASCII font banners. This tool includes advanced styling flags for dynamic colorization, output redirection, text alignment, and an inverse matrix engine capable of reversing ASCII art files back into plain text.

## 🛠️ Features

**Font Banner Selection**: Supports standard, shadow, and thinkertoy graphical font configurations.
**Professional Text Alignment**: System ioctl integration to pool hardware constraints dynamically, supporting --align=[left|center|right|justify].
**Targeted Substring Colorization**: State-machine color engine matching full strings or specialized subsections with specific ANSI configurations via --color=[color] [substring].
**POSIX File Redirection**: Decoupled layout printer allowing stream interception and output preservation using --output=[filename.txt].
**Matrix Reversal Parser**: Highly optimized inverse matrix matching algorithm capable of scanning historical ASCII generation files back to basic alphanumeric runes with --reverse=[filename.txt].

## 📂 Project Architecture

The codebase is engineered strictly around decoupled structural patterns:Plaintextascii-art/
├── functions/
│   ├── args_handling.go     # Position-independent flag extraction engines
│   ├── ascii_handling.go    # Structural painters & alignment layout managers
│   ├── color_handling.go    # Boolean mask generators & ANSI configurations
│   ├── content_handling.go  # File system I/O streams and storage routers
│   ├── error_handling.go    # Standard execution guard rails & validations
│   ├── file_handling.go     # Dynamic vertical sliding window reverse lookups
│   ├── justify.go           # Text spacing controllers
│   └── loadbanner.go        # Banner index database cache parsers
└── main.go                  # Central system orchestrator


🚀 Installation & SetupEnsure you have Go installed on your local operating environment.Bash# Clone the repository
git clone https://github.com/maddox-bayn/ascii-art.git
cd ascii-art

**Initialize modules**
go mod tidy

## 💻 Usage & Flag Specifications
The system utilizes custom flag validation engines following a standard POSIX convention structure:
go run . [OPTION] [SUBSTRING] [STRING] [BANNER]

### Basic Art Generation

go run . "Hello World" standard
**Redirection to File Descriptor**
go run . --output=banner.txt "Data Streams" shadow

**Colorizing Explicit Substrings**
***// Colors only the word "logic" inside the string payload***
go run . --color=red "logic" "Building logic structures" thinkertoy

**Dynamic Layout Alignment**
go run . --align=center "Centered Graphics" standard

**Running the Reverse Matrix Engine**
go run . --reverse=banner.txt

### 🎨 Supported Colors
The ANSI color configuration engine currently registers the following internal mappings:
-black, red, green, yellow, blue, magenta, cyan, orange, white
### 🧪 Algorithmic Complexity Specifications
-Matrix Generation: Runs at a clean linear computational execution scale of $O(N)$ against input lengths.
-Matrix Reversal: Leverages structural index cache lookups. Because font blocks utilize constant spatial configurations (8 rows $\times$ fixed horizontal bounds), evaluation matrices process input lines in predictable $O(N)$ runtime tracks without exponential hardware strain.
