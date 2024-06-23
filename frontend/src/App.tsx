import { useEffect } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import axios from 'axios'
import { CsrfToken } from './types'
import Game from './layouts/Game.tsx'
import History from './layouts/History.tsx'
import Top from './layouts/Top.tsx'
import Auth from './layouts/Auth.tsx'
import { AppDispatch } from "./stores/store";
import { useDispatch } from "react-redux"

const apiUrl = import.meta.env.VITE_API_KEY;

import { setLoginState } from './stores/authSlice.ts';

function App() {
  const dispatch = useDispatch<AppDispatch>();
  useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${apiUrl}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
    const checkAuth = async () => {
      try {
        axios.defaults.withCredentials = true
        const response = await axios.get(`${apiUrl}/check-auth`, {
          withCredentials: true,
          validateStatus: (status) => {
            return status === 200 || status === 201;
          }
        });
        if (response.status === 200 && response.data.loggedIn) {
          dispatch(setLoginState({ is_LogIn: response.data.loggedIn }))
        } else if (response.status === 201) {
          console.log(response.data.message);
        }
      } catch (error) {
        console.error('認証チェックに失敗しました:', error);
      }
    }
    checkAuth()
  }, [])
  return (
    <BrowserRouter >
      <Routes>
        <Route path="/" element={<Top />} />
        <Route path="/auth" element={<Auth />} />
        <Route path="/game" element={<Game />} />
        <Route path="/history" element={<History />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
