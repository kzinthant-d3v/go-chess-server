import { BACKEND_URL, http } from "./config";
import { CreateGame, Games, JoinGame } from "./interface";

export const createGameApi = async ({
  playerId,
  playerColor,
  playerTime,
}: CreateGame): Promise<Games[]> => {
  const response = await fetch(`${http}${BACKEND_URL}/create-game`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      playerId,
      playerColor,
      playerTime,
    }),
  });
  return (await response.json()) as Games[];
};

export const listGamesApi = async (): Promise<Games[]> => {
  const response = await fetch(`${http}${BACKEND_URL}/list-games`);
  const gameData = await response.json();
  if (gameData) return gameData.gameList as Games[];
  return [];
};

export const joinGameApi = async ({ gameId, playerId }: JoinGame) => {
  const response = await fetch(`${http}${BACKEND_URL}/join-game`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      gameId,
      playerId,
    }),
  });
  return (await response.json()) as string;
};
