// @TOOD: Allow value to be set by env vars
export const getReqUrl = () => "http://localhost:8080";

export class NetworkError extends Error {
  constructor() {
    super("A network error occured, please try again.");
    this.name = "NetworkError";
  }
}

export const getAuthHeaders = () => ({
  Authorization: `Bearer ${sessionStorage.getItem("token")}`,
});
