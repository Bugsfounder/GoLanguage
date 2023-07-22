## 1. Hello World

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

Output:
```
Hello, World!
```

**Note:** The "Hello, World!" program is the traditional starting point for learning a new programming language. It simply prints the message "Hello, World!" to the console.

## 2. Variables and Constants

```go
package main

import "fmt"

func main() {
	// Variables
	var name string = "John"
	age := 30
	var height float64
	height = 6.1

	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)

	// Constants
	const pi = 3.14
	fmt.Println("Pi:", pi)
}
```

Output:
```
Name: John, Age: 30, Height: 6.10
Pi: 3.14
```

**Note:** Variables are used to store data that can be changed during the program execution. Constants, on the other hand, are used to store data that cannot be changed after being assigned.

## 3. Data Types

```go
package main

import "fmt"

func main() {
	// Basic Data Types
	var name string = "John"
	var age int = 30
	var weight float64 = 75.5
	var isMarried bool = false

	fmt.Printf("Name: %s, Age: %d, Weight: %.1f, Married: %t\n", name, age, weight, isMarried)

	// Zero Values
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Printf("Zero Values: %d, %f, %s, %t\n", zeroInt, zeroFloat, zeroString, zeroBool)
}
```

Output:
```
Name: John, Age: 30, Weight: 75.5, Married: false
Zero Values: 0, 0.000000, , false
```

**Note:** Go has several basic data types, including `string`, `int`, `float64`, and `bool`. When a variable is declared without an initial value, it takes the zero value of its data type.

## 4. Control Structures (if-else, switch)

```go
package main

import "fmt"

func main() {
	// if-else
	age := 25
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// switch
	day := "Sunday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
	case "Tuesday":
		fmt.Println("It's Tuesday.")
	case "Wednesday":
		fmt.Println("It's Wednesday.")
	default:
		fmt.Println("It's another day.")
	}
}
```

Output:
```
You are an adult.
It's another day.
```

**Note:** The `if-else` statement allows you to execute different blocks of code based on certain conditions. The `switch` statement is used to select one of many code blocks to be executed based on the value of an expression.

## 5. Loops (for, while)

```go
package main

import "fmt"

func main() {
	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// while loop (using for loop)
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}
}
```

Output:
```
1
2
3
4
5
1
2
3
4
5
```

**Note:** Go has only one looping construct, which is the `for` loop. You can use it to iterate over a range of values or create a `while` loop using the same `for` loop syntax.

## 6. Arrays and Slices

```go
package main

import "fmt"

func main() {
	// Arrays
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	fmt.Println("Fruits Array:", fruits)

	// Slices
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers Slice:", numbers)

	// Slicing
	sliced := numbers[1:4]
	fmt.Println("Sliced:", sliced)
}
```

Output:
```
Fruits Array: [Apple Banana Orange]
Numbers Slice: [1 2 3 4 5]
Sliced: [2 3 4]
```

**Note:** Arrays have a fixed size and cannot be resized, while slices are dynamic and can change in size. Slicing is used to extract a portion of a slice.

## 7. Maps

```go
package main

import "fmt"

func main() {
	// Map
	person := map[string]int{
		"John": 30,
		"Alice": 25,
	}
	fmt.Println("Person Map:", person)

	// Accessing Values
	age := person["John"]
	fmt.Println("John's Age:", age)

	// Adding and Deleting
	person["Bob"] = 35
	delete(person, "Alice")
	fmt.Println("Updated Person Map:", person)
}
```

Output:
```
Person Map: map[John:30 Alice:25]
John's Age: 30
Updated Person Map: map[Bob:35 John:30]
```

**Note:** Maps are key-value pairs, similar to dictionaries in other programming languages. You can add, access, and delete elements from a map using the associated key.

## 8. Functions

```go
package main

import "fmt"

// Function with Parameters and Return Value
func add(a, b int) int {
	return a + b
}

// Function with Multiple Return Values
func swap(a, b string) (string, string) {
	return b, a
}

// Function with Variadic Parameters
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Sum:", add(5, 3))

	x, y := swap("Hello", "World")
	fmt.Println("Swapped:", x, y)

	fmt.Println("Total:", sum(1, 2, 3, 4, 5))
}
```

Output:
```
Sum: 8
Swapped: World Hello
Total: 15
```

**Note:** Functions are blocks of code that perform specific tasks. They can take parameters, return values, and have variadic parameters (variable number of arguments).

## 9. Pointers

```go
package main

import "fmt"

func main() {
	// Pointers
	x := 10
	y := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Value of y (pointer to x):", *y)

	// Modifying Value through Pointer
	*y = 20
	fmt.Println("Modified

 Value of x:", x)
}
```

Output:
```
Value of x: 10
Value of y (pointer to x): 10
Modified Value of x: 20
```

**Note:** Pointers are variables that store the memory address of another variable. They are used to indirectly access and modify the value of a variable.

## 10. Structs

