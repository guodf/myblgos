import { ArticleModel, ResultData } from "../types";
import request from "./request";

class ArticleApi {
  async getArticleById(id: number) {
    const resp = await request.get<ResultData<ArticleModel>>(`/articles/${id}`)
    if (resp.status === 200) {
      if (resp.data.code === 1000) {

        return resp.data.data
      }
    }
    return null
  }

  async getArticles() {
    const resp = await request.get<ResultData<ArticleModel[]>>("/articles")
    if (resp.status === 200) {
      if (resp.data.code === 1000) {
        return resp.data.data ?? [];
      }
    }
    return [];
  }

  async addArticle(article: any) {
    const resp = await request.post<ResultData<ArticleModel>>("/articles", article)
    if (resp.status === 200) {
      if (resp.data.code === 1000) {
        return resp.data.data;
      }
    }
    return null;
  }
}


export const articleApi = new ArticleApi();