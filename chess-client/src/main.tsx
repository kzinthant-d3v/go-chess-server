import React from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Home from './pages/home.tsx'
import Games from './pages/games.tsx'
import Game from './pages/game.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
    children: [
      {
        path: "games",
        element: <Games />
      },
      {
        path: "game/:id",
        element: <Game />,
        loader: async ({ params }) => {
          return { id: params.id }
        }
      }

    ]
  },
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
