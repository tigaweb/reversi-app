import { useState } from "react"

const Auth = () => {
  const [email, setEmail] = useState('')
  const [pw, setPw] = useState('')
  const [isLogin, setIsLogin] = useState(true)
  // const {logIn}
  return (
    <>
      <div className="pt-4 flex justify-center items-center flex-col text-gray-600 max-w-full">
        <div className="flex items-center">
          <span className="text-center text-3xl font-extrabold">
            Todo app by React/Go(Echo)
          </span>
        </div>
        <h2 className="my-6">{isLogin ? 'Login' : 'Create a new account'}</h2>
        <form>
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
              {isLogin ? 'Login' : 'Sign Up'}
            </button>
          </div>
        </form>
      </div>
    </>
  )
}

export default Auth
