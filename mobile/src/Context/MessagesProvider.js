import {createContext, useState, useEffect, useRef} from "react"
import { SERVER } from "../utils/Repo/envirement"
import { getToken } from "../utils/Repo/secureStorage"
import { getUser } from "../utils/Repo/Auth"
// import { socket } from "../utils/Repo/socket"
export const MessagesContext = createContext()

export default function MessagesProvider ({ children }){
    const [token, setToken ] = useState(null)
    const [username, setUsername] = useState(null)
    const WSRef = useRef(null)
    useEffect(()=>{

        (async()=>{
            const tokenFound = token || await getToken()
            if (tokenFound){
                try{

                    const currentUser = await getUser(tokenFound)
                    
                    setUsername(currentUser.username)
                    setToken(tokenFound)
                    let ws =  new WebSocket(`${SERVER}/ws/${currentUser.username}?token=${tokenFound}`,)
                    ws.onopen = () =>console.log("user is connected")
                    ws.onclose = (e) => console.log("user is diconnected:", e)
                    ws.onerror = (err) => console.log("ws error", err)
                    WSRef.current = ws
                } catch (err){
                    console.log(err)
                }
            } else {
                console.log("user not logged")
            }
        })()
        

        
    }   , [token])    
    return (
        <MessagesContext.Provider value={{setToken}}>
            {children}
        </MessagesContext.Provider>
    )
}