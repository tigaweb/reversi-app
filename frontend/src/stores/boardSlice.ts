import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { boardPosition, gameId, registerTurnRequestBody } from "../types/index";
import axios from "axios";
import { CsrfToken } from '../types'

const apiUrl = import.meta.env.VITE_API_KEY;

const EMPTY: number = 0;
const DARK: number = 1;
const LIGHT: number = 2;

export const registerTurn = createAsyncThunk(
  'data/placeStone',
  async (payload: { game_id: number, turn_count: number, disc: number, x: number, y: number }) => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${apiUrl}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
    const { game_id, turn_count, disc, x, y } = payload;
    const requestBody: registerTurnRequestBody = {
      game_id,
      turn_count,
      move: {
        disc,
        x,
        y,
      }
    }
    try {
      const registerTurnResult = await axios.post(apiUrl + '/games/latest/turns/', requestBody, {
        headers: {
          'Content-Type': 'application/json'
        },
        withCredentials: true,
      })
      console.log(`registerTurnResult : ${registerTurnResult.status}`)
      const findLatestTurnResult = await axios.get(apiUrl + `/games/latest/turns/${requestBody.game_id}`, {
        headers: {
          'Content-Type': 'application/json'
        },
      })
      console.log(`findLatestTurnResult : ${findLatestTurnResult.status}`)
      return findLatestTurnResult.data
    } catch (error) {
      if (axios.isAxiosError(error)) {
        const errorMessage = error.response?.data || 'Something went wrong';
        throw new Error(errorMessage);
      }
    }
  }
);

export const restartGame = createAsyncThunk(
  'data/restartGame',
  async (payload: { game_id: number }) => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${apiUrl}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
    const { game_id } = payload;
    try {
      const findLatestTurnResult = await axios.get(apiUrl + `/games/latest/turns/${game_id}`, {
        headers: {
          'Content-Type': 'application/json'
        },
      })
      console.log(`findLatestTurnResult : ${findLatestTurnResult.status}`)
      return findLatestTurnResult.data
    } catch (error) {
      if (axios.isAxiosError(error)) {
        const errorMessage = error.response?.data || 'Something went wrong';
        throw new Error(errorMessage);
      }
    }
  }
);


type BoardState = {
  game_id: number,
  board: number[][],
  turn_count: number,
  next_disc: number,
  winner_disc: number | null,
  loading: boolean,
  error: string,
}

const initialBoard: number[][] = [
  [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
  [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
  [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
  [EMPTY, EMPTY, EMPTY, LIGHT, DARK, EMPTY, EMPTY, EMPTY],
  [EMPTY, EMPTY, EMPTY, DARK, LIGHT, EMPTY, EMPTY, EMPTY],
  [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
  [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
  [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
]


const initialState: BoardState = {
  game_id: 0,
  board: initialBoard,
  turn_count: 0,
  next_disc: 1,
  winner_disc: null,
  loading: false,
  error: "",
}

export const boardSlice = createSlice({
  name: 'board',
  initialState,
  reducers: {
    set: (state, { payload }) => {
      const { line, squre }: boardPosition = payload;
      state.board[line][squre] = DARK;
    },
    startGame: (state, { payload }) => {
      const { game_id }: gameId = payload;
      state.game_id = game_id;
      state.board = initialBoard;
      state.turn_count = 0;
      state.next_disc = 1;
      state.winner_disc = null;
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(registerTurn.pending, (state) => {
        state.loading = true;
        state.error = '';
        console.log('ロード中');
      })
      .addCase(registerTurn.fulfilled, (state, { payload }) => {
        console.log(payload);
        state.loading = false;
        state.game_id = payload.game_id;
        state.turn_count = payload.turn_count;
        state.board = payload.board;
        state.next_disc = payload.next_disc;
        state.winner_disc = payload.winner_disc === 0 ? null : payload.winner_disc;
      })
      .addCase(registerTurn.rejected, (state, action) => {
        state.loading = false;
        state.error = typeof (action.error.message) === "string" ? action.error.message : "";
        console.log(state.error);
        console.log('失敗');
      })
      .addCase(restartGame.pending, (state) => {
        state.loading = true;
        state.error = '';
        console.log('ロード中');
      })
      .addCase(restartGame.fulfilled, (state, { payload }) => {
        console.log(payload);
        state.loading = false;
        state.game_id = payload.game_id;
        state.turn_count = payload.turn_count;
        state.board = payload.board;
        state.next_disc = payload.next_disc;
        state.winner_disc = payload.winner_disc === 0 ? null : payload.winner_disc;
      })
      .addCase(restartGame.rejected, (state, action) => {
        state.loading = false;
        state.error = typeof (action.error.message) === "string" ? action.error.message : "";
        console.log(state.error);
        console.log('失敗');
      });
  },
});

export const { set, startGame } = boardSlice.actions;

export default boardSlice.reducer;
