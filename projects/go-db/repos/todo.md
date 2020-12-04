/* TODO: Examples */
func (repository *PlayerRepository) GetPlayerByName(name string) (dao.PlayerModel, error) {

	row, err :=repository.Query(fmt.Sprintf("SELECT * FROM player_models WHERE name = '%s'", name))
	if err != nil {
		return dao.PlayerModel{}, err
	}

	var player dao.PlayerModel

	row.Next()
	row.Scan(&player.Id, &player.Name, &player.Score)

	return player, nil
}

//TODO: use this later
func (repository *PlayerRepositoryWithCircuitBreaker) GetPlayerByName(name string) (dao.PlayerModel, error) {

	output := make(chan dao.PlayerModel, 1)
	hystrix.ConfigureCommand("get_player_by_name", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_player_by_name", func() error {

		player, _ := repository.PlayerRepository.GetPlayerByName(name)

		output <- player
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return dao.PlayerModel{}, err
	}
}
