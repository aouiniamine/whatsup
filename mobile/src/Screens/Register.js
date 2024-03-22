import { KeyboardAvoidingView, Platform, StyleSheet, View, Text } from "react-native"
import CustomInput from "../components/atoms/CustomInput"
import { darkgreen, lightgreen, windowWidth } from "../Styles/GlobalStyles"
import { useState } from "react"
import CustomButton from "../components/atoms/CustomButton"
import { connect, register } from "../utils/Repo/Auth"
import { validateEmail, validateUsername } from "../utils/general"

const Register = ({navigation}) =>{
    const [email, setEmail ] = useState("")
    const [error, setError] = useState('')
    const [username, setUsername] = useState('')
    const registerAccount = async() =>{
        try{
            const validEmail = validateEmail(email)
            const validUsername = validateUsername(username)
            if(validEmail && validUsername){
                
                const credential = {
                    email, username
                }
                const res = await register(credential)
                console.log(res, credential)
            } else{
                
                console.log("email is valid:", validEmail)
                console.log("username is valid:", validUsername)
                 
            }

        }catch(err){ console.log(err); throw err}
        
    }
    return (
        <KeyboardAvoidingView behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
            style={style.wrapper}>
                <View style={style.mainContainer}>

                    <View style={style.container}>
                        <Text style={style.title}>Register Account</Text>
                        <CustomInput placeholder={"E-mail"} onChange={setEmail} />
                        <CustomInput placeholder={"Username"} onChange={setUsername}/>
                        <CustomButton text={"Connect"} onPress={registerAccount} style={{width: 250}}/>
                    </View>
                    
                    {/* {error && <Text style={style.error} >{error}</Text>} */}
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

export default Register