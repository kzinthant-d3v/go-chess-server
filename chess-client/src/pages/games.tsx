import { useEffect, useState } from "react";
import type { Games } from "../api/interface";
import { createGame, joinGame, listGames } from "../services";
import { useNavigate } from "react-router-dom";
import { usePlayer } from "../context/PlayerProvider";

function Games() {
  const [currentGames, setCurrentGames] = useState<Games[]>([]);
  const { playerId, updatePlayerId } = usePlayer();
  const navigate = useNavigate()
  const areCurrentGamesEmpty = currentGames.length === 0;

  useEffect(() => {
    (async () => {
      setCurrentGames(await listGames());
    })();
  }, []);

  const createGameSubmit = async () => {
    if (!playerId) return;
    const createdGames = await createGame({
      playerId,
      playerColor: "white",
      playerTime: 11111
    })
    setCurrentGames(createdGames)
  };

  const joinAGame = async (gameId: string) => {
    if (!playerId) return;

    await joinGame({
      gameId,
      playerId
    })

    navigate(`/game/${gameId}`)
  }


  return (
    <div>
      <form>
        <input name="playerId" onChange={(e) => updatePlayerId(e.target.value)} />
        <button onClick={createGameSubmit}>Create a test game</button>
        <div>Games</div>
      </form>
      {!areCurrentGamesEmpty &&
        currentGames.map((game) => (
          <div key={game.gameId}>
            <div>
              <div>
                {game.gameId}
              </div>
              <div>
                <h1>Is game running</h1>
                {game.isRunning}
                <button onClick={() => joinAGame(game.gameId)}>Join</button>
              </div>
            </div>
          </div>
        ))}
    </div>
  );
}

export default Games;
