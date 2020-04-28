import request from "./request";
import { ResultData, AuthModel } from "../types";
import authState from "../state/auth";

class UseApi {
  async login(name: string, pwd: string) {
    const resp = await request.post<ResultData<AuthModel>>("/login", {
      name,
      pwd,
    });

    if (resp.status === 200) {
      console.log(resp.data);
      if (resp.data.code === 1000) {
        authState.login(resp.data.data);
        return true;
      }
    }
    return false;
  }
}

export const useApi = new UseApi();