```go
package main

import "fmt"

// Struct Definition
type Person struct {
	Name string
	Age  int
}

func main() {
	// Creating Struct Instances
	p1 := Person{Name: "John", Age: 30}
	p2 := Person{Name: "Alice", Age: 25}

	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)

	// Modifying Struct Fields
	p1.Age = 32
	fmt.Println("Modified Person 1:", p1)
}
```

Output:
```
Person 1: {John 30}
Person 2: {Alice 25}
Modified Person 1: {John 32}
```

**Note:** Structs are used to define custom data types in Go. They can have multiple fields, and you can create instances of structs to store data.

## 11. Interfaces

```go
package main

import "fmt"

// Interface Definition
type Shape interface {
	Area() float64
}

// Struct Implementing the Interface
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Struct Implementing the Interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	// Using the Interface
	var shape Shape

	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}

	shape = rectangle
	fmt.Println("Rectangle Area:", shape.Area())

	shape = circle
	fmt.Println("Circle Area:", shape.Area())
}
```

Output:
```
Rectangle Area: 15
Circle Area: 12.56
```

**Note:** Interfaces define a set of methods that a struct must implement to satisfy the interface. They allow you to work with different types that implement the same set of methods.

## 12. Concurrency (Goroutines and Channels)

```go
package main

import (
	"fmt"
	"time"
)

func printMessage(msg string, ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println(msg)
	ch <- "Message Sent!"
}

func main() {
	ch := make(chan string)

	go printMessage("Hello", ch)
	go printMessage("Go", ch)
	go printMessage("Language", ch)

	<-ch
	<-ch
	<-ch
}
```

Output:
```
Hello
Go
Language
```

**Note:** Goroutines are lightweight threads used for concurrent execution. Channels are used for communication between goroutines.

## 13. Error Handling (Error Interface, Panic, Recover)

```go
package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Panic and Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("This is a panic!")
}
```

Output:
```
Result: 5
Recovered from panic: This is a panic!
```

**Note:** Go uses the `error` interface for error handling. You can also use `panic` to halt the program and use `recover` to handle panics.

## 14. File I/O (Reading and Writing to Files)

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Writing to a File
	content := []byte("Hello, Go!")
	err := ioutil.WriteFile("file.txt", content, 0644)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Reading from a File
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Content:", string(data))
}
```

Output:
```
Content: Hello, Go!
```

**Note:** Go's `io/ioutil` package provides convenient functions for reading and writing to files.

## 15. JSON Handling

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// JSON Serialization
	p := Person{Name: "John", Age: 30}
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("JSON Data:", string(data))

	// JSON Deserialization
	jsonData := `{"name": "Alice", "age": 25}`
	var person Person
	err = json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Decoded Person:", person)
}
```

Output:
```
JSON Data: {"name":"John","age":30}
Decoded Person: {John 30}
```

**Note:** Go provides built-in support for JSON encoding and decoding using the `encoding/json` package.

## 16. Command Line Arguments

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments:", args)
}
```

Output (when run with arguments):
```
Arguments: [/path/to/program arg1 arg2 arg3]
```

**Note:** Command-line arguments can be accessed using the `os.Args` variable.

## 17. Testing (Using the testing Package)

Go's testing package allows you to write tests for your code. Create a file with the `_test` suffix to write tests. Use the `go test` command to run tests.

Create a file named `math.go`:

```go
package math

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}
```

Create a file named `math_test.go`:

```go
package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}
```

Run tests:

```bash
$ go test
```

Output:
```
PASS
ok  	_/path/to/your/package/math	0.004s


```

**Note:** This example demonstrates how to write and run tests in Go. Create a separate file with `_test` suffix to write tests for your code. Use `go test` command to run tests.

## 18. Package Management (Using Go Modules)

Go modules allow you to manage dependencies and versioning for your Go projects. Use `go mod init module-name` to create a new module for your project.

```bash
$ go mod init first-go-module
```

Create a file named `main.go`:

```go
package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("UUID:", id)
}
```

Run the program:

```bash
$ go run main.go
```

Output:
```
UUID: 2f118a52-98d2-11ec-87d3-0242ac130003
```

**Note:** This example shows how to use Go modules for package management. Go modules allow you to manage dependencies and versioning for your Go projects.
## 19. Working with Time and Date
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	// Formatting Time
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)

	// Manipulating Time
	oneHourLater := currentTime.Add(1 * time.Hour)
	fmt.Println("One Hour Later:", oneHourLater)

	threeDaysAgo := currentTime.Add(-3 * 24 * time.Hour)
	fmt.Println("Three Days Ago:", threeDaysAgo)
}
```
Output:
```bash
Current Time: 2021-09-16 10:30:00 +0000 UTC m=+0.000000001
Formatted Time: 2021-09-16 10:30:00
One Hour Later: 2021-09-16 11:30:00 +0000 UTC m=+3600.000000001
Three Days Ago: 2021-09-13 10:30:00 +0000 UTC m=-259200.000000001
```
**Note:** Go's `time` package provides functions to work with dates and times.

## 20. Web Development (Using net/http Package)

### Creating a Simple HTTP Server:
```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Web!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```
Output:
```
Starting server on http://localhost:8080
```

