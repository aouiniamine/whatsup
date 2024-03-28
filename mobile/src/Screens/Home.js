import { FlatList } from "react-native"
import NavBar from "../components/atoms/NavBar"
import Message from "../components/Home/Message"
import { useEffect } from "react"
import { getToken, removeToken } from "../utils/Repo/secureStorage"

const Home = () =>{

    const dummyChats = [
        {message: "Heyyyyy!ddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd", name: "Jhon Doe", image: ""}, 
        {message: "Heyy I'd love tooo!!", name: "Diana Doe", image: ""}
    ]
    // useEffect(()=>{removeToken()}, [])
    return (
        <>
            <NavBar></NavBar>
            <FlatList data={dummyChats}
            renderItem={({item, i}) => <Message key={i} name={item.name} message={item.message}/>}/>
        </>
    )
}

export default Home