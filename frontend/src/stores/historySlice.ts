import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";
import { CsrfToken, GameHistory } from '../types'

const apiUrl = import.meta.env.VITE_API_KEY;

export const getGameResultHistory = createAsyncThunk(
  'data/getHistory',
  async () => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${apiUrl}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
    try {
      const historyResponse = await axios.get(apiUrl + '/games', {
        headers: {
          'Content-Type': 'application/json'
        },
      });
      return historyResponse.data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        const errorMessage = error.response?.data?.message || 'Something went wrong';
        throw new Error(errorMessage);
      }
    }
  }
);

type HistoryState = {
  game_history: GameHistory[],
  loading: boolean,
  error: string,
}

const initialState: HistoryState = {
  game_history: [],
  loading: false,
  error: "",
}

function formatDateString(isoString: string): string {
  const date = new Date(isoString);

  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const hours = date.getHours();
  const minutes = date.getMinutes();

  return `${year}年${month}月${day}日 ${hours}時${minutes}分`;
}

// TODO 対戦履歴のsliceを定義
export const historySlice = createSlice({
  name: 'board',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getGameResultHistory.pending, (state) => {
        state.loading = true;
        state.error = '';
        console.log('対戦履歴ロード中');
      })
      .addCase(getGameResultHistory.fulfilled, (state, { payload }) => {
        state.game_history = payload
        state.game_history.forEach((game_history) => {
          game_history.end_at = formatDateString(game_history.end_at as unknown as string);
          game_history.started_at = formatDateString(game_history.started_at as unknown as string);
        });
        state.loading = false;
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
