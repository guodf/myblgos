import request from "./request";
import { ResultData, CategoryModel } from "../types";

var categoriesCache: CategoryModel[] = []

class CategoryApi {
  async getCategories() {
    if (categoriesCache.length > 0) {
      return categoriesCache;
    }
    const resp = await request.get<ResultData<CategoryModel[]>>("/categories")
    if (resp.status === 200) {
      if (resp.data.code === 1000) {
        categoriesCache = resp.data.data??[];
        return categoriesCache;
      }
    }
    return [];
  }
  async addCategory(name: string) {
    const resp = await request.post<ResultData<CategoryModel>>(`/categories/${encodeURIComponent(name)}`)
    if (resp.status === 200) {
      if (resp.data.code === 1000) {
        debugger;
        if (categoriesCache.filter(category => category.id === resp.data.data.id).length === 0) {
          categoriesCache.unshift(resp.data.data)
        }
        return resp.data.data
      }
    }
    return null;
  }
}

export const categoryApi = new CategoryApi();