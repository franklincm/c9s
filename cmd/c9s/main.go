package main

import (
	"c9s/pkg/builds"
	"context"
	"fmt"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1"
	"google.golang.org/api/iterator"
	cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
	"gopkg.in/yaml.v3"
)

func check_err(err error) {
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
}

type Step struct {
	Id         string
	Name       string
	Entrypoint string
	Args       []string
	Env        []string
}

type Config struct {
	projectId string
	region    string
	filter    string
}

func NewStep(s *cloudbuildpb.BuildStep) Step {
	return Step{
		Id:         s.Id,
		Name:       s.Name,
		Entrypoint: s.Entrypoint,
		Args:       s.Args,
		Env:        s.Env,
	}
}

func PrintBuilds(ctx context.Context, client cloudbuild.Client) {
	buildsList := builds.ListBuilds(ctx, client, "MY_PROJECT", `status=SUCCESS`)

	for i := 0; i < 20; i++ {
		resp, err := buildsList.Next()
		if err == iterator.Done {
			break
		}

		fmt.Printf("%-40s%-30s%s\n", resp.Id, resp.Substitutions["TRIGGER_NAME"], resp.Status.String())
	}
}

func PrintSteps(ctx context.Context, client cloudbuild.Client) {
	testBuild, err := builds.GetBuildById(ctx, client, "MY_PROJECT", "9830acaf-7a51-4a64-a085-451b0a5bc7cb")
	check_err(err)

	for _, step := range testBuild.Steps {
		t := NewStep(step)

		out, _ := yaml.Marshal(t)
		fmt.Println()
		fmt.Println(string(out))
	}
}

func main() {
	ctx := context.Background()
	client, err := cloudbuild.NewClient(ctx)
	check_err(err)

	PrintBuilds(ctx, *client)
	PrintSteps(ctx, *client)

	fmt.Println(cloudbuildpb.Build_Status_name)
}
