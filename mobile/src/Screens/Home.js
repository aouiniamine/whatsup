import { FlatList, StyleSheet, TouchableOpacity, } from "react-native"
import NavBar from "../components/atoms/NavBar"
import Message from "../components/Home/Message"
import { useContext, useEffect, useState } from "react"
import { FontAwesome } from '@expo/vector-icons';
import { darkgreen, lightgreen } from "../Styles/GlobalStyles"
import SendMessage from "../components/atoms/SendMessage"
import { MessagesContext } from "../Context/MessagesProvider";

const Home = ({navigation}) =>{
    const [sendingMesssage, setSendingMessage] = useState(false)
    const closeForm = ()=> setSendingMessage(false)
    const openForm = ()=>setSendingMessage(true)
    const {chats} = useContext(MessagesContext)
    const dummyChats = [
        {message: "Heyyyyy!ddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd", name: "Jhon Doe", image: ""}, 
        {message: "Heyy I'd love tooo!!", name: "Diana Doe", image: ""}
    ]
    // useEffect(()=>{removeToken(); navigation.navigate("Connect")}, [])
    const renderChat = ({item, i}) =>{
        const lastMessage = item[item.length-1]
        console.log(lastMessage)
        return <Message key={i} name={lastMessage.username} message={lastMessage.message}/>
    }

    // useEffect(()=>{console.log(chats)}, [chats])
    return (
        <>

            <NavBar></NavBar>
            <TouchableOpacity style={[style.addContainer, sendingMesssage ? {backgroundColor: lightgreen}: null]} onPress={openForm}>
                <FontAwesome name="plus" size={40} color={"white"} />

            </TouchableOpacity>
            {sendingMesssage && <SendMessage isVisible={sendingMesssage} closeForm={closeForm}/>}
            <FlatList data={chats} renderItem={renderChat}/>
        </>
    )
}

const style = StyleSheet.create({
    addContainer: {
        backgroundColor: darkgreen,
        flex: 1,
        width: 60,
        maxHeight: 60,
        minHeight: 60,
        borderRadius: 100,
        position: "absolute",
        bottom: 50,
        right: 50,
        alignItems: 'center',
        justifyContent: "center",
        zIndex: 200
    }
})
export default Home