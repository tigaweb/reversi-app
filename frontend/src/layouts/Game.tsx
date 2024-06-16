import Header from '../components/Header';
import Board from "../components/Board";
import { useSelector } from 'react-redux';
import { RootState } from '../stores/store';

const Game = () => {
  const error = useSelector((state: RootState) => state.boardState.error);
  return (
    <>
      <Header title={"リバーシアプリケーション"} />
      <h2 className="text-2xl">ホーム画面</h2>
      <main>
        <div className='warning-message-area h-8'>
          {
            error !== ""
              ? <p className="warning-message h-8 w-[37rem] m-0 m-auto bg-red-500 text-white flex justify-center items-center rounded-lg animate-fadeIn">{error}</p>
              : ""
          }
        </div>
        <Board />
      </main>
    </>
  );
};

export default Game;
