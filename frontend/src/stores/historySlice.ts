import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";


// TODO 対戦履歴取得のAPI実行処理を記載
export const getGameResultHistory = createAsyncThunk(
  'data/getHistory',
  async () => {
    console.log('ASDAS');
    const result = await fetch('http://localhost:3000/api/games/', {
      method: 'GET'
    });

    if (!result.ok) {
      const errorData = await result.json();
      throw new Error(errorData.message || 'Something went wrong');
    };

    return result.json();
  }
);

// TODO 対戦履歴のsliceを定義
export const historySlice = createSlice({
  name: 'board',
  initialState: {
    games: [],
    loading: false,
    error: "",
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getGameResultHistory.pending, (state) => {
        state.loading = true;
        state.error = '';
        console.log('対戦履歴ロード中');
      })
      .addCase(getGameResultHistory.fulfilled, (state, { payload }) => {
        console.log(payload);
        state.loading = false;
        state.games = payload.games;
      })
      .addCase(getGameResultHistory.rejected, (state, action) => {
        state.loading = false;
        state.error = typeof (action.error.message) === "string" ? action.error.message : "";
        console.log(state.error);
        console.log('失敗');
      });
  },
});

export default historySlice.reducer;
