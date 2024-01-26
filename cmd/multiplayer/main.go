package main

import (
	"fmt"
	"sync"
	"time"
)

// main provides an example for how to use 'sync.Cond' with broadcasting.
func main() {
	cond := sync.NewCond(&sync.Mutex{})
	waitingForPlayers := 4
	for playerId := 0; playerId < 4; playerId++ {
		go playerHandler(cond, &waitingForPlayers, playerId)
		time.Sleep(1 * time.Second)
	}
}

func playerHandler(cnd *sync.Cond, pendingPlayers *int, playerId int) {
	cnd.L.Lock()
	fmt.Println(playerId, ": Connected")
	*pendingPlayers--
	for *pendingPlayers > 0 {
		fmt.Println(playerId, ": Waiting for more players")
		cnd.Wait()
	}
	cnd.Broadcast()
	cnd.L.Unlock()
	fmt.Println("All players connected. Ready player", playerId)
}
