# Main package

**Main** is used as entry point to program

**Every** every main package must have main func

main package can be separated into multiple files that completlly normal

## Multiple packages

one package == one directory, You shouldn't have `.go` files with different package names inside one directory

## Importing/Exporting

To export something it needs to be **Capitalized**, if it's start with lower case letter it should be considered private for that package

Importing something if it's in the same module, start the path from `go.mod` file then path to the desired package

`import {module path}/{path to the package relative to your go.mod file}`

## Making complex file structure

This is a bit over complex for this it's just to test exec app

now app is executable with `./cmd/cli`
relative paths ensures that path is interpreted correctly

**any packages under internal cannot be imported by code outside of the "package" directory.**
