import { useEffect } from "react";
import { Link } from 'react-router-dom'
import { getGameResultHistory } from "../stores/historySlice";
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "../stores/store";

const MatchHistoryTable = () => {
  const game_history = useSelector((state: RootState) => state.historyState.game_history)
  const dispatch = useDispatch<AppDispatch>();
  useEffect(() => {
    dispatch(getGameResultHistory());
  }, []);
  return (
    <table className="pt-6 mt-5 size-full w-10/12 mx-auto">
      <thead>
        <tr>
          <th className="border-2 border-sky-500 w-1/6">対戦ID</th>
          <th className="border-2 border-sky-500 w-1/6">状況</th>
          <th className="border-2 border-sky-500 w-1/6">勝者</th>
          <th className="border-2 border-sky-500 w-1/6">勝った石</th>
          <th className="border-2 border-sky-500 w-1/6">対戦開始時刻</th>
          <th className="border-2 border-sky-500 w-1/6">対戦終了時刻</th>
        </tr>
      </thead>
      <tbody>
        {game_history.map((g) => {
          return (
            <tr key={g.game_id}>
              <td className="border-2 border-sky-500 w-1/6">{g.game_id}</td>
              <td className="border-2 border-sky-500 w-1/6">

              {g.game_state !== 0
                ? <Link to={`/game/${g.game_id}`} className="text-white min-w-40 bg-blue-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 focus:outline-none dark:focus:ring-red-800">終了</Link>
                : <Link to={`/game/${g.game_id}`} className="text-white min-w-40 bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 focus:outline-none dark:focus:ring-green-800">再開</Link>
              }

              </td>
              <td className="border-2 border-sky-500 w-1/6">{g.winner_user_name !== "" ? g.winner_user_name : "-"}</td>
              <td className="border-2 border-sky-500 w-1/6">{g.winner_disc === 0 ? "-" : g.winner_disc === 1 ? "黒" : "白"}</td>
              <td className="border-2 border-sky-500 w-1/6 whitespace-pre-wrap">{String(g.started_at).replace(' ', '\n')}</td>
              <td className="border-2 border-sky-500 w-1/6 whitespace-pre-wrap">{g.game_state === 0 ? "-" : String(g.end_at).replace(' ', '\n')}</td>
            </tr>
          )
        })}
      </tbody>
    </table>
  );
};

export default MatchHistoryTable;
