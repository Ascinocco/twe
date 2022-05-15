export type SignUpFormErrors =
  | {
      email?: string;
      username?: string;
      password?: string;
      passwordConfirmation?: string;
      pmcName?: string;
    }
  | undefined;

export type SignUpResponse = {
  id: string;
  username: string;
  email: string;
  error: string;
};

export type SignUpFormFields = {
  email?: string;
  username?: string;
  password?: string;
  passwordConfirmation?: string;
  pmcName?: string;
};
