import { useSelector } from "react-redux";
import Header from "../components/Header"
import { Link } from 'react-router-dom'
import { RootState } from "../stores/store";

const Top = () => {
  // TODO ページリロードするとis_LogInの状態が初期化されるので、CookieにJWTトークンが保持されている場合に認証状態と判定するようにしたい
  const is_LogIn = useSelector((state: RootState) => state.authState.is_LogIn);
  return (
    <>
      <Header title="TOPページ" />
      <div>Reversi app by React/Go(Echo)</div>
      {is_LogIn
        ? (
          <div className="pt-5 flex items-center flex-col justify-center">
        <Link
          to="/game"
          className="text-white min-w-40 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
        >
          一人で遊ぶ
        </Link>
        <Link
          to="/history"
          className="text-white min-w-40 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
        >
          対戦履歴
        </Link>
      </div>
        )
        :(
          <div>ログインしてください</div>
        )
      }

    </>
  );
};

export default Top;
