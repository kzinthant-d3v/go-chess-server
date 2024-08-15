import { BACKEND_URL, http } from "./config";
import { jsonFetch } from "./helper";
import { CreateGame, GameList, Games, JoinGame } from "./interface";

export const createGameApi = async ({
  playerId,
  playerColor,
  playerTime,
}: CreateGame): Promise<Games[]> => {
  const gameData = await jsonFetch<GameList>(`${http}${BACKEND_URL}/create-game`, {
    method: "POST",
    body: JSON.stringify({
      playerId,
      playerColor,
      playerTime,
    }),
  });
  if (gameData) return gameData.gameList;
  return [];
};

export const listGamesApi = async (): Promise<Games[]> => {
  const gameData = await jsonFetch<GameList>(`${http}${BACKEND_URL}/list-games`);
  if (gameData) return gameData.gameList as Games[];
  return [];
};

export const joinGameApi = async ({ gameId, playerId }: JoinGame) => {
  const result = await jsonFetch<string>(`${http}${BACKEND_URL}/join-game`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      gameId,
      playerId,
    }),
  });
  return result;
};
