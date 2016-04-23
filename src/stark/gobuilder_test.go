package main

import (
	"testing"
	"model"
	"encoding/json"
	"fmt"
)

const (
	STS_ENDPOINT = "http://pete.production.maint.bulletproof.net/api/v1/awsid"
)

func TestStark(t *testing.T) {
	p := model.Project{
		ArtifactId:"artifact-id",
		GeneratePom:true,
		Name:"gobuilder-test",
		GroupId:"test-group",
		ProjectType:"go",
		Packaging:"tar.gz",
	}
	rSnap := model.Repository{
		Id:"bp-product-snapshots",
		URL:"http://pc-mgmt01.products.bulletproof.net:8081/nexus/content/repositories/snapshots",
		Type: "snapshots",
	}

	rRel := model.Repository{
		Id:"bp-product-releases",
		URL:"http://pc-mgmt01.products.bulletproof.net:8081/nexus/content/repositories/releases",
		Type: "releases",
	}
	p.Repositories = []model.Repository{rSnap, rRel}

	resp, _ := json.MarshalIndent(&p, "", "  ")
	fmt.Println(string(resp))
}