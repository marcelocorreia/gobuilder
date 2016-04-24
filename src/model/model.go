package model

type Project struct {
	GroupId      string              `json:"group-id"`
	ArtifactId   string              `json:"artifact-id"`
	Name         string              `json:"name"`
	ProjectType  string              `json:"project-type"`
	Version      string              `json:"version"`
	File         string              `json:"file"`
	Packaging    string              `json:"packaging"`
	GeneratePom  bool                `json:"generate-pom"`
	Repositories []Repository        `json:"repositories"`
	Builds       []Build         	 `json:"builds"`
}

type Repository struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type Build struct {
	OS   string
	Arch string
}