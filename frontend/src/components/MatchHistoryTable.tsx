import { useEffect } from "react";
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
    <table className="pt-6 size-full w-10/12 mx-auto">
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
              <td className="border-2 border-sky-500 w-1/6">{g.game_state === 0 ? "対戦中" : "対戦終了"}</td>
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
