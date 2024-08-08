import { BACKEND_URL, ws } from "./config"
import { JoinGame } from "./interface"

export const getConnectWebSocketApi = ({ gameId, playerId }: JoinGame) => {
  if (WebSocket.OPEN) throw new Error('A WebSocket is already connected')

  const params = new URLSearchParams({ gameId, playerId }).toString()
  const socket = new WebSocket(`${ws}${BACKEND_URL}?${params}`)
  return socket
}

export const closeWebSocketApi = (socket: WebSocket) => {
  socket.close(1000, 'Unjoined')
}
