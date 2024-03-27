import {createContext, useEffect, useState} from "react"
import { getToken } from "../utils/Repo/secureStorage"
import { getUser } from "../utils/Repo/Auth"
export const UserContext = createContext()

export default function UserProvider ({ children }){
    const [token, setToken] = useState(null)
    const [username, setUsername] = useState(null)
    useEffect(()=>{
        (async()=>{
            const tokenFound = token || await getToken()
            if (tokenFound){
                try{

                    const user = await getUser(tokenFound)
                    
                    setUsername(user.username)
                    setToken(tokenFound)
                } catch (err){
                    console.log(err)
                }
            }
        })()
        
    }   , [])

    return (
        <UserContext.Provider value={{token, username, setToken}}>
            {children}
        </UserContext.Provider>
    )
}