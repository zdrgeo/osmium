package github

import (
	"errors"
	"log"

	"github.com/zdrgeo/osmium/internal/analysis"

	"github.com/cli/go-gh/v2/pkg/api"
)

type CommitAnalysisSource struct {
	client          *api.GraphQLClient
	repositoryOwner string
	repositoryName  string
}

func NewCommitAnalysisSource(client *api.GraphQLClient, repositoryOwner, repositoryName string) *CommitAnalysisSource {
	return &CommitAnalysisSource{client: client, repositoryOwner: repositoryOwner, repositoryName: repositoryName}
}

func (source *CommitAnalysisSource) Query() *analysis.Analysis {
	log.Fatal(errors.New("not implemented"))

	return nil
}
