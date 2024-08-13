import { useParams } from 'react-router-dom'
import { usePlayer } from '../context/PlayerProvider';
import { useEffect, useRef, useState } from 'react';
import { closeWebSocket, getConnectWebSocket } from '../services';

function Game() {
  const gameId = useParams()?.id ?? "";
  const [test, setTest] = useState("")
  const { playerId } = usePlayer();
  const socketRef = useRef<WebSocket | null>(null)

  useEffect(() => {
    socketRef.current = getConnectWebSocket({ gameId, playerId })
    const socket = socketRef.current

    socket.onmessage = (event) => {
      console.log(event.data)
    }
    socket.onopen = () => {
      console.log("Connected")
    }
    socket.onclose = () => {
      console.log("Disconnected")
    }
    socket.onerror = (error) => {
      console.log("Error", error)
    }

    return () => closeWebSocket(socket)

  }, []);
  return (
    <div>GameId: {gameId}, PlayerId: {playerId}
      <input value={test} onChange={(e) => setTest(e.target.value)} />
      <button onClick={() => socketRef.current?.send(test)}>Send something</button>
    </div>
  )
}

export default Game
