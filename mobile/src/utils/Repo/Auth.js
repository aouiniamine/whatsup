import axios from 'axios'
import { SERVER } from './envirement'
import { setToken } from './secureStorage'
export const connect = (credential) =>
    axios.post(`${SERVER}/api/auth/connect`, {credential})
    .then(res => {
        if (res.data?.email){
            return res.data
        } else throw res.data
    })

export const register = (credential) =>
    axios.post(`${SERVER}/api/auth/register`, credential)
    .then(res => {
        if (res.data?.email){
            return res.data
        } else throw res.data
    })

export const validate = (body) =>
    axios.post(`${SERVER}/api/auth/validate`, body)
    .then(res => setToken(res.data))