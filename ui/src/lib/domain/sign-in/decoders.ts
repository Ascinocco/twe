import { string, object } from "decoders";

export const loginResponseDecoder = object({
  token: string,
  error: string,
});
