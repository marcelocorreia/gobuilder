# Turtle

- [About Turtle](#About Turtle)
- [Turtle file](#Turtle file)
- [Turtle and GO](#Turtle and Go)
- [Packing Stuff with Turtle](#Packing Stuff with Turtle)
- [Deploying Stuff with Turtle](#Deploying Stuff with Turtle)
- [Project Types](#Project Types)

##About Turtle
Turtle is a build and deployment helper tool where goodies will be implemented in a turtle pace.
It is trying to bring some of the concepts from Maven other platforms. Not too ambitious just filling a gap on BAU operations.


## Project Types
Currently Turtle supports two project types.
- Go Projects (via gb wrapper)
- Generic static projects, some different projecys might fit under this category

##Packing Stuff with Turtle

###GO builds



## Deploying Stuff with Turtle


## Turtle file

Turtle file is the project definition used by turtle to define properties of the project as well as packaging, builds,

```
{
  "group-id": "io.correia",
  "artifact-id": "turtle",
  "name": "turtle",
  "version": "0.0.1-SNAPSHOT",
  "packaging": "tar.gz",
  "generate-pom": false,
  "project-type": "go",
  "repositories": [
      {
        "id": "my-nexus-repo",
        "type": "nexus",
        "build-type": "snapshots",
        "url": "http://pc-mgmt01.products.bulletproof.net:8081/nexus/content/repositories/snapshots"
      },
      {
        "id": "my-nexus-repo",
        "type": "nexus",
        "build-type": "releases",
        "url": "http://pc-mgmt01.products.bulletproof.net:8081/nexus/content/repositories/releases"
      }
    ],
  "builds": [
    {
      "OS": "darwin",
      "Arch": "amd64"
    },
    {
      "OS": "linux",
      "Arch": "amd64"
    },
    {
      "OS": "windows",
      "Arch": "amd64"
    }
  ]
}
```
## Turtle and Go

Turtle started and a Golang helper tool and started the shifting into a more generic and smarter build helper.
It offers a wrapper for gb https://getgb.io. Other Go build tool might be integrated but for the moment only gb is available.

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
