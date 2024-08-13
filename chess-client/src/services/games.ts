import { api, socketApi, types } from "../api";

export const createGame = async (createGame: types.CreateGame) => {
  return api.createGameApi(createGame)
}

export const listGames = async () => {
  return api.listGamesApi()
}

export const joinGame = async (joinGame: types.JoinGame) => {
  return api.joinGameApi(joinGame)
}

export const getConnectWebSocket = (joinGame: types.JoinGame) => {
  return socketApi.getConnectWebSocketApi(joinGame)
}

export const closeWebSocket = (socket: WebSocket) => {
  return socketApi.closeWebSocketApi(socket)
}
