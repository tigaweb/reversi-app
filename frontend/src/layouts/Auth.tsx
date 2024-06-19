import { useState } from "react"
import { useLocation } from 'react-router-dom';


const Auth = () => {
  const [userName, setUserName] = useState('')
  const [email, setEmail] = useState('')
  const [pw, setPw] = useState('')
  const location = useLocation();
  const { login } = location.state;
  return (
    <>
      <div className="pt-4 flex justify-center items-center flex-col text-gray-600 max-w-full">
        <div className="flex items-center">
          <span className="text-center text-3xl font-extrabold">
            Reversi app by React/Go(Echo)
          </span>
        </div>
        <h2 className="my-6">{login ? 'Login' : 'Create a new account'}</h2>
        <form>
          {login ? '' :
            <div>
              <input
                className="mb-3 px-3 text-sm py-2 border border-gray-300 w-full"
                name="userName"
                type="text"
                autoFocus
                placeholder="User Name"
                onChange={(e) => setUserName(e.target.value)}
                value={userName}
              />
            </div>
          }
          <div>
            <input
              className="mb-3 px-3 text-sm py-2 border border-gray-300 w-full"
              name="email"
              type="email"
              autoFocus
              placeholder="Email address"
              onChange={(e) => setEmail(e.target.value)}
              value={email}
            />
          </div>
          <div>
            <input
              className="mb-3 px-3 text-sm py-2 border border-gray-300 w-full"
              name="password"
              type="password"
              placeholder="Password"
              onChange={(e) => setPw(e.target.value)}
              value={pw}
            />
          </div>
          <div className="flex justify-center my-2">
            <button
              className="disabled:opacity-40 py-2 px-4 rounded text-white bg-indigo-600"
              disabled={!email || !pw}
              type="submit"
            >
              {login ? 'Login' : 'Sign Up'}
            </button>
          </div>
        </form>
      </div>
    </>
  )
}

export default Auth
