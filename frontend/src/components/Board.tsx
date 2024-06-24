import { useDispatch, useSelector } from "react-redux";
import { RootState, AppDispatch } from '../stores/store';
import { registerTurn, startGame } from "../stores/boardSlice";
import { useEffect } from "react";
import axios from 'axios';

const apiUrl = import.meta.env.VITE_API_KEY;

const Board = () => {
  const board = useSelector((state: RootState) => state.boardState.board);
  const turn_count = useSelector((state: RootState) => state.boardState.turn_count);
  const next_disc = useSelector((state: RootState) => state.boardState.next_disc);
  const winner_disc = useSelector((state: RootState) => state.boardState.winner_disc);
  const game_id = useSelector((state: RootState) => state.boardState.game_id);
  const dispatch = useDispatch<AppDispatch>();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.post(apiUrl + '/games', {}, {
          headers: {
            'Content-Type': 'application/json'
          },
          withCredentials: true
        });
        dispatch(startGame({ game_id: response.data.game_id }))
      } catch (error) {
        if (axios.isAxiosError(error)) {
          console.error("Axios error:", error.response?.data);
        } else {
          console.error("Unexpected error:", error);
        }
      }
    };
    fetchData();
  }, []);

  return (
    <>
      <h2 className="mt-3 text-xl">{turn_count === 0 ? "~ゲーム開始~" : `${turn_count}手目`}</h2>
      {
        winner_disc === null
          ? <h3>{next_disc === 1 ? "黒" : "白"}の順番です</h3>
          : <h3>{winner_disc === 1 ? "黒" : "白"}の勝利です</h3>
      }
      <div className='Board
        w-[30rem] m-0 m-auto mt-5 grid grid-cols-8 place-content-center place-items-center divide-solid border-0 border-t-[1px] border-l-[1px] bg-gray-300
        [&>div]:w-[60px] [&>div]:h-[60px] [&>div]:border-r-[1px] [&>div]:border-b-[1px] [&>div]:flex [&>div]:justify-center [&>div]:items-center
        '>
        {board.map((line, lineIndex) => {
          return line.map((square, squareIndex) => {
            if (square === 0) {
              return <div key={`${lineIndex}+${squareIndex}`} className="" onClick={() => dispatch(registerTurn({ game_id: game_id, turn_count: turn_count + 1, disc: next_disc, x: squareIndex, y: lineIndex }))}></div>
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
