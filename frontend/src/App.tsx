import { useEffect } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import axios from 'axios'
import { CsrfToken } from './types'
import Game from './layouts/Game.tsx'
import History from './layouts/History.tsx'
import Top from './layouts/Top.tsx'
import Auth from './layouts/Auth.tsx'

const apiUrl = import.meta.env.VITE_API_KEY;

function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${apiUrl}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
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
