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
  is_LogIn: boolean
}

export type boardPosition = {
  line: number,
  squre: number,
}

export type registerTurnRequestBody = {
  turnCount: number,
  move: {
    disc: number,
    x: number,
    y: number,
  }
}

export type gameId = {
  game_id: number
}
