import request from "./request";
import { TagModel, ResultData } from "../types";

var tagsCache: TagModel[] = []

class TagApi {
  async getTags() {
    if (tagsCache?.length > 0) {
      return tagsCache;
    }
    const resp = await request.get<ResultData<TagModel[]>>("/tags")
    if (resp.status === 200) {
      if (resp.data.code === 1000) {
        tagsCache = resp.data.data ?? [];
        return tagsCache;
      }
    }
    return []
  }
  async addTag(name: string) {
    const resp = await request.post<ResultData<TagModel>>(`/tags/${name}`)
    if (resp.status === 200) {
      if (resp.data.code === 1000) {
        if (tagsCache.filter(tag => tag.id === resp.data.data.id).length === 0) {
          tagsCache.unshift(resp.data.data)
        }
        return resp.data.data;
      }
    }
    return null;
  }
}

export const tagApi = new TagApi();