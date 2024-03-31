import { StyleSheet, View, Text } from "react-native"
import { Avatar } from 'react-native-elements';
import { nameStyle, windowWidth } from "../../Styles/GlobalStyles";
import { useEffect, useState } from "react";
const Message = ({message, name, image,}) =>{
    const [showMessage, setShowMessage] = useState('')
    useEffect(()=>{
        if (message.length > 27){
            const msg = message.slice(0, 27).concat("...")
            setShowMessage(msg)
        } else setShowMessage(message)
    }, [message])
    return (
        <View style={style.container}>
            <Avatar
                rounded
                size={70}
                source={{
                    uri:
                    'https://s3.amazonaws.com/uifaces/faces/twitter/ladylexy/128.jpg',
                }}
            />
            <View style={{marginLeft: 20, marginRight: 90, marginTop: 15}}>
                <Text style={nameStyle}>{name}</Text>
                <Text style={style.message}>{showMessage}</Text>

            </View>

        </View>
    )
}

const style = StyleSheet.create({
    container: {
        marginTop: 20,
        marginLeft: 20,
        flex: 1,
        flexDirection: "row",
        paddingVertical: 10
    },
    message: {

        color: "grey",
        fontSize: 16,
        fontWeight: "400"
    }

})

export default Message