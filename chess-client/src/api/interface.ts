import { ChessColors } from "../types"

type Player = {
  "PlayerId": string,
  "Color": ChessColors,
  "RemainingTime": number,
  "LastActivity": number,
  "Joined": boolean
}

export type Games = {
  "gameId": string,
  "isFinished": boolean,
  "player1": Player,
  "player2": Player
  "totalTime": number,
  "isRunning": boolean
}

export type CreateGame = {
  "playerId": string,
  "playerColor": ChessColors,
  "playerTime": number
}

export type JoinGame = {
  "gameId": string,
  "playerId": string
}
