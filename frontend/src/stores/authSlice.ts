import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { SignUp, Credential, IsLogin } from "../types/index";
import axios from 'axios';

const apiUrl = import.meta.env.VITE_API_KEY;

export const signUp = createAsyncThunk(
  'data/signup',
  async (payload: SignUp) => {
    const { user_name, email, password } = payload;
    const requestBody: SignUp = {
      user_name,
      email,
      password
    }
    try {
      const signUpResponse = await axios.post(apiUrl + '/signup', requestBody, {
        headers: {
          'Content-Type': 'application/json'
        },
      });
      return signUpResponse.data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        const errorMessage = error.response?.data || 'Something went wrong';
        throw new Error(errorMessage);
      }
    }
  }
);

export const logIn = createAsyncThunk(
  'data/login',
  async (payload: Credential) => {
    const { email, password } = payload;
    const requestBody: Credential = {
      email,
      password
    }
    try {
      const logInResponse = await axios.post(apiUrl + '/login', requestBody, {
        headers: {
          'Content-Type': 'application/json'
        },
      });
      return logInResponse.data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        const errorMessage = error.response?.data || 'Something went wrong';
        throw new Error(errorMessage);
      }
    }
  }
);

export const logOut = createAsyncThunk(
  'data/logout',
  async () => {
    try {
      const logOutResponse = await axios.post(apiUrl + '/logout', {
        headers: {
          'Content-Type': 'application/json'
        },
      });
      return logOutResponse.data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        const errorMessage = error.response?.data?.message || 'Something went wrong';
        throw new Error(errorMessage);
      }
    }
  }
);


export const authSlice = createSlice({
  name: 'board',
  initialState: {
    is_login: false,
    is_Authenticated: false,
    user_name: '',
    error: '',
  },
  reducers: {
    setLoginState: (state, { payload }) => {
      const { is_login }: IsLogin = payload;
      state.is_login = is_login;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(signUp.pending, (state) => {
        state.error = '';
      })
      .addCase(signUp.fulfilled, (state) => {
        state.is_Authenticated = true;
      })
      .addCase(signUp.rejected, (state, action) => {
        state.error = typeof (action.error.message) === "string" ? action.error.message : "";
        console.log(state.error);
      })
      .addCase(logIn.pending, (state) => {
        state.error = '';
      })
      .addCase(logIn.fulfilled, (state) => {
        state.is_login = true;
        state.is_Authenticated = false;
      })
      .addCase(logIn.rejected, (state, action) => {
        state.error = typeof (action.error.message) === "string" ? action.error.message : "";
        console.log(state.error);
      })
      .addCase(logOut.pending, (state) => {
        state.error = '';
      })
      .addCase(logOut.fulfilled, (state) => {
        state.is_login = false;
        state.is_Authenticated = false;
      })
      .addCase(logOut.rejected, (state, action) => {
        state.error = typeof (action.error.message) === "string" ? action.error.message : "";
        console.log(state.error);
      });
  },
});

export const { setLoginState } = authSlice.actions;

export default authSlice.reducer;
