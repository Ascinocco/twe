import type { LoginFormFields, LoginFormErrors } from "./types";

const pwRegex = /^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$/;

export const validate = ({ email, password }: LoginFormFields) => {
  let errors: LoginFormErrors;

  if (!email) {
    errors = {
      email: "Email is required",
    };
  }

  if (!pwRegex.test(password || "")) {
    errors = {
      password: "Password must be at least 8 characters, and contain a number and symbol.",
    };
  }

  return errors;
};
