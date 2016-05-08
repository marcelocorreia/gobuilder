package model

type Project struct {
	GroupId       string              `json:"group-id"`
	ArtifactId    string              `json:"artifact-id"`
	Name          string              `json:"name"`
	ProjectType   string              `json:"project-type"`
	Version       string              `json:"version"`
	File          string              `json:"file,ommitempty"`
	Packaging     string              `json:"packaging"`
	GeneratePom   bool                `json:"generate-pom,ommitempty"`
	Repositories  []Repository        `json:"repositories,ommitempty"`
	Builds        []Build             `json:"builds,ommitempty"`
	VersionString string              `json:"version-string,ommitempty"`
}

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

type Build struct {
	OS   string `json:"os,ommitempty"`
	Arch string `json:"arch,ommitempty"`
	Type string `json:"type"`
	ID   string `json:"id"`
}