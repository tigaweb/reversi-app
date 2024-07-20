// import { useDispatch, useSelector } from "react-redux";
import { Link } from "react-router-dom";
import Header from "../components/Header";
import MatchHistoryTable from '../components/MatchHistoryTable';

const History = () => {
  return (
    <main>
      <Header title={"対戦履歴"} />
      <Link className="text-white min-w-40 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5  dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800" to={`/`}>HOMEへ戻る</Link>
      <MatchHistoryTable />
    </main>
  )
}

export default History
