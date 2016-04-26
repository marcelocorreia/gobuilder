# Turtle

- [TL;DR](#tldr)
- [About Turtle](#about-turtle)
- [Turtle file](#turtle-file)
- [Turtle config file](#turtle-config-file)
- [Turtle and GO](#turtle-and-Go)
- [Packing Stuff with Turtle](https://github.com/marcelocorreia/turtle#packing-stuff-with-turtle)
- [Deploying Stuff with Turtle](https://github.com/marcelocorreia/turtle#deploying-stuff-with-turtle)
- [Project Types](#project-types)



##About Turtle
Turtle is a build and deployment helper tool where goodies will be implemented in a turtle pace.
It is trying to bring some of the concepts from Maven to other platforms and help when possible. Not too ambitious just filling a gap on BAU operations.
The name Turtle was picked as the previous name was already taken and the lack of creativity at 3am, make me pick up the first cool animal picture I could find by browsing any random stuff, nothing fancy about it. Besides turtles are cool
If anyone can bothered to help with ideas for a logo, please don't be shy. Would be cool to give a face to the Turtle. #giveturtleaface
Also feel free to fork or send pull requests.

### TL;DR
#### Turtle in a nutshell
- [x] Helper tool writen in go to make devs and devops life easier.
- [x] Artifact Management via [Sonatype Nexus](http://www.sonatype.org/nexus/)
- [x] Constant crazy development pace atm. (Things are moving quickly, mess still around).
- [x] Bunch of helpers
- [ ] Auto generation of Documentation

#### TODO
Appart from lists below

- [ ] Clean the mess
- [ ] Unit tests


#### Project types
- [x] [Go](https://golang.org)
- [x] Static directory based projects
    - AngularJS
    - Perl
    - NodeJS Applications
    - Template projects (ansible, terraform, packer)
- [ ] NPM Packages
- [ ] Ruby Gems
- [ ] Docker Containers

#### Golang
- [x] Dependency management - via [gb](https://getgb.io)
    - [ ] Other options i.e. ([Glide](https://glide.sh) or some other). gb has done the job well so far. but with Go 1.6+ out, it might be a good idea to explore what's around
- [x] Build helpers
    - [x] gb build wrapper
    - [x] gb test helper

#### Packaging
- [x] [Distribution packaging](https://github.com/marcelocorreia/turtle#packing-stuff-with-turtle)
  - [x] tar.gz type based on [Turtle File](#full-turtle-file)

#### Installers
 - [x] install [gb](https://getgb.io)
 - [ ] [Hashicorp](http://www.hashicorp.com) stuff
 - [ ] [Ansible](http://www.ansible.org) stuff

#### CI / SCM Tools & helpers
- [ ] SCM Release

#### Deployment Tools & helpers
- [x] [Sonatype Nexus](http://www.sonatype.org/nexus/) deployment
- [ ] [Github](http://www.github.com) deployment
- [ ] Git deployment
- [ ] SSH deployment
- [ ] Docker deployment

#### Documentation generation
- [ ] Hugo integration
- [ ] Confluence sync
- [ ] markdown2confluence integration

#### Account Management
- [ ] Confluence
- [ ] Jira
- [ ] Nexus
- [ ] Jenkins
- [ ] SSH

#### Integration with cool toys
- [ ] Hashicorp Ecosystem
    - [ ] [Consul](https://consul.io)
    - [ ] [Packer](https://packer.io)
    - [ ] [Vault](https://vaultproject.io)
    - [ ] [Nomad](https://nomadproject.io)
    - [ ] [Otto](https://ottoproject.io)
    - [ ] [Vagrant](https://vagrantup.com)
- [ ] [Ansible](http://www.ansible.org)
- [ ] [Docker Family](https://www.docker.com)
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

## Packing Stuff with Turtle

### GO builds

Turtle generates tarball packages using the definitions in [turtle.json](#full-turtle-file)

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
    - deploy2 nexus - Deploys artifact to [Sonatype Nexus](http://www.sonatype.org/nexus/)
- [ ] Git
    - [ ] deploy2 git - Deploys artifact to [Github](http://github.com)
    - [ ] deploy2 github - Deploys artifact to [Github](http://github.com)
    - [ ] deploy2 gitblit - Deploys artifact to [Gitblit](http://gitblit.com)

##### Help deploy2 nexus
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



#### About [Sonatype Nexus](http://www.sonatype.org/nexus/)
Nexus Repository Manager and Nexus Repository Manager OSS manage software components required for development, deployment, and provisioning. If you develop software, the repository manager can help you share those components with other developers and end users. It greatly simplifies the maintenance of your own internal repositories and access to external repositories. With Nexus Repository Manager and Nexus Repository Manager OSS you can completely control access to, and deployment of, every component in your organization from a single location.



#### Deploy to [Sonatype Nexus](http://www.sonatype.org/nexus/)
```
$:> turtle deploy nexus -f $DIST_FILE -repoId 
```
or
```
$:> turtle deploy nexus -b build-id1, build-id2
```

## Turtle Home
Turtle needs home, by default it creates his lair at $HOME/.turtle | %HOME%/.turtle
You can override this by setting the Environment Variable TURTLE_HOME, FROM now this doc will refer it as TURTLE_HOME

## Turtle Config file
The configuration file is placed at TURTLE_HOME/config.json

So far its only making use of Repositories section.

You can define your repositories in this config and it will be available for all jobs.
Repositories can be also defined on the [Turtle File](#full-turtle-file) of each job.

> NOTE: Repositories define in the [Turtle File](#full-turtle-file) will override repositories on [Turtle Config File](#turtle-config-file)


```
{
  "external-apps": [],
  "accounts": [],
  "repositories": [
    {
      "id": "default",
      "type": "snapshots",
      "url": "http://myrepo:8081/nexus/content/repositories/snapshots",
      "build-type": "snapshots"
    },
    {
      "id": "my-company-releases",
      "type": "releases",
      "url": "http://myrepo:8081/nexus/content/repositories/releases",
      "build-type": "releases"
    }
  ]
}
```

#### Model Structure.
```
type Repository struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	User      string `json:"user,omitempty"`
	Password  string `json:"password,omitempty"`
	BuildType string `json:"build-type"`
}

type TurtleConfig struct {
	ExternalApps []ExternalApp `json:"external-apps,ommitempty"`
	Accounts     []Account    `json:"accounts,ommitempty"`
	Repositories []Repository `json:"repositories,ommitempty"`
}

type Account struct {
	Type     string `json:"account-type"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Endpoint string `json:"endpoint,ommitempty"`
}

type ExternalApp struct {
	Available           bool     `json:"available"`
	Executable          string   `json:"executable"`
	Mandatory           bool     `json:"mandatory"`
	Name                string   `json:"name"`
	SupportedVersions   []string `json:"supported-versions"`
	VersionCheckCommand []string `json:"version-check-command"`
	DownloadURL         string   `json:"download-url,ommitempty"`
}
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

Turtle started as Golang helper tool and quickly started shifting into a more generic and smarter build helper.
It offers a wrapper for gb https://getgb.io. Other Go build tool might be integrated but for the moment only gb is available.

gb is an awesome tool for Golang projects and provides features like:
- Project based workflow
- Automatic project detection
- Dependency management
- Test run

An alternative build tool for the [Go programming language](https://golang.org/).

[Read more about the rationale for gb](https://getgb.io/rationale).

Project-Based
gb operates on the concept of a project. A gb project is a workspace for all the Go code that is required to build your project.

A gb project is a folder on disk that contains a sub directory named src/. That’s it, no environment variables to set. For the rest of this document we’ll refer to your gb project as $PROJECT.

You can create as many projects as you like and move between them simply by changing directories.

[Read more about setting up a gb project](https://getgb.io/docs/project).
