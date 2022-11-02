package builds

import (
	"context"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1"
	cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)

func ListBuilds(
	ctx context.Context,
	client cloudbuild.Client,
	projectId string,
	filter string) *cloudbuild.BuildIterator {

	req := &cloudbuildpb.ListBuildsRequest{
		ProjectId: projectId,
		Filter:    filter,
	}

	return client.ListBuilds(ctx, req)
}

func GetBuildById(
	ctx context.Context,
	client cloudbuild.Client,
	projectId string,
	id string) (*cloudbuildpb.Build, error) {

	req := &cloudbuildpb.GetBuildRequest{
		ProjectId: projectId,
		Id:        id,
	}

	return client.GetBuild(ctx, req)
}
