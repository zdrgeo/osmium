package git

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/zdrgeo/osmium/pkg/analysis"
)

type CommitAnalysisSource struct {
	progressFunc analysis.GitAnalysisProgressFunc
}

func NewCommitAnalysisSource(progressFunc analysis.GitAnalysisProgressFunc) *CommitAnalysisSource {
	return &CommitAnalysisSource{
		progressFunc: progressFunc,
	}
}

func (source *CommitAnalysisSource) Query(spanSize int, repositoryURL, repositoryPath string) (*analysis.Analysis, error) {
	// cloneOptions := &git.CloneOptions{
	// 	URL:      repositoryURL,
	// 	Progress: os.Stdout,
	// }

	// repository, err := git.PlainClone(repositoryPath, false, cloneOptions)
	repository, err := git.PlainOpen(repositoryPath)
	// repository, err := git.Clone(memory.NewStorage(), nil, cloneOptions)
	// repository, err := git.Open(memory.NewStorage(), nil)

	if err != nil {
		log.Panic(err)
	}

	reference, err := repository.Head()

	if err != nil {
		log.Panic(err)
	}

	// since := time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC)
	// until := time.Date(2200, 1, 1, 0, 0, 0, 0, time.UTC)

	logOptions := &git.LogOptions{From: reference.Hash() /* , Order=LogOrderCommitterTime, Since: &since, Until: &until */}

	commitIter, err := repository.Log(logOptions)

	if err != nil {
		log.Panic(err)
	}

	modules := map[string]*analysis.Module{}
	spans := map[string]*analysis.Span{}
	changes := map[string]*analysis.Change{}
	nodes := map[string]*analysis.Node{}

	err = commitIter.ForEach(func(commit *object.Commit) error {
		changeName := fmt.Sprint(commit.Hash)

		changes[changeName] = &analysis.Change{Name: changeName}

		nodeFileIter, err := commit.Files()

		if err != nil {
			return err
		}

		err = nodeFileIter.ForEach(func(nodeFile *object.File) error {
			if node, ok := nodes[nodeFile.Name]; ok {
				edgeFileIter, err := commit.Files()

				if err != nil {
					return err
				}

				err = edgeFileIter.ForEach(func(edgeFile *object.File) error {
					if edge, ok := node.Edges[edgeFile.Name]; ok {
						edge.ChangeNames = append(edge.ChangeNames, changeName)
					} else {
						node.Edges[edgeFile.Name] = &analysis.Edge{ChangeNames: []string{changeName}}
					}

					return nil
				})

				if err != nil {
					return err
				}
			} else {
				edges := map[string]*analysis.Edge{}

				edgeFileIter, err := commit.Files()

				if err != nil {
					return err
				}

				err = edgeFileIter.ForEach(func(edgeFile *object.File) error {
					edges[edgeFile.Name] = &analysis.Edge{ChangeNames: []string{changeName}}

					return nil
				})

				if err != nil {
					return err
				}

				nodes[nodeFile.Name] = &analysis.Node{Edges: edges}
			}

			return nil
		})

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	span := &analysis.Span{Changes: changes, Nodes: nodes}

	spans[span.Name] = span

	return &analysis.Analysis{Modules: modules, Spans: spans}, nil
}
