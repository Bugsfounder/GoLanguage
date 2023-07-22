# Go Programming Language

## Download and Install

To get started with Go programming, you first need to download and install Go on your system. Visit the official Go website (https://golang.org/) and follow the installation instructions for your operating system.

## Setup Project and Get Started with Coding in Go

### Create a Folder and Navigate into it

Create a new folder for your first Go project and navigate into it using the following terminal command:

```bash
mkdir first-go-project && cd first-go-project
```

### Open the Folder in Your Favorite IDE

Open the newly created folder in your favorite Integrated Development Environment (IDE). If you're using Visual Studio Code, you can enter the following command in the terminal inside the `first-go-project` folder:

```bash
code .
```

### Initialize the Project Module

In Go, everything is organized into modules and packages. To get started, create a module for your Go project using the following terminal command:

```bash
go mod init first-go-project
```

### Get Started with Coding

Create a new file named `main.go` in your project directory and write the following code:

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello World\n")
}
```

Save the file. This code defines a simple Go program with a `main` function that prints "Hello World" to the console.

### Running the Code

To run the `main.go` file, enter the following command in the terminal:

```bash
go run main.go
```

You should see the following output:

```terminal
Hello World
```

## Variables

In Go, if you define a variable and do not use it anywhere in your codebase, it will result in an error. To demonstrate, consider the following example:

```go
var variableName = "Bugs"
```

When you run this file, you will encounter the following error:

```bash
# command-line-arguments
./main.go:8:6: variableName declared and not used
```

To fix the error, you need to use the variable in your codebase:

```go
var variableName = "Bugs"
fmt.Println("Welcome", variableName)
```

Now, running the file will produce the following output:

```terminal
Welcome Bugs
```

### Constants

If you want to create a constant variable, use the `const` keyword to define it:

```go
const constVariableName = 50
```

> Note: Constant variables cannot be updated; their values remain the same throughout the program execution.

Feel free to experiment with variables and constants to familiarize yourself with Go's syntax and features.

## Conclusion

Congratulations! You've taken your first steps into the world of Go programming. You've learned how to set up a project, create a simple Go program, work with variables, and define constants. Now you're ready to explore more advanced concepts and build exciting projects with Go. Happy coding!

# Go Multi-Package Project Example

This is a simple Go project that demonstrates how to work with multiple packages in Go.

## Project Overview

The project consists of two main packages, `main` and `helper`, each residing in their respective directories.

1. `main` package:
   - File: `main.go`
   - Description: The main package serves as the entry point of the application. It imports the `helper` package and calls the `Greet()` function from the `helper` package to print a greeting to the console.

2. `helper` package:
   - File: `helper/helper.go`
   - Description: The `helper` package contains utility functions used by the main package. It defines the `Greet()` function that generates a greeting message.


## Package Management in Go

In Go, proper package management is crucial for organizing code and ensuring smooth compilation and dependency resolution. Here are some key points to keep in mind:

1. Each directory should contain only one package, and the name of the directory should match the package name. For instance, the `helper` package resides in the `helper` directory.

2. To use functions and types from another package, you need to import the package using its relative or absolute import path.

3. When you have multiple packages in the same directory, it leads to conflicts during compilation. Thus, it's essential to separate packages into their respective directories.

## How to Run the Application

To run the application, follow these steps:

1. Clone this repository to your local machine.

2. Open a terminal or command prompt and navigate to the root directory of the project.

3. Execute the following command to run the application:
   ```go
   go run main.go
   ```
This will execute the `main` package, and you should see the greeting message printed to the console.
## Code Examples

### main.go

```go
package main

import (
	"fmt"
	"practice-mod/helper" // Importing the helper package
)

func main() {
	message := helper.Greet() // Using the Greet() function from the helper package
	fmt.Println(message)
}
```

### helper/helper.go

```go
package helper

func Greet() string {
	return "Hello from the helper package!"
}
```

## Additional Notes

- Adhering to Go's best practices and idiomatic style enhances code readability and maintainability.
- The Go standard library provides many useful packages to simplify various tasks, so make sure to explore it when building your projects.
- Actively participating in the Go community helps you learn from others and share your knowledge.

## Revision

If you encounter issues or need to revisit the project structure, follow these steps:
1. Ensure that each package resides in its own directory.
2. The main package should import the `helper` package using the correct path.
3. Each file containing a package should start with the `package <package-name>` statement.
4. To run the application, use the `go run` command in the root directory of the project.

Feel free to refer to this README whenever you need a quick reminder or share the project with others.

Happy coding!

With the added code examples, you can now see how the `main.go` file imports the custom `helper` package and uses its `Greet()` function. The code examples make it easier to understand how to import custom packages and use functions from those packages in your Go project. This way, you can quickly revise the concepts and avoid getting stuck in the future.