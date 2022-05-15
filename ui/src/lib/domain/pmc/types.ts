export type CreatePmcFormFields = {
  name: string;
};

export type CreatePmcResponse = {
  id: string;
  userId: string;
  name: string;
  error: string;
};

export type Pmc = {
  id: string;
  userId: string;
  name: string;
};
