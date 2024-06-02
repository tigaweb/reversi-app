import { useDispatch, useSelector } from "react-redux";
import { RootState, AppDispatch } from './stores/store';
import { decrease, increase } from "./stores/counterSlice";

const App = () => {
  const count = useSelector((state: RootState) => state.counter.count);
  const dispatch = useDispatch<AppDispatch>();

  return (
    <main>
      {count}
      {/* {message} */}
      <button onClick={() => dispatch(increase())}>Up</button>
      <button onClick={() => dispatch(decrease())}>Down</button>
    </main>
  )
}

export default App
