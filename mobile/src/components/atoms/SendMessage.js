import { useContext, useEffect, useRef, useState } from "react";
import { Modal, Pressable, StyleSheet, Text, TextInput, View } from "react-native";
import { darkgreen, warningRed } from "../../Styles/GlobalStyles";
import CustomInput from "./CustomInput";
import { MessagesContext } from "../../Context/MessagesProvider";
import { validateUsername } from "../../utils/general";


export default function SendMessage ({isVisible, closeForm}){
    const [to, setTo ] = useState("")
    const [message, setMesssage] = useState("")
    const {sendMessage} = useContext(MessagesContext)
    const [invalidUsernameErr, setInvalidUsernameErr] = useState(false) 
    const handleSend = ()=>{
        if(validateUsername(to)){
            sendMessage(to, message)
            closeForm()

        }else{
            setInvalidUsernameErr(true)
        }
    }
    useEffect(()=>{
        if (invalidUsernameErr){
            const timeoutId = setTimeout(() => {
                setInvalidUsernameErr(false)
              }, 1200);
          
              return () => clearTimeout(timeoutId);
        }
    }, [invalidUsernameErr]) 
    return (

            <Modal
                animationType="slide"
                transparent={true}
                visible={isVisible}
                onRequestClose={closeForm}
                >
                
                <View style={styles.centeredView}>
                <View style={styles.modalView}>
                    <CustomInput onChange={setTo} placeholder={"Sending To:"}
                    style={[{width: "100%", height: 50,}, invalidUsernameErr ? {borderColor: warningRed}: null]}></CustomInput>
                    <CustomInput onChange={setMesssage} placeholder={"Message:"} style={{width: "100%", height: 50}}></CustomInput>
                    
                    <Pressable
                    style={[styles.button, styles.send]}
                    onPress={handleSend}
                    >
                    <Text style={styles.textStyle}>Send</Text>
                    </Pressable>
                </View>
                </View>
            </Modal>
    )
    
}

const styles = StyleSheet.create({ 
    centeredView: {
      flex: 1,
      justifyContent: 'center',
      alignItems: 'center',
      marginTop: 22,
    },
    modalView: {
        gap: 20,
      margin: 20,
      backgroundColor: 'white',
      borderRadius: 20,
      padding: 35,
      alignItems: 'center',
      shadowColor: '#000',
      shadowOffset: {
        width: 0,
        height: 2,
      },
      shadowOpacity: 0.25,
      shadowRadius: 4,
      elevation: 5,
      minWidth: 300,
      maxWidth: 300,
      maxHeight: 250,
      minHeight: 250
    },
    button: {
      borderRadius: 20,
      padding: 10,
      elevation: 2,
      width: 80,
      alignSelf: 'flex-end'
    },
    buttonOpen: {
      backgroundColor: '#F194FF',
    },
    send: {
      backgroundColor: darkgreen,
    },
    textStyle: {
      color: 'white',
      fontWeight: 'bold',
      textAlign: 'center',
    },
    modalText: {
      marginBottom: 15,
      textAlign: 'center',
    },
  });