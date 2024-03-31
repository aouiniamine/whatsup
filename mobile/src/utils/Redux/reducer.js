import { createSlice, configureStore } from '@reduxjs/toolkit'
import {produce} from "immer"
export const messagesSlice = createSlice({
  name: 'messages',
  initialState: {
    allMessages: []
  },
  reducers: {
    addMessage: (state, action) => {
      return produce(state, draft =>{
        
        for (let i in draft.allMessages){
            let conv = draft.allMessages[i]
            for(let j in conv){
                const convMessage = conv[j]
                if (convMessage.username === action.payload.username){
                    draft.allMessages[i].push(action.payload)
                    return
                }
                if (convMessage.username !== username || convMessage.username !== action.payload.username){
                    break
                }
            }
        }
        draft.allMessages.push([action.payload])
        
      })
        
    },
  }
})

export const { addMessage, } = messagesSlice.actions
export const messagesSelector = messagesSlice.selectors