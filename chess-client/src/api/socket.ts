import { BACKEND_URL, ws } from "./config"
import { JoinGame } from "./interface"

export const getConnectWebSocketApi = ({ gameId, playerId }: JoinGame) => {
  const params = new URLSearchParams({ gameId, playerId }).toString()
  const socket = new WebSocket(`${ws}${BACKEND_URL}/connect-game?${params}`)
  return socket
}

export const closeWebSocketApi = (socket: WebSocket) => {
  socket.close(1000, 'Unjoined')
}
