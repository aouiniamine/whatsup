import { KeyboardAvoidingView, Platform, StyleSheet, View, Text } from "react-native"
import CustomInput from "../components/atoms/CustomInput"
import { darkgreen, lightgreen, lightgrey, windowWidth } from "../Styles/GlobalStyles"
import { useState } from "react"
import CustomButton from "../components/atoms/CustomButton"
import { validate } from "../utils/Repo/Auth"

const Validate = ({route, navigation}) =>{
    const {email} = route.params;
    const [code, setCode ] = useState("")
    const [error, setError] = useState('')
    const login = ()=> navigation.navigate("Register")
    const onValidate = async() =>{
        try{
            const body ={
                email, code: Number(code)
            }
            console.log(body)
            await validate(body).then(navigation.navigate("Home"))
            
        }catch(err) { console.log('error: ', err.message); setError(err.message)}
        
    }
    return (
        <KeyboardAvoidingView behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
            style={style.wrapper}>
                <View style={style.mainContainer}>

                    <View style={style.container}>
                        <View>
                            <Text style={style.title}>Validation</Text>
                            <Text style={style.for}>For: {email}</Text>
                            
                        </View>
                        <CustomInput placeholder={"Format: XXXXXX"} onChange={setCode} type={'numeric'}/>
                        <CustomButton text={"Validate"} onPress={onValidate} style={{width: 250}}/>
                    </View>
                    <Text onPress={login} style={style.login}>create an account ?</Text>
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
    for: {
        fontSize: 16,
        fontWeight: '500',
        color: darkgreen,
    },
    error: {
        
        textAlign: "center",
        color: 'white',
        fontSize: 19
    },
    login: {
        textAlign: "center",
        color: "grey",
        fontSize: 19
    }
})

export default Validate