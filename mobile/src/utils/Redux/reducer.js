import { createSlice, configureStore } from '@reduxjs/toolkit'

export const messagesSlice = createSlice({
  name: 'messages',
  initialState: {
    allMessages: []
  },
  reducers: {
    addMessage: (state, message) => {
        const allMessages = state.allMessages
        for (let i in allMessages){
            let conv = allMessages[i]
            for(let j in conv){
                const convMessage = conv[j]
                if (convMessage.username === message.payload.username){
                    allMessages[i].push(message.payload)
                    state.allMessages = allMessages
                    return
                }
                if (convMessage.username !== username || convMessage.username !== message.payload.username){
                    break
                }
            }
        }
        allMessages.push([message.payload])
        
    },
  }
})

export const { addMessage, } = messagesSlice.actions
export const messagesSelector = messagesSlice.selectors