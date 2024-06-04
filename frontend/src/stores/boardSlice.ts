import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";

const apiUrl = import.meta.env.REACT_APP_API_URL;

const EMPTY: number = 0;
const DARK: number = 1;
const LIGHT: number = 2;

type boardPosition = {
  line: number,
  squre: number,
}

type registerTurnRequestBody = {
  turnCount: number,
  move: {
    disc: number,
    x: number,
    y: number,
  }
}

export const registerTurn = createAsyncThunk(
  'data/placeStone',
  async (payload: { turnCount: number, disc: number, x: number, y: number }) => {
    const { turnCount, disc, x, y } = payload;
    const requestBody: registerTurnRequestBody = {
      turnCount,
      move: {
        disc,
        x,
        y,
      }
    }
    const result = await fetch(apiUrl + 'games/latest/turns', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestBody)
    });

    if (!result.ok) {
      const errorData = await result.json();
      throw new Error(errorData.message || 'Something went wrong');
    };

    const response = await fetch(apiUrl + `games/latest/turns/${turnCount}`, {
      method: 'GET'
    });
    return response.json();
  }
);


export const boardSlice = createSlice({
  name: 'board',
  initialState: { // NOTE ターン数(turn_count)、次のコマの色(next_disc)も管理すべき?
    board: [
      [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
      [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
      [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
      [EMPTY, EMPTY, EMPTY, LIGHT, DARK, EMPTY, EMPTY, EMPTY],
      [EMPTY, EMPTY, EMPTY, DARK, LIGHT, EMPTY, EMPTY, EMPTY],
      [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
      [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
      [EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY],
    ],
    turn_count: 0,
    next_disc: 1,
    winner_disc: null,
    loading: false,
    error: "",
  },
  reducers: { // NOTE 自分の色は何か、boardの位置をどのように渡すか
    set: (state, { payload }) => {
      const { line, squre }: boardPosition = payload;
      state.board[line][squre] = DARK;
    },
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
        state.turn_count = payload.turnCount;
        state.board = payload.board;
        state.next_disc = payload.nextDisc;
        state.winner_disc = payload.winnerDisc;
      })
      .addCase(registerTurn.rejected, (state, action) => {
        state.loading = false;
        state.error = typeof (action.error.message) === "string" ? action.error.message : "";
        console.log(state.error);
        console.log('失敗');
      });
  },
});

// export const { set } = boardSlice.actions;

export default boardSlice.reducer;
