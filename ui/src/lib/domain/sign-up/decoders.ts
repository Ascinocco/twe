import { string, object } from "decoders";

export const signUpResponseDecoders = object({
  id: string,
  username: string,
  email: string,
  error: string,
});
