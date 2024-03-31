import { FlatList, StyleSheet, TouchableOpacity, } from "react-native"
import NavBar from "../components/atoms/NavBar"
import Message from "../components/Home/Message"
import { useContext, useEffect, useState } from "react"
import { FontAwesome } from '@expo/vector-icons';
import { darkgreen, lightgreen } from "../Styles/GlobalStyles"
import SendMessage from "../components/atoms/SendMessage"
import store from "../utils/Redux/store";
import { connect, useSelector } from "react-redux";
import { messageStateToProps } from "../utils/Redux/reducer";

const Home = ({navigation, lastMessages}) =>{
    const [sendingMesssage, setSendingMessage] = useState(false)
    const closeForm = ()=> setSendingMessage(false)
    const openForm = ()=>setSendingMessage(true)
    // useEffect(()=>{removeToken(); navigation.navigate("Connect")}, [])

    const renderChat = ({item, i}) =>{
        // const lastMessage = convo[convo.length-1]
        return <Message key={i} name={item.username} message={item.message}/>
    }

    return (
        <>

            <NavBar></NavBar>
            <TouchableOpacity style={[style.addContainer, sendingMesssage ? {backgroundColor: lightgreen}: null]} onPress={openForm}>
                <FontAwesome name="plus" size={40} color={"white"} />

            </TouchableOpacity>
            {sendingMesssage && <SendMessage isVisible={sendingMesssage} closeForm={closeForm}/>}
            <FlatList data={lastMessages} renderItem={renderChat}/>
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

export const lastMessagesToProps = state => {
    // passing only what's needed as props
    const messages = new Array(...state.allMessages)
    const lastMessages = messages.map(chat => chat[chat.length-1])
    return {lastMessages}
}
export default connect(lastMessagesToProps)(Home)
// export default Home