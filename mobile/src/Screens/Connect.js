import { KeyboardAvoidingView, Platform, StyleSheet, View, Text } from "react-native"
import CustomInput from "../components/atoms/CustomInput"
import { darkgreen, lightgreen, windowWidth } from "../Styles/GlobalStyles"
import { useState } from "react"
import CustomButton from "../components/atoms/CustomButton"

const Connect = () =>{
    const [credentail, setCredential ] = useState("")
    const logAccount = () =>{
        console.log(credentail)
    }
    return (
        <KeyboardAvoidingView behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
            style={style.wrapper}>
                <View style={style.container}>
                    <Text style={style.title}>Connect Account</Text>
                    <CustomInput placeholder={"E-mail or username"} onChange={setCredential} />
                    <CustomButton text={"Connect"} onPress={logAccount} style={{width: 250}}/>
                
                </View>
        </KeyboardAvoidingView>
    )
}

const style = StyleSheet.create({
    wrapper: {
        backgroundColor: lightgreen,
        flex: 1,
        alignItems: 'center',
        justifyContent: "center"
    },
    container: {
        gap: 20,
    },
    title: {
        fontSize: 25,
        fontWeight: '500',
        color: darkgreen,
    }
})

export default Connect