import request from "./request";
import { ResultData, ResModel } from "../types";

class ResApi {
  async getRes() {
    const resp = await request.get<ResultData<ResModel[]>>("/res")
    if (resp.status === 200) {
      if (resp.data.code === 1000) {
        return resp.data.data ?? []
      }
    }
    return []
  }
}


export const resApi = new ResApi();