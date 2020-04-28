import { RouteComponentProps } from "react-router-dom";
import { StringDecoder } from "string_decoder";

export interface RouterProps extends RouteComponentProps {
  routers?: RouterItem[];
}
export interface RouterItem {
  path: string;
  name: string;
  icon: any;
  link?: boolean;
  exact?: boolean;
  strict?: boolean;
  component: (props: RouterProps) => JSX.Element;
}

export interface ResultData<T> {
  code: number;
  msg: string;
  data: T;
}

export interface AuthModel {
  name: string;
  token: string;
  exp: number;
}

export interface TagModel {
  id: number;
  name: string;
}

export interface CategoryModel {
  id: number;
  name: string;
}

export interface ArticleModel {
  id: number;
  title: string;
  logoUrl: string;
  overview: string;
  content: string;
  categoryId: number;
  createTime: number;
  updateTime: number;
  publishTime: number;
  tags: number[];
  status: number;
}

export interface ResModel {
  url: string | undefined;
  fileHash: string;
  name: string;
  size: string;
  ext: string;
  uploadTime: number;
}

export const defaultArticle: ArticleModel = {
  id: 0,
  title: "",
  logoUrl: "",
  overview: "",
  content: "",
  categoryId: 0,
  tags: [],
  createTime: 0,
  updateTime: 0,
  publishTime: 0,
  status: 0,
};
