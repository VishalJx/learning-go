# Understanding Packages in Go

## Exported Function name must be in uppercase

To make a function accessible outside its package, the function name must start with an `uppercase letter`.

**Example:**
```go
func AddNumbers(num1 int, num2 int) int {
    return num1 + num2
}
```
In this example, `AddNumbers` is exported because it starts with an uppercase letter. This means it can be used in other packages that import this package.

## Same Package Names

If we have separate files but same package names then there is no need to import it explicitly. 

**Project Structure Example:**
```
Root Folder:
├── main.go
├── add.go
```

**Key Points:**
- The package name must be `main` for all files in the same folder
- No explicit import is needed when files share the same package name

## Different Package Names

When need a different package file, it must be imported wherever it's needed, using its path.

**Import Syntax:**
```go
import x "go-a-to-z/2.Packages/extra_folder"
```

**Note:**
- You can use an `alias` for the package to simplify its usage
- In the example above, `x` is an alias for the `extra_folder` package
- This allows you to reference package functions using the alias

## Best Practices

1. Use uppercase first letters for functions you want to export
2. Keep package names short and meaningful
3. Use aliases when package names are long or potentially confusing
4. Organize your packages logically to improve code readability and maintainability

## Example Project Structure

```
2.Packages/
│
├── main.go
├── add.go
├── extra_folder/
│   └── insidefolder_add.go
├── go.mod

```
