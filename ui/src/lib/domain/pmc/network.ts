import { getReqUrl, NetworkError, getAuthHeaders } from "$lib/utils/network";
import type { CreatePmcFormFields, CreatePmcResponse } from "./types";
import { pmcResponseDecoder } from "./decoders";

const pmcEndpoint = () => `${getReqUrl()}/pmc/create`;

export const createPmc = (fields: CreatePmcFormFields): Promise<CreatePmcResponse> =>
  fetch(pmcEndpoint(), {
    method: "POST",
    headers: getAuthHeaders(),
    body: JSON.stringify({
      ...fields,
    }),
  })
    .then((res) => res.json())
    .then((res) => {
      const pr = pmcResponseDecoder.verify(res);

      if (!pr) {
        throw new NetworkError();
      }

      return pr;
    });
