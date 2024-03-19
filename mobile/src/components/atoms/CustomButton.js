import { Pressable, StyleSheet, Text } from "react-native";
import { darkgreen, windowWidth } from "../../Styles/GlobalStyles";

export default function CustomButton({text, onPress, style, textStyle}) {
    return (
        <Pressable onPress={onPress} style={[style1.button, style]}>
            <Text style={[style1.text, textStyle]}>{text}</Text>
        </Pressable>
    )
}

const style1 = StyleSheet.create({
    button: {
        backgroundColor: darkgreen,
        borderRadius: 50,
        padding: 20,
        width: windowWidth * 0.5,
        alignSelf: "center"
    },
    text: {
        textAlign: "center",
        fontSize: 19,
        color: "white",
    }
})