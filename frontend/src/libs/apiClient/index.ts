import { useMemo } from "react";
import { apiClient } from "../../generated/apiclient";
import { getBaseUrl } from "../baseUrl";

export const useApiClient = () => {
  const baseUrl = getBaseUrl();
  const client = useMemo(() => apiClient(baseUrl), [baseUrl]);
  return client;
};
