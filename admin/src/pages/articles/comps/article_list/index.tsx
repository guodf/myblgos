import React, { useState, useEffect } from "react";
import { Table, Tag, Space } from "antd";
import Column from "antd/lib/table/Column";
import { Link } from "react-router-dom";
import { formatTime } from "../../../../utils";
import { ArticleModel } from "../../../../types";
import { articleApi } from "../../../../api/article";
import { CheckCircleOutlined, EditOutlined } from "@ant-design/icons";

const ArticleListComp = (props: any) => {
  const [articles, setArticles] = useState<ArticleModel[]>([]);

  useEffect(() => {
    articleApi.getArticles().then((data) => {
      setArticles(data);
    });
  }, []);

  const getArticleStatus = (article: ArticleModel) => {
    if (article.status === 1) {
      return (
        <Tag icon={<CheckCircleOutlined />} color="success">
          已发布
        </Tag>
      );
    }

    return (
      <Tag icon={<EditOutlined />} color="warning">
        草稿
      </Tag>
    );
  };

  return (
    <div>
      <Table dataSource={articles} pagination={false}>
        <Column title="文章ID" dataIndex="id" key="id" />
        <Column title="标题" dataIndex="title" key="title" />
        <Column title="文章状态" key="status" render={getArticleStatus} />
        <Column
          title="创建时间"
          key="createTime"
          render={(value) => formatTime(value.createTime)}
        />
        <Column
          title="最后修改时间"
          key="updateTime"
          render={(value) => formatTime(value.updateTime)}
        />
        <Column
          title="发布时间"
          key="publishTime"
          render={(value) => formatTime(value.publishTime)}
        />

        <Column
          key="action"
          render={(text, record: any) => (
            <Space>
              <Link
                to={{
                  pathname: `/article/edit/${record.id}`,
                }}
                className="ant-dropdown-link"
              >
                <EditOutlined />
                编辑
              </Link>
              <Link
                to={{
                  pathname: `/article/edit/${record.id}`,
                  state: record,
                }}
                className="ant-dropdown-link"
              >
                发布
              </Link>
            </Space>
          )}
        />
      </Table>
    </div>
  );
};

export default ArticleListComp;
