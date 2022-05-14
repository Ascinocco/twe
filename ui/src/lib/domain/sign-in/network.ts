import { getReqUrl } from "$lib/utils/network";
import type { LoginFormFields, LoginResponse } from "./types";
import { loginResponseDecoder } from "./decoders";

const loginEndpoint = () => `${getReqUrl()}/session/create`;

export const login = (fields: LoginFormFields): Promise<LoginResponse> =>
  fetch(loginEndpoint(), {
    method: "POST",
    body: JSON.stringify({
      ...fields,
    }),
  })
    .then((res) => res.json())
    .then((res) => {
      const lr = loginResponseDecoder.verify(res);
      if (!lr) {
        throw new Error("Network error occured");
      }

      return lr;
    });
