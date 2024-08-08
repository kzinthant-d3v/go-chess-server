import { useEffect, useState } from 'react'

function App() {
  const [socketInstance, setSocketInstance] = useState<WebSocket | null>(null);

  useEffect(() => {
    let socket = new WebSocket('ws://localhost:5000/ws/12131')
    socket.onopen = () => {
      console.log('socket connected')
      if (!socketInstance) setSocketInstance(socket)
    }
    socket.onmessage = (e) => {
      console.log(e.data)
    }
    socket.onclose = () => {
      console.log('socket closed')
    }
    return () => socket.close()
  }, [])
  console.log(socketInstance)

  return (
    <>
      <h1>Hello client</h1>
      <button onClick={() => socketInstance?.send(JSON.stringify('no json'))}>Click to send a message to server</button>
    </>
  )
}

export default App
