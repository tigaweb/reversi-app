import { Link } from 'react-router-dom'
import { useSelector } from "react-redux"
import { RootState } from '../stores/store'

export type HeaderProps = {
  title: string
}

export type LoginOrSignUp = {
  login: boolean
}

const Header = ({ title }: HeaderProps) => {
  const is_login = useSelector((state: RootState) => state.authState.is_login);
  let login : LoginOrSignUp ;
  return (
    <header className="relative mb-8 text-2xl font-bold h-12 flex justify-center items-center font-mono">
      <h1>{title}</h1>
      <div className='absolute right-0'>
        {is_login !== true
          ? (
            <>
              <Link
                to={{
                  pathname: "/auth",
                }}
                state={{ login: false } }
                className="text-white min-w-40 bg-lime-500 hover:bg-lime-800 focus:ring-4 focus:ring-lime-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2"
              >
                SignUp
              </Link>
              <Link
                to={{
                  pathname: "/auth",
                }}
                state={{ login: true }}
                className="text-white min-w-40 bg-amber-500 hover:bg-amber-800 focus:ring-4 focus:ring-amber-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2"
              >
                LogIn
              </Link>
            </>
          )
          : <Link to="/auth" className="text-white min-w-40 bg-rose-500 hover:bg-rose-800 focus:ring-4 focus:ring-rose-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2">LogOut</Link>
        }
      </div>
    </header>
  );
};

export default Header;
