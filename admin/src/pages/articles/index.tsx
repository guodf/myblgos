import React from "react";
import ArticleListComp from "./comps/article_list";
import { useLocation } from "react-router-dom";

const ArticlesPage = (props: any) => {
  return (
    <div>
      <ArticleListComp />
    </div>
  );
};

export default ArticlesPage;
