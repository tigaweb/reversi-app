import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { SignUp, Credential } from "../types/index";

const apiUrl = import.meta.env.REACT_APP_API_URL;

const IS_LOGIN: boolean = true;
const NO_LOGIN: boolean = false;

export const signUp = createAsyncThunk(
  'data/signup',
  async (payload: SignUp) => {
    const { user_name, email, password } = payload;
    const requestBody: SignUp = {
      user_name,
      email,
      password
    }
    const result = await fetch(apiUrl + '/signup', {
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

    //TODO ユーザー作成成功したら続けてログインする
    
    return result.json();
  }
);


export const authSlice = createSlice({
  name: 'board',
  initialState: { // NOTE フロントで管理する状態は、ログイン状態とユーザー名
    is_login: NO_LOGIN,
    user_name: '',
    error: '',
  },
  reducers: {
  },
  extraReducers: (builder) => {
    builder
      .addCase(signUp.pending, (state) => {
        state.error = '';
        console.log('ユーザー作成中');
      })
      .addCase(signUp.fulfilled, (state, { payload }) => {
        console.log(payload);
        state.is_login = IS_LOGIN;
        state.user_name = payload.user_name;
      })
      .addCase(signUp.rejected, (state, action) => {
        state.error = typeof (action.error.message) === "string" ? action.error.message : "";
        console.log(state.error);
        console.log('失敗');
      });
  },
});

// export const { set } = boardSlice.actions;

export default authSlice.reducer;
