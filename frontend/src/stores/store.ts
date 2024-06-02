import { configureStore } from "@reduxjs/toolkit";
import counterReducer from "./counterSlice";
import boardReducer from "./boardSlice";
import historyReducer from "./historySlice";

export const store = configureStore({
  reducer: {
    counter: counterReducer,
    boardState: boardReducer,
    historyState: historyReducer,
  }
});

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
