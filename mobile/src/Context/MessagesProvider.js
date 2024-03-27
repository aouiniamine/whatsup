import {createContext, useEffect} from "react"
import { socket } from "../utils/Repo/socket"
const MessagesContext = createContext()

export default function MessagesProvider ({ children }){

    useEffect(()=>{
        socket.onopen = () =>console.log("user is connected")
        socket.onclose = (e) => console.log("user is diconnected:", e)
        socket.onerror = (err) => console.log("ws error", err)
    }   , [])

    return (
        <MessagesContext.Provider value={{}}>
            {children}
        </MessagesContext.Provider>
    )
}