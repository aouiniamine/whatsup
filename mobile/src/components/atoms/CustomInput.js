import { StyleSheet, TextInput } from "react-native";
import { darkgreen, windowWidth } from "../../Styles/GlobalStyles";

export default function CustomInput ({value="", placeholder, onChange, style=null, type}) {
    return (
        <TextInput placeholder={placeholder} defaultValue={value} keyboardType={type}
        style={[style1.input, style]} onChangeText={onChange}/>
    )
}

const style1 = StyleSheet.create({
    input: {
        borderColor: darkgreen,
        padding: 15,
        borderWidth: 3,
        borderRadius: 50,
        backgroundColor: "white",
        width: windowWidth * 0.8
    }
})