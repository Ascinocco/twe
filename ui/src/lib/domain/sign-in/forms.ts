import type { LoginFormFields, LoginFormErrors } from "./types";

export const validate = ({ email, password }: LoginFormFields) => {
  let errors: LoginFormErrors = undefined;

  if (!email) {
    errors = {
      email: "Email is required",
    };
  }

  // @TODO: Better client side validation
  if (!password) {
    errors = {
      ...errors,
      password: "Password is required",
    };
  }

  return errors;
};
