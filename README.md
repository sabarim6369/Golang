# Golang


# Modulecreation
module go init modulename

Files under same folder cant have different packages.All should be same package


1. Package
A package is a collection of Go source files in the same folder that are compiled together.

Each Go file in that folder starts with the same package name.

Packages organize code into reusable units.

You import packages in your Go code to use their functions, types, variables, etc.

Example:

Package name: utils

Folder: utils/

You import it as: import "your_module_name/utils"

2. Module
A module is a collection of related Go packages together as a single unit.

A module is defined by a go.mod file at its root.

It represents your whole project or a collection of packages and dependencies.

Modules manage:

Your own packages

External dependencies with versions

Dependency graph and reproducible builds

You use the module name as the root path for importing your own packages.

How they relate
Module = project root + all packages inside it + dependencies + go.mod

Package = folder inside the module with Go source files sharing the same package name.

