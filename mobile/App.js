import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View } from 'react-native';
import Home from './src/Screens/Home';
import { lightgreen } from './src/Styles/GlobalStyles';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import Connect from './src/Screens/Connect';
import Register from './src/Screens/Register';
import Validate from './src/Screens/Validate';
import MessagesProvider from './src/Context/MessagesProvider';
import UserProvider from './src/Context/UserProvider';

const Stack = createNativeStackNavigator();

export default function App() {
  return (
    <UserProvider>
    <MessagesProvider>
      <NavigationContainer>
        <Stack.Navigator screenOptions={{headerShown: false}} initialRouteName='Connect'>
          <Stack.Screen name="Home" component={Home} />
          <Stack.Screen name="Connect" component={Connect} />
          <Stack.Screen name="Register" component={Register} />
          <Stack.Screen name='Validate' component={Validate} />
        </Stack.Navigator>
      </NavigationContainer>
    </MessagesProvider>
    </UserProvider>
  );
}

