import { createContext, useCallback, useContext, useState } from "react"

type PlayerContextType = {
    playerId: string
    updatePlayerId: (id: string) => void
}

const PlayerContext = createContext<PlayerContextType>({} as PlayerContextType)

function PlayerProvider({ children }: { children: React.ReactNode }) {
    const [playerId, setPlayerId] = useState<string>("")

    const updatePlayerId = useCallback((playerId: string) => setPlayerId(playerId), [])

    return (
        <PlayerContext.Provider value={{ playerId, updatePlayerId }}>
            {children}
        </PlayerContext.Provider>
    )
}

export const usePlayer = () => {
    const context = useContext(PlayerContext)
    if (!context) {
        throw new Error("usePlayer must be used within a PlayerProvider")
    }
    return context
}

export default PlayerProvider