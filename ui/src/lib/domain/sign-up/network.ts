import { getReqUrl, NetworkError } from "$lib/utils/network";
import type { SignUpFormFields, SignUpResponse } from "./types";
import { signUpResponseDecoders } from "./decoders";

const signInEndpoint = () => `${getReqUrl()}/user/create`;

export const signUp = (fields: SignUpFormFields): Promise<SignUpResponse> =>
  fetch(signInEndpoint(), {
    method: "POST",
    body: JSON.stringify({
      ...fields,
    }),
  })
    .then((res) => res.json())
    .then((res) => {
      const sr = signUpResponseDecoders.verify(res);

      if (!sr) {
        throw new NetworkError();
      }

      return sr;
    });
