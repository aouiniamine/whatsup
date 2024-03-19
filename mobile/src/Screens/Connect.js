import { KeyboardAvoidingView, Platform, StyleSheet, View, Text } from "react-native"
import CustomInput from "../components/atoms/CustomInput"
import { darkgreen, lightgreen, windowWidth } from "../Styles/GlobalStyles"

const Connect = () =>{
    return (
        <KeyboardAvoidingView behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
            style={style.wrapper}>
                <View style={style.container}>
                    <Text style={style.title}>Connect Account</Text>
                    <CustomInput placeholder={"E-mail or username"}/>

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
        gap: 20
    },
    title: {
        fontSize: 25,
        fontWeight: '500',
        color: darkgreen,
    }
})

export default Connect