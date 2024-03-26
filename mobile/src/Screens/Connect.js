import { KeyboardAvoidingView, Platform, StyleSheet, View, Text } from "react-native"
import CustomInput from "../components/atoms/CustomInput"
import { darkgreen, lightgreen, windowWidth } from "../Styles/GlobalStyles"
import { useEffect, useState } from "react"
import CustomButton from "../components/atoms/CustomButton"
import { connect } from "../utils/Repo/Auth"
import { getToken } from "../utils/Repo/secureStorage"

const Connect = ({navigation}) =>{
    const [credential, setCredential ] = useState("")
    const [error, setError] = useState('')
    const register = ()=> navigation.navigate("Register")
    const logAccount = async() =>{
        try{
            const result = await connect(credential)
            // assuming result's format {email: '...'}
            navigation.navigate("Validate", result)
            
        }catch(err) { console.log('error: ', err); setError(err.message)}
        
    }

    useEffect(()=> {
        if (getToken()){
            navigation.navigate("Home")
        }
    }, [])
    return (
        <KeyboardAvoidingView behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
            style={style.wrapper}>
                <View style={style.mainContainer}>

                    <View style={style.container}>
                        <Text style={style.title}>Connect Account</Text>
                        <CustomInput placeholder={"E-mail or username"} onChange={setCredential} />
                        <CustomButton text={"Connect"} onPress={logAccount} style={{width: 250}}/>
                    </View>
                    <Text onPress={register} style={style.register}>create an account ?</Text>
                    {error && <Text style={style.error} >{error}</Text>}
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
    mainContainer: {
        gap: 10
    },
    container: {
        gap: 20,
    },
    title: {
        fontSize: 25,
        fontWeight: '500',
        color: darkgreen,
    },
    error: {
        
        textAlign: "center",
        color: 'white',
        fontSize: 19
    },
    register: {
        textAlign: "center",
        color: "grey",
        fontSize: 19
    }
})

export default Connect