package interfaces

type IGameRepository interface {
	CreateGame() (string, error)
}
