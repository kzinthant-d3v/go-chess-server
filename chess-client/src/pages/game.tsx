import { useParams } from 'react-router-dom'

function Game() {
  const gameId = useParams()?.id
  console.log(gameId)
  return (
    <div>GameId</div>
  )
}

export default Game
