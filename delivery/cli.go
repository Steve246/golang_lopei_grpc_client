package delivery

import (
	"fmt"
	"golang_lopei_grc_client/config"
	"golang_lopei_grc_client/manager"
)

type cli struct {
	useCaseManager manager.UseCaseManager
}

func (c *cli) Run() {
	balance, err := c.useCaseManager.CheckBalanceUsecase().GetBalance(int32(4))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(balance)
}


func Cli() *cli {
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)

	useCaseManager := manager.NewUseCaseManager(repoManager)

	return &cli{useCaseManager}
}