import * as SecureStore from 'expo-secure-store';

const secretTokenKey = "____________________TOKEN_KEY____________________"
export async function addToken (token){
    await SecureStore.setItemAsync(secretTokenKey, token)
}

export async function getToken(){
    return await SecureStore.getItemAsync(secretTokenKey)
}

export async function removeToken(){
    await SecureStore.deleteItemAsync(secretTokenKey)
    console.log('user is logged out')
}