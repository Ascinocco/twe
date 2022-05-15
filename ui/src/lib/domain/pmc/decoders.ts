import { object, string } from "decoders";

export const pmcResponseDecoder = object({
  id: string,
  userId: string,
  name: string,
  error: string,
});
