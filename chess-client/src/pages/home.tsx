import { Outlet } from 'react-router-dom'
import PlayerProvider from '../context/PlayerProvider'

function Home() {
  return (
    <div>
      <PlayerProvider>
        <Outlet />
      </PlayerProvider>
    </div>
  )
}

export default Home
