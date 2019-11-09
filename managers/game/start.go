package game

import (
	"fmt"

	"github.com/minhajuddinkhan/webrung/entities"
	"github.com/minhajuddinkhan/webrung/iorpc"
)

func (gm *gameManager) StartGame(gameID string, startBy *entities.Player) (bool, error) {

	var started bool
	gameToStart, err := gm.store.GetGame(gameID)
	if err != nil {
		return started, err
	}

	if gameToStart.GetHostID() != startBy.ID {
		return started, fmt.Errorf("game cannot be started by someone other than the host")
	}

	players, err := gm.store.GetPlayersInGame(gameID)
	if err != nil {
		return started, err
	}

	if len(players) != 4 {
		return started, fmt.Errorf("cannot start game until 4 players have joined")
	}

	playerIds := make([]string, len(players))
	for i, p := range players {
		playerIds[i] = p.GetPlayerID()
	}

	started, err = gm.ioclient.StartGame(iorpc.DistributeCardsRequest{
		PlayerIds: playerIds,
		GameID:    gameID,
	})

	return started, err
}
