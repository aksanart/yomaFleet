package repository

import (
	"fmt"

	"github.com/aksanart/tracker_service/pkg/config"
	repointerface "github.com/aksanart/tracker_service/repository/repo_interface"
)

type RepoConf interface {
	Init(*Repository) error
	GetRepoName() string
}

func (r *Repository) PrintHealthy(repoName string) {
	fmt.Printf("+ %s Repository is healthy! \n", repoName)
}

func (r *Repository) PrintNotHealthy(repoName string) {
	fmt.Printf("- %s Repository is not healthy!\n", repoName)
}

func (r *Repository) PrintSkipHealthy(repoName string) {
	fmt.Printf("- %s Repository Healthy skipped due config= false!\n", repoName)
}

var repositoriesPointer *Repository

func NewRepository(rf []RepoConf) (*Repository, error) {
	if repositoriesPointer != nil {
		return repositoriesPointer, nil
	}
	repositoriesPointer = &Repository{}
	for _, rc := range rf {
		err := rc.Init(repositoriesPointer)
		if !config.GetConfig("check_healthy_repo").GetBool() {
			repositoriesPointer.PrintSkipHealthy(rc.GetRepoName())
			return repositoriesPointer, nil
		}
		if err != nil {
			repositoriesPointer.PrintNotHealthy(rc.GetRepoName())
			return nil, err
		}
		repositoriesPointer.PrintHealthy(rc.GetRepoName())
	}
	return repositoriesPointer, nil
}

type Repository struct {
	MongoDb MongoCollections
	Redis   repointerface.RedisInterface
}

type MongoCollections struct {
	Tracker repointerface.MongoInterface
}
