# webrung

### Application workflow:

1. Initial Application status:
    * The players are already hardcoded in the database for seed data. No creation of players is required at this point in time.
2. Player login: 
    * When a player logs in to the game, a token is given back to the user.
    * This token in saved in a redis (`key=token+gameID`, `value=playerWithCards`)
3. Hosting a Game:
    * A new game is created on the game database (game_history)
    * `playersJoined` should be incremented by one.
    * At this point, game should have one player linked to it, which is the host.
    * Host player has no cards right now
    * The game record created in the database should have a deck, i.e to be distributed.
    * The game record created in the database should have the state of the hand being played
    * The game record should have the turn thats being played
    * The game record should maintain the state of the hand being played.
    * The web application should establish a socket connection with the i/o server with the token provided
    * The web application stores game id and token in socket redis against socket id
    * The web application should put the cards against the token
4. Joining a Game:
    * the dashboard should show a list of hosted games.
    * `playersJoined` increments by one
    * the web server checks 
    * the game record should maintain which player has joined.
    * another player is added in the redis with no cards
    * 
5. Disconnecting/Rejoining a Game:
    * if a player disconnects, its socket id should be searched in the socket redis.
    * it reveals the game id, and the player it was linked to.
    * a broadcast message is sent to all other players that a player X has disconnected
    * when the player reconnects with game id and token, it is authenticated by verifying in the cards redis
    * if verified, the socket id is stored back in the socket redis


## Todos:
1. Architecture of the game database. (Normalized)
2. Figure out a strategy to broadcast players to the disconnected players.

### Links

| Name | Resource |
| ------ | ------ |
| Architecture Diagram | <link> |


# Test

## Integration

```bash
$ sudo docker-compose -f docker-compose-test.yml up --abort-on-container-exit
```
