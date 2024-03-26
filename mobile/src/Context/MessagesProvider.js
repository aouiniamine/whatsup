import {createContext, useEffect} from "react"
import { socket } from "../utils/Repo/socket"
const MessagesContext = createContext()

export default function MessagesProvider ({ children }){

    useEffect(()=>{
        socket.on("connect", ()=>{
            console.log("User is connect to web socket!")
        })
        socket.on("disconnect", ()=>{
            console.log("user has disconnected")
        })
        socket.on("connect_error", err => console.log(err, err.stack))
        return ()=>{
            socket.off("connect")
            socket.off("disconnect")
            socket.off("connect_error")
        }
    }, [])

    return (
        <MessagesContext.Provider value={{}}>
            {children}
        </MessagesContext.Provider>
    )
}