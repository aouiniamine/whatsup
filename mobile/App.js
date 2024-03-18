import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View } from 'react-native';
import Home from './src/Screens/Home';
import { lightgreen } from './src/Styles/GlobalStyles';

export default function App() {
  return (
    <View style={styles.container}>
      <Home></Home>
      
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: lightgreen,
    
  },
});
