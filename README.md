# IronPass – Secure CLI Password Generator in Go

## 1. Title & Objective

**Title:** IronPass – Secure CLI Password Generator in Go

**Chosen technology:** Go (Golang)

**Why I chose it**

I chose Go because it is simple to learn, fast to compile, and very well suited for command-line tools. For this capstone, I wanted a technology that would let me build a small but practical security-focused application without depending on many external packages. Go was a strong fit because it has a clean syntax, excellent standard library support, and built-in tools for creating reliable CLI programs.

**Project introduction**

IronPass is a lightweight command-line password generator built with Go. It helps users create strong, secure, and customizable passwords directly from the terminal. Instead of reusing weak passwords or depending on web-based generators, a user can run one command and instantly get a secure password tailored to their needs.

This project supports password length customization and lets the user decide whether to include lowercase letters, uppercase letters, numbers, and symbols. The passwords are generated using Go's `crypto/rand` package so the randomness is suitable for security-related use cases.

**End goal**

The goal of this project is to build a simple, secure, and customizable CLI tool that generates strong passwords directly from the terminal, while also documenting the learning process in a beginner-friendly way.

## 2. Quick Summary of the Technology

**What is Go?**

Go, also called Golang, is an open-source programming language created by Google. It is designed for simplicity, performance, readability, and reliability.

**Where is it used?**

- Command-line tools
- Backend APIs and web services
- Cloud-native applications
- DevOps and automation tools
- Security and infrastructure utilities

**One real-world example**

Docker, the widely used containerization platform, is written in Go.

## 3. System Requirements

- OS: Linux, macOS, or Windows
- Editor: VS Code or any text editor
- Go version: Go 1.20 or newer recommended
- Dependencies: None outside the Go standard library

## 4. Installation & Setup Instructions

### Clone or open the project

If using Git:

```bash
git clone https://github.com/ANNGLORIOUS/Generative-AI
cd FOLDER-2
```

If the project is already on your machine, just open the folder in your terminal:

```bash
cd /home/annglorious/Documents/code/GENERATIVE-AI/FOLDER-2
```

### Make sure Go is available

Check whether Go is installed:

```bash
go version
```

If `go` is not installed system-wide, this project also includes a local Go setup for this folder. In that case, run:

```bash
source ./env.sh
go version
```

### Build the project

```bash
go build -o ironpass
```

### Run the project

Generate one password with default settings:

```bash
./ironpass
```

Generate a 20-character password:

```bash
./ironpass -length 20
```

Generate 3 passwords without symbols:

```bash
./ironpass -count 3 -symbols=false
```

Generate a password using only lowercase letters and numbers:

```bash
./ironpass -upper=false -symbols=false
```

## 5. Minimal Working Example

### What the example does

This example generates one secure password of length 16 using lowercase letters, uppercase letters, numbers, and symbols.

### Example command

```bash
./ironpass -length 16 -count 1
```

### Example code

The application uses secure random generation from Go's standard library:

```go
index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
if err != nil {
    return 0, fmt.Errorf("failed to generate secure random index: %w", err)
}
```

This is important because `crypto/rand` is more appropriate for password generation than non-secure random generators.

### Expected output

```text
9@Qa!Lm2#Zt7^PxR
```

The exact password will be different every time because it is randomly generated.

## 6. AI Prompt Journal

This section documents how generative AI supported my learning and project development.

### Prompt 1

**Prompt used:**  
"Give me a beginner-friendly explanation of Go and why it is a good language for building CLI tools."

**Curriculum link for the prompt:**  
ai.moringaschool.com

**Response summary:**  
The AI explained Go's strengths, including simplicity, fast compilation, and a strong standard library.

**Helpful part of the response:**  
It helped me understand why Go was a better fit for this project than choosing a heavier framework or language with more setup overhead.

**My evaluation:**  
Helpful. It gave me the confidence to choose Go for a first beginner-focused toolkit project.

### Prompt 2

**Prompt used:**  
"Show me how to build a secure password generator in Go using only the standard library."

**Curriculum link for the prompt:**  
ai.moringaschool.com

**Response summary:**  
The AI suggested using `crypto/rand`, building character sets, and making the tool configurable with flags.

**Helpful part of the response:**  
It pointed me toward secure randomness instead of weaker random number generation methods.

**My evaluation:**  
Very helpful. This directly influenced the core implementation of IronPass.

### Prompt 3

**Prompt used:**  
"Help me write a beginner-friendly README for a Go CLI password generator based on a capstone assignment format."

**Curriculum link for the prompt:**  
ai.moringaschool.com

**Response summary:**  
The AI helped organize the README into installation steps, example usage, common errors, and learning reflections.

**Helpful part of the response:**  
It improved the clarity and structure of the documentation so another beginner could follow it more easily.

**My evaluation:**  
Helpful. It saved time and improved the final presentation of the project.

### Reflection on AI usage

Generative AI made it easier for me to learn faster, especially when breaking down unfamiliar concepts and turning them into actionable steps. It was most useful for structuring my learning, debugging environment issues, and improving documentation. I still needed to test everything locally and confirm that the commands actually worked.

## 7. Common Issues & Fixes

### Issue 1: `go: command not found`

**Cause:**  
Go was not installed or not available in the current shell.

**Fix:**  
Install Go or load the local environment file:

```bash
source ./env.sh
```

Then confirm:

```bash
go version
```

### Issue 2: Build cache permission problems

**Cause:**  
Some environments restrict writing to the default Go cache directory.

**Fix:**  
This project uses a local cache configuration in `env.sh`, which keeps Go's cache inside the project folder.

### Issue 3: No output file after build

**Cause:**  
The build command may not have completed successfully.

**Fix:**  
Run:

```bash
go build -o ironpass
ls
```

Make sure the `ironpass` binary appears in the project folder.

### Issue 4: Invalid password settings

**Cause:**  
A password cannot be generated if all character groups are disabled or if the requested length is too short.

**Fix:**  
Use at least one character set and keep the length high enough. For example:

```bash
./ironpass -length 12 -symbols=false -numbers=true -upper=true -lower=true
```

## 8. References

- Go official documentation: https://go.dev/doc/
- Go installation guide: https://go.dev/doc/install
- Go `flag` package: https://pkg.go.dev/flag
- Go `crypto/rand` package: https://pkg.go.dev/crypto/rand
- Docker official website: https://www.docker.com/

## Project Files

- `main.go` contains the CLI application logic.
- `go.mod` defines the Go module.
- `env.sh` sets up the local Go environment for this project if Go is not installed globally.

## How to Test Quickly

```bash
go build -o ironpass
./ironpass
./ironpass -length 20 -count 2
./ironpass -symbols=false -numbers=false
```

## Conclusion

IronPass is a small but practical beginner toolkit project that shows how Go can be used to create secure command-line applications. It demonstrates setup, secure random password generation, argument parsing, troubleshooting, and documentation in a way that another beginner can follow and reproduce.
