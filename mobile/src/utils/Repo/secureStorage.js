import * as SecureStore from 'expo-secure-store';

const secretTokenKey = "____________________TOKEN_KEY____________________"
export async function setToken (token){
    await SecureStore.setItemAsync(secretTokenKey, token)
}

export async function getToken(){
    return await SecureStore.getItemAsync(secretTokenKey)
}
