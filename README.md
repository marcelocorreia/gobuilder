# Turtle

- [TL;DR](#tldr)
- [About Turtle](#about-turtle)
- [Turtle file](#turtle-file)
- [Turtle and GO](#turtle-and-Go)
- [Packing Stuff with Turtle](#packing-stuff-with-turtle)
- [Deploying Stuff with Turtle](#seploying-stuff-with-turtle)
- [Project Types](#project-types)



##About Turtle
Turtle is a build and deployment helper tool where goodies will be implemented in a turtle pace.
It is trying to bring some of the concepts from Maven to other platforms. Not too ambitious just filling a gap on BAU operations.

### TL;DR
#### Turtle in a nutshell
- Helper tool writen in go to make devs and devops life easier.
- Constant development atm. (Things my change quickly).
- Bunch of helpers
    - Golang
        - [x] Dependency management - via [gb](https://getgb.io)
            - [ ] Other options i.e. ([Glide](https://glide.sh) or some other). gb has done the job well so far. but with Go 1.6+ out, it might be a good idea to explore what's around
        - [x] Build helpers
            - [x] gb build wrapper
            - [x] gb test helper
        - Packaging
            - [x] Distribution packaging
              - [x] Type tar.gz based on [Turtle File](#full-turtle-file)
    - Installers
        - [x] install [gb](https://getgb.io)
        - [ ] [Hashicorp](http://www.hashicorp.com) stuff
        - [ ] [Ansible](http://www.ansible.org) stuff

    - CI / SCM Tools & helpers
        - [ ] SCM Release

    - Integration with cool toys
        - [ ] Hashicorp Ecosystem
            - [ ] [Consul](https://consul.io)
            - [ ] [Packer](https://packer.io)
            - [ ] [Vault](https://vaultproject.io)
            - [ ] [Nomad](https://nomadproject.io)
            - [ ] [Otto](https://ottoproject.io)
            - [ ] [Vagrant](https://vagrantup.com)
        - [ ] [Ansible](http://www.ansible.org)
        - [ ] [Docker](https://www.docker.com)
            - [ ] [Docker Compose](https://www.docker.com/products/docker-compose)
            - [ ] [Docker Swarm](https://www.docker.com/products/docker-swarm)
            - [ ] [Docker Registry](https://www.docker.com/products/docker-registry)

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

The [Build Section](#project-section) in [Turtle File](#full-turtle-file) defines the build packages files to be created

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
- [ ] SSH

Deploying to Nexus
Repositories are defined in the [Turtle File](#full-turtle-file).

Deployment type available and roadmap

Deployment Commands and subcommands Available
- [x] [Sonatype Nexus](http://www.sonatype.org/nexus/)
    - deploy2
        - nexus -> Deploys artifact to [Sonatype Nexus](http://www.sonatype.org/nexus/)

##### Help
```
$:> turtle deploy2 nexus -h
usage: turtle deploy2 nexus

Deploy to Nexus

Flags:
  -h, --help                 Show context-sensitive help (also try --help-long and --help-man).
      --path="."             Go project path
      --batch                Runs commands without asking for any input
  -r, --repoId=REPOID        Repository ID
  -g, --generate-pom="true"  Generate POM
  -f, --file=FILE            Package to Deploy
```

- [ ] Git
    - [ ] [Github](http://github.com)
        - deploy2
            - github -> Deploys artifact to [Github](http://github.com)

#### About [Sonatype Nexus](http://www.sonatype.org/nexus/)
Nexus Repository Manager and Nexus Repository Manager OSS manage software components required for development, deployment, and provisioning. If you develop software, the repository manager can help you share those components with other developers and end users. It greatly simplifies the maintenance of your own internal repositories and access to external repositories. With Nexus Repository Manager and Nexus Repository Manager OSS you can completely control access to, and deployment of, every component in your organization from a single location.



#### Deploy to [Sonatype Nexus](http://www.sonatype.org/nexus/)
```
$:> turtle deploy nexus -f $DIST_FILE
```
## Turtle file

Turtle file is the project definition used by turtle to define properties of the project as well as packaging, builds,

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
