import Header from "./Header"
import { Link } from 'react-router-dom'

const Top = () => {
  return (
    <>
      <Header title="TOPページ" />
      <div>Reversi app by React/Go(Echo)</div>
      <div className="pt-5 flex items-center flex-col justify-center">
        <Link to="/game" className="text-white min-w-40 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">一人で遊ぶ</Link>
        <Link to="/history" className="text-white min-w-40 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">対戦履歴</Link>
      </div>
    </>
  );
};

export default Top;
