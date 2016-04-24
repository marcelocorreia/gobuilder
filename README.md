# Turtle - Build Helper

Turtle is a build helper tool where goodies will be implemented. It is trying to bring some of the concepts from Maven other platforms.

- [Turtle and GO](#Turtle and Go)

## Turtle and Go

Turtle offers a wrapper for gb https://getgb.io

gb is an awesome tool for Golang projects and provides features like:
- Project based workflow
- Automatic project detection
- Dependency management
- Test run

gb is an alternative build tool for the [Go programming language](https://golang.org/).

[Read more about the rationale for gb](https://getgb.io/rationale).

Project-Based
gb operates on the concept of a project. A gb project is a workspace for all the Go code that is required to build your project.

A gb project is a folder on disk that contains a sub directory named src/. That’s it, no environment variables to set. For the rest of this document we’ll refer to your gb project as $PROJECT.

You can create as many projects as you like and move between them simply by changing directories.

[Read more about setting up a gb project](https://getgb.io/docs/project).
