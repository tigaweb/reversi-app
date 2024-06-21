import { useDispatch, useSelector } from "react-redux";
import { RootState, AppDispatch } from '../stores/store';
import { registerTurn } from "../stores/boardSlice";
import { useEffect } from "react";

const Board = () => {
  const board = useSelector((state: RootState) => state.boardState.board);
  const turnCount = useSelector((state: RootState) => state.boardState.turn_count);
  const nextDisc = useSelector((state: RootState) => state.boardState.next_disc);
  const winnerDisc = useSelector((state: RootState) => state.boardState.winner_disc);
  const dispatch = useDispatch<AppDispatch>();

  useEffect(() => {
    const registerGame = async () => {
      await fetch('http://localhost:3000/api/games', {
        method: 'POST'
      })
    };
    registerGame();
  }, []);

  return (
    <>
      <h2 className="mt-3 text-xl">{turnCount === 0 ? "~ゲーム開始~" : `${turnCount}手目`}</h2>
      {
        winnerDisc === null
          ? <h3>{nextDisc === 1 ? "黒" : "白"}の順番です</h3>
          : <h3>{winnerDisc === 1 ? "黒" : "白"}の勝利です</h3>
      }
      <div className='Board
        w-[30rem] m-0 m-auto mt-5 grid grid-cols-8 place-content-center place-items-center divide-solid border-0 border-t-[1px] border-l-[1px] bg-gray-300
        [&>div]:w-[60px] [&>div]:h-[60px] [&>div]:border-r-[1px] [&>div]:border-b-[1px] [&>div]:flex [&>div]:justify-center [&>div]:items-center
        '>
        {board.map((line, lineIndex) => {
          return line.map((square, squareIndex) => {
            if (square === 0) {
              return <div key={`${lineIndex}+${squareIndex}`} className="" onClick={() => dispatch(registerTurn({ turnCount: turnCount + 1, disc: nextDisc, x: squareIndex, y: lineIndex }))}></div>
            }
            return (
              square === 1
                ? <div key={`${lineIndex}+${squareIndex}`} className="stone dark"></div>
                : <div key={`${lineIndex}+${squareIndex}`} className="stone light"></div>
            );
          })
        })}
      </div >
    </>
  );
}
export default Board;
