import { useState, useEffect } from "react"
import { useDispatch, useSelector } from "react-redux";
import { useLocation } from 'react-router-dom';
import { RootState, AppDispatch } from "../stores/store";
import { logIn, signUp } from "../stores/authSlice";
import { LoginOrSignUp } from "../types";
import { useNavigate } from "react-router-dom";

const Auth = () => {
  const navigate = useNavigate();
  const is_Authenticated = useSelector((state: RootState) => state.authState.is_Authenticated);
  const is_login = useSelector((state: RootState) => state.authState.is_login);
  useEffect(() => {
    if (is_Authenticated !== is_login) {
      navigate('/');
    }
  }, [is_Authenticated, is_login, navigate]);
  const [userName, setUserName] = useState('')
  const [email, setEmail] = useState('')
  const [pw, setPw] = useState('')

  const location = useLocation();
  const { login }: LoginOrSignUp = location.state;

  const dispatch = useDispatch<AppDispatch>();

  return (
    <>
      <div className="pt-4 flex justify-center items-center flex-col text-gray-600 max-w-full">
        <div className="flex items-center">
          <span className="text-center text-3xl font-extrabold">
            Reversi app by React/Go(Echo)
          </span>
        </div>
        <h2 className="my-6">{login ? 'Login' : 'Create a new account'}</h2>
        <div>
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
              autoFocus={login ? true : false}
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
            {login
              ? <button
                className="disabled:opacity-40 py-2 px-4 rounded text-white bg-indigo-600"
                disabled={!email || !pw}
                // type="submit"
                onClick={() => dispatch(logIn({ email: email, password: pw }))}
              >
                Login
              </button>
              : <button
                className="disabled:opacity-40 py-2 px-4 rounded text-white bg-indigo-600"
                disabled={!userName || !email || !pw}
                // type="submit"
                onClick={() => dispatch(signUp({ user_name: userName, email: email, password: pw }))}
              >
                SignUp
              </button>
            }
          </div>
        </div>
      </div>
    </>
  )
}

export default Auth