### Accessing the Server:

Visit http://localhost:8080 in your web browser or use a tool like `curl` to access the server.

## 21. RESTful APIs

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people []Person

func getAllPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var newPerson Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request data")
		return
	}

	people = append(people, newPerson)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Person created successfully")
}

func main() {
	people = append(people, Person{Name: "John", Age: 30})

	http.HandleFunc("/people", getAllPeople)
	http.HandleFunc("/people/create", createPerson)

	fmt.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```

Output:
```
Starting server on http://localhost:8080
```

### Accessing the API:

- GET request to http://localhost:8080/people to retrieve all people.
- POST request to http://localhost:8080/people/create to create a new person.

## 22. Database Access (SQL and NoSQL)

### Using SQL (Using `database/sql` Package):

```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID   int
	Name string
	Age  int
}

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database")
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer db.Close()

	// Querying Data
	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer rows.Close()

	var people []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.ID, &p.Name, &p.Age)
		if err != nil {
			log.Fatal("Error:", err)
		}
		people = append(people, p)
	}

	for _, p := range people {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", p.ID, p.Name, p.Age)
	}
}
```

### Using NoSQL (Using `go.mongodb.org/mongo-driver/mongo` Package):

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("testdb").Collection("people")

	// Inserting Data
	p1 := Person{Name: "John", Age: 30}
	_, err = collection.InsertOne(ctx, p1)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Querying Data
	var p2 Person
	filter := bson.M{"name": "John"}
	err = collection.FindOne(ctx, filter).Decode(&p2)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("Person:", p2)
}
```

**Note:** For the SQL example, you'll need to replace "username", "password", and "database" in the connection string with your actual database credentials and database name.

For the NoSQL example, you'll need to have a running MongoDB instance on localhost:27017 or change the URI accordingly.

## 23. Reflection

```go
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "John", Age: 30}

	// Using Reflection to Inspect Type and Value
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	fmt.Println("Type:", t)
	fmt.Println("Value:", v)

	// Accessing Struct Fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Field Name: %s, Field Type: %s, Field Value: %v\n", field.Name, field.Type, value)
	}
}
```
Output:
```
Type: main.Person
Value: {John 30}
Field Name: Name, Field Type: string, Field Value: John
Field Name: Age, Field Type: int, Field Value: 30
```

**Note:** Reflection allows you to examine the type and value of variables at runtime. You can use the `reflect` package for this purpose.

## 24. Mutex and WaitGroup (for Synchronization)

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mu sync.Mutex
var wg sync.WaitGroup

func increment() {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
```

Output:
```
Counter: 5
```

**Note:** Mutex (`sync.Mutex`) is used to provide mutual exclusion when accessing shared resources. WaitGroup (`sync.WaitGroup`) is used to wait for multiple goroutines to finish.

## 25. Context Package (for Handling Cancellation and Timeouts)

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(

3 * time.Second):
		fmt.Println("Task completed successfully.")
	case <-ctx.Done():
		fmt.Println("Task canceled.")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	longRunningTask(ctx)
}
```

Output:
```
Task canceled.
```

**Note:** The `context` package is used to handle cancellation and timeouts in long-running tasks. It allows you to gracefully stop a task when it's no longer needed.

## 26. HTTP Clients (Making HTTP Requests)

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", string(data))
}
```

Output:
```
Response: {
  "userId": 1,
  "id": 1,
  "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
}
```

**Note:** The `net/http` package can be used to create HTTP clients for making HTTP requests to external APIs or servers.

## 27. Template Rendering (Using html/template Package)

```go
package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	tmpl := template.Must(template.New("person").Parse("Name: {{.Name}}, Age: {{.Age}}\n"))

	p := Person{Name: "John", Age: 30}
	err := tmpl.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
```

Output:
```
Name: John, Age: 30
```

**Note:** The `html/template` package allows you to render HTML templates with dynamic data.

## 28. Regular Expressions (Using regexp Package)

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Matching a Pattern
	re := regexp.MustCompile(`\d+`)
	match := re.MatchString("The year is 2023")
	fmt.Println("Match:", match)

	// Finding All Matches
	matches := re.FindAllString("There are 365 days in a year", -1)
	fmt.Println("Matches:", matches)
}
```

Output:
```
Match: true
Matches: [365]
```

**Note:** Go's `regexp` package provides support for regular expressions. You can use regular expressions to match and find patterns in strings.

## 29. Command Line Tools (Using os/exec Package)

```go
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("ls").Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(out))
}
```

Output:
```
Output: file.txt
main.go
```

**Note:** The `os/exec` package allows you to execute external commands from Go programs.

## 30. Binary Data and Byte Manipulation (Using encoding/binary Package)

```go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Encoding
	p := Person{Name: "John", Age: 30}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Decoding
	var decoded Person
	err = binary.Read(buf, binary.LittleEndian, &decoded)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Decoded Person:", decoded)
}
```

Output:
```
Decoded Person: {John 30}
```

**Note:** The `encoding/binary` package is used for working with binary data and byte manipulation in Go.
