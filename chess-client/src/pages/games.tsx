import { FormEventHandler, useEffect, useState } from "react";
import type { Games } from "../api/interface";
import { listGames } from "../services";

function Games() {
  const [currentGames, setCurrentGames] = useState<Games[]>([]);
  const areCurrentGamesEmpty = currentGames.length === 0;

  useEffect(() => {
    (async () => {
      setCurrentGames(await listGames());
    })();
  }, []);

  const createGame: FormEventHandler<HTMLFormElement> = (element) => {
    element.preventDefault();
    const formData = new FormData(element.currentTarget);
    console.log(formData.get("playerId"));
  };
  return (
    <div>
      <form onSubmit={createGame}>
        <input name="playerId" />
        <button>Create a test game</button>
        <div>Games</div>
      </form>
      {!areCurrentGamesEmpty &&
        currentGames.map((game) => (
          <div key={game.gameId}>
            <div>
              <div>
                <h1>Game Id</h1>
                {game.gameId}
              </div>
              <div>
                <h1>Is game running</h1>
                {game.isRunning}
              </div>
            </div>
          </div>
        ))}
    </div>
  );
}

export default Games;
