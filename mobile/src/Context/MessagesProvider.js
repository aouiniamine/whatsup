import {createContext, useState, useEffect, useRef} from "react"
import { SERVER } from "../utils/Repo/envirement"
import { getToken } from "../utils/Repo/secureStorage"
import { getUser } from "../utils/Repo/Auth"
import { useDispatch } from "react-redux"
import { addMessage } from "../utils/Redux/reducer"
// import { socket } from "../utils/Repo/socket"
export const MessagesContext = createContext()

export default function MessagesProvider ({ children }){
    const [token, setToken ] = useState(null)
    const [username, setUsername] = useState(null)
    const WSRef = useRef(null)
    const [isOpen, setIsOpen] = useState(false)
    const dispatch = useDispatch()
    const recieveMessage = (message) =>{
        
        dispatch(addMessage(message))
    }

    const sendMessage = (username, message) =>{
        if (isOpen){
            WSRef.current.send(JSON.stringify({username, message}))
        }
    }
    useEffect(()=>{

        (async()=>{
            const tokenFound = token || await getToken()
            if (tokenFound){
                try{

                    const currentUser = await getUser(tokenFound)
                    setUsername(currentUser.username)
                    setToken(tokenFound)
                    let ws =  new WebSocket(`${SERVER}/ws/${currentUser.username}?token=${tokenFound}`,)
                    ws.onopen = () =>{
                        setIsOpen(true)
                        console.log("user is connected")
                    }
                    ws.onclose = (e) => {
                        setIsOpen(false)
                        console.log("user is diconnected:", e)
                        
                    } 
                    ws.onerror = (err) => console.log("ws error", err)
                    ws.onmessage = (event) =>{
                        const messageData = JSON.parse(event.data)
                        switch(messageData.type){
                            case "message:recieve":
                                const {username, message} = messageData
                                
                                console.log("message from", username, "is recieved")
                                recieveMessage({message, username})
                                break;
                            default:
                                console.log("Still working on it!")
                        }
                        
                    }
                    ws.addEventListener('recieve:message', (e)=> console.log("revieved message: ", e.data))
                    
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
        <MessagesContext.Provider value={{setToken, sendMessage, username}}>
            {children}
        </MessagesContext.Provider>
    )
}