package git

import (
	"errors"
	"log"

	"github.com/zdrgeo/osmium/pkg/analysis"
)

type CommitAnalysisSource struct {
}

func NewCommitAnalysisSource() *CommitAnalysisSource {
	return &CommitAnalysisSource{}
}

func (source *CommitAnalysisSource) Query(options map[string]string) *analysis.Analysis {
	log.Fatal(errors.New("not implemented"))

	return nil
}
