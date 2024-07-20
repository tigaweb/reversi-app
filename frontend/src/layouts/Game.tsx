import Header from '../components/Header';
import Board from "../components/Board";
import { useParams } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { RootState } from '../stores/store';
import { Link } from 'react-router-dom'

const Game = () => {
  const { game_id } = useParams();
  const error = useSelector((state: RootState) => state.boardState.error);
  return (
    <>
      <Header title={"リバーシアプリケーション"} />
      <main>
        <div className='warning-message-area'>
          {
            error !== ""
              ? <p className="warning-message h-8 w-[37rem] m-0 m-auto bg-red-500 text-white flex justify-center items-center rounded-lg animate-fadeIn">{error}</p>
              : ""
          }
        </div>
        <Board gameId={game_id} />
        <Link className="text-white min-w-40 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5  dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800" to={`/`}>HOMEへ戻る</Link>
      </main>
    </>
  );
};

export default Game;
