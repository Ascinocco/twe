export type LoginFormErrors =
  | {
      email?: string;
      password?: string;
    }
  | undefined;

export type LoginFormFields = {
  email?: string;
  password?: string;
};

export type LoginResponse = {
  token: string;
  error: string;
};
