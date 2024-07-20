export type CsrfToken = {
  csrf_token: string
}

export type SignUp = {
  user_name: string
  email: string
  password: string
}

export type Credential = {
  email: string
  password: string
}

export type HeaderProps = {
  title: string
}

export type LoginOrSignUp = {
  login: boolean | null
}

export type IsLogin = {
  is_login: boolean
}

export type boardPosition = {
  line: number,
  squre: number,
}

export type registerTurnRequestBody = {
  game_id: number,
  turn_count: number,
  move: {
    disc: number,
    x: number,
    y: number,
  }
}

export type gameId = {
  game_id: number
}

export type GameHistory = {
  game_id: number,
  game_state: number,
  winner_user_name: string | null,
  winner_disc: number | null,
  started_at: Date | string,
  end_at: Date | string,
}
