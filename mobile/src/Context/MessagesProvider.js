import {createContext, useContext, useEffect, useRef} from "react"
import { UserContext } from "./UserProvider"
import { SERVER } from "../utils/Repo/envirement"
// import { socket } from "../utils/Repo/socket"
const MessagesContext = createContext()

export default function MessagesProvider ({ children }){
    const {username} = useContext(UserContext)
    const WSRef = useRef(null)
    useEffect(()=>{
        if (username){
            console.log("user ready for messages")
            let ws =  new WebSocket(`${SERVER}/ws/${username}`)
            ws.onopen = () =>console.log("user is connected")
            ws.onclose = (e) => console.log("user is diconnected:", e)
            ws.onerror = (err) => console.log("ws error", err)
            WSRef.current = ws

        }
    }   , [username])    
    return (
        <MessagesContext.Provider value={{}}>
            {children}
        </MessagesContext.Provider>
    )
}