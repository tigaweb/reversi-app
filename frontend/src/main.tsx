import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import Game from './Game.tsx'
import './index.css'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import { store } from './stores/store';
import { Provider } from "react-redux";
import History from './History.tsx'
const router = createBrowserRouter([
  {
    path: "/app",
    element: <App />,
  },
  {
    path: "/history",
    element: <History />,
  },
  {
    path: "/game",
    element: <Game />,
  },
]);
console.log(router);

ReactDOM.createRoot(document.getElementById('root')!).render(
  // <React.StrictMode>
    <Provider store={store}>
      <RouterProvider router={router} />
    </Provider>
  // </React.StrictMode>,
)
