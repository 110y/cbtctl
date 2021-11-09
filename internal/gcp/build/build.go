package build

import (
	"context"
	"fmt"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1"
	cloudbuildpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)

func Run(ctx context.Context, project, trigger string, substitutions map[string]string) error {
	cli, err := cloudbuild.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create cloud build client: %w", err)
	}

	req := &cloudbuildpb.RunBuildTriggerRequest{
		ProjectId: project,
		TriggerId: trigger,
		Source: &cloudbuildpb.RepoSource{
			Substitutions: substitutions,
		},
	}

	_, err = cli.RunBuildTrigger(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to run trigger: %w", err)
	}

	return nil
}
