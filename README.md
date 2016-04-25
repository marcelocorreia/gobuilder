# Turtle

- [About Turtle](#about-turtle)
- [Turtle file](#turtle-file)
- [Turtle and GO](#turtle-and-Go)
- [Packing Stuff with Turtle](#packing-stuff-with-turtle)
- [Deploying Stuff with Turtle](#seploying-stuff-with-turtle)
- [Project Types](#project-types)

##About Turtle
Turtle is a build and deployment helper tool where goodies will be implemented in a turtle pace.
It is trying to bring some of the concepts from Maven to other platforms. Not too ambitious just filling a gap on BAU operations.


## Project Types
Currently Turtle supports two project types.
- [x] Go Projects (via gb wrapper)
- [x] AngularJS Apps
- Generic static projects, some different projects fit under this category. Been using for
    - [x] AngularJS apps
    - [Terraform](https://terraform.io) templates

##Packing Stuff with Turtle

###GO builds

Turtle generates packages tarball packages using the definitions in turtle.json

It creates one package per definition in [Turtle File](#full-turtle-file), using a name convention similar to Maven style.
The (Build section)[#project-section]defines the build packages files to be created
Example:
For a project as below it will generate the following:
- turtle-darwin-amd64-0.0.1-SNAPSHOT.tar.gz
- turtle-linux-amd64-0.0.1-SNAPSHOT.tar.gz
- turtle-windows-amd64-0.0.1-SNAPSHOT.tar.gz

#### Run
```
$:> turtle dist
```

### Turtle file definition

> NOTE: However packaging is present in the Turtle file, ONLY .tar.gz is supported at the moment.

## Deploying Stuff with Turtle

Supported Schemes
- [x] Sonatype Nexus
- [ ] Git
- [x] SSH

Deploying to Nexus
Repositories are defined in the turtle file.
//TODO: Auth



#### Deploy to Sonatype Nexus
```
$:> turtle deploy nexus -f $DIST_FILE
```
## Turtle file

Turtle file is the project definition used by turtle to define properties of the project as well as packaging, builds,


#### Project Section
```
{
  "group-id": "io.correia",
  "artifact-id": "turtle",
  "name": "turtle",
  "version": "0.0.1-SNAPSHOT",
  "packaging": "tar.gz",
  "generate-pom": false,
  "project-type": "go",
  ...
}
```

#### Builds Section
```
{
  ...
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
  ...
}

```
#### Repositories Section
```
 "repositories": [
      {
        "id": "my-nexus-repo",
        "type": "nexus",
        "build-type": "snapshots",
        "url": "http://my-nexus:8081/nexus/content/repositories/snapshots"
      },
      {
        "id": "my-nexus-repo",
        "type": "nexus",
        "build-type": "releases",
        "url": "http://my-nexus:8081/nexus/content/repositories/releases"
      }
    ],
```

#### Full Turtle File
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
        "url": "http://my-nexus:8081/nexus/content/repositories/snapshots"
      },
      {
        "id": "my-nexus-repo",
        "type": "nexus",
        "build-type": "releases",
        "url": "http://my-nexus:8081/nexus/content/repositories/releases"
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
