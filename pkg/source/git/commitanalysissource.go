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

func (source *CommitAnalysisSource) Query() *analysis.Analysis {
	log.Fatal(errors.New("not implemented"))

	return nil
}
