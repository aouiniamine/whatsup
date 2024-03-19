import { StyleSheet, View, Text } from "react-native"
import { darkgreen, lightgreen, windowWidth } from "../../Styles/GlobalStyles"
import { FontAwesome } from '@expo/vector-icons';
import { Entypo } from '@expo/vector-icons';
const NavBar = () =>{
    return (
        <View style={style.wrapper}>
            <View style={style.mainContainer}> 

                <View style={style.container}>
                    <Text style={style.title}>WhatsUp</Text>

                    <View style={style.icons}>
                        <FontAwesome name="search" size={22} color="white" />
                        <Entypo name="dots-three-vertical" size={22} color="white" />
                    </View>

                </View>
                <View style={[style.container, {marginTop: 20, gap: 30}]}>
                    <Text style={[style.navs, style.openedNav]}>Chats</Text>
                    <Text style={style.navs}>Calls</Text>
                    <Text style={style.navs}>Updates</Text>
                </View>
            </View>
        </View>
    )
}

const style = StyleSheet.create({
    wrapper: {
        backgroundColor: darkgreen,
        width: windowWidth,
        height: "15%",
        marginTop: 0,
    },
    mainContainer: {
        marginTop: 30,
        flex: 1,
        flexDirection: "column"
    },
    container: {
        flex: 1,
        flexDirection: 'row',
        // marginTop: 40,

    },
    title: {
        flex: 1,
        fontWeight: "400",
        fontSize: 25,
        color: "white",
        marginLeft: 10
    },
    icons: {
        
        flex: 1,
        flexDirection: "row",
        marginTop: 10,
        justifyContent: "flex-end",
        gap: 30,
        marginRight: 25
    },
    navs: {
        textAlignVertical: "center",
        textAlign: "center",
        flex: 1,
        color: lightgreen,
        fontSize: 18,
        fontWeight: "600",

    },
    openedNav: {
        color: "white",
        borderBottomColor: "white",
        borderBottomWidth: 2
    }
})

export default NavBar

