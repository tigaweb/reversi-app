// import { useDispatch, useSelector } from "react-redux";
import { Link } from "react-router-dom";
import Header from "../components/Header";
import MatchHistoryTable from '../components/MatchHistoryTable';

const History = () => {
  return (
    <main>
      <Header title={"対戦履歴"} />
      <Link to="/game" className="pt-3 text-blue-600 visited:text-purple-600">対戦開始</Link>
      <MatchHistoryTable />
    </main>
  )
}

export default History
