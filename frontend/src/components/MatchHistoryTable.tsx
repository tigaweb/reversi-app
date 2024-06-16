import { useEffect } from "react";
import { getGameResultHistory } from "../stores/historySlice";
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch,RootState } from "../stores/store";

const MatchHistoryTable = () => {
  const games = useSelector((state: RootState) => state.historyState.games)
  const dispatch = useDispatch<AppDispatch>();
  useEffect(()=>{
    dispatch(getGameResultHistory());
  },[]);
  return (
    <table className="pt-5">
      <thead className="size-full">
        <tr>
          <th className="w-1/5">黒を打った回数</th>
          <th className="w-1/5">白を打った回数</th>
          <th className="w-1/5">勝った石</th>
          <th className="w-1/5">対戦開始時刻</th>
          <th className="w-1/5">対戦終了時刻</th>
        </tr>
      </thead>
      <tbody className="size-full">
        {/* <tr>
          <td className="w-1/5">1</td>
          <td className="w-1/5">1</td>
          <td className="w-1/5">1</td>
          <td className="w-1/5">2017-07-21T17:32:28z</td>
          <td className="w-1/5">2017-07-21T17:32:28z</td>
        </tr> */}
        {games.map((g)=>{
          return(
            <tr key={g.id}>
              <td className="w-1/5">{g.darkMoveCount}</td>
              <td className="w-1/5">{g.lightMoveCount}</td>
              <td className="w-1/5">{g.winnerDisc}</td>
              <td className="w-1/5">{g.startedAt}</td>
              <td className="w-1/5">{g.endAt}</td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
};

export default MatchHistoryTable;
