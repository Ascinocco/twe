import type { SignUpFormFields, SignUpFormErrors } from "./types";

const pwRegex = /^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$/;

export const validate = ({
  email,
  username,
  password,
  passwordConfirmation,
  pmcName,
}: SignUpFormFields) => {
  let errors: SignUpFormErrors;

  if (!email) {
    errors = {
      email: "Email is required",
    };
  }

  if (!username || username.length < 3) {
    errors = {
      username: "Username must be at least 3 characters",
    };
  }

  if (!pwRegex.test(password || "")) {
    errors = {
      password: "Password must be at least 8 characters, and contain a number and symbol.",
    };
  }

  if (passwordConfirmation !== password) {
    errors = {
      ...errors,
      passwordConfirmation: "Password do not match",
    };
  }

  if (!pmcName || pmcName.length < 5) {
    errors = {
      ...errors,
      pmcName: "PMC name must be at least 5 characters",
    };
  }

  return errors;
};
