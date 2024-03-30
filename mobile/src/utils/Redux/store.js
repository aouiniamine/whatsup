import { createSlice, configureStore } from '@reduxjs/toolkit'
import { messagesSlice } from "./reducer"

export default configureStore({
  reducer: messagesSlice.reducer
})
