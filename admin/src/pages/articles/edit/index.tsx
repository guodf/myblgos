import React, { useState, useEffect, useMemo, useCallback } from "react";
import {
  Input,
  Card,
  Upload,
  Divider,
  Tag,
  Button,
  Space,
  Switch,
  Tooltip,
  Alert,
  Modal,
  message,
  Spin,
} from "antd";
import MarkDwonEditComp from "../../../comps/markdown";
import SelectCategoryComp from "../comps/select_category";
import "./index.scss";
import { PlusOutlined, FileImageFilled } from "@ant-design/icons";
import SelectTagsComp from "../comps/select_tags";
import { tagApi } from "../../../api/tag";
import { TagModel, defaultArticle, ArticleModel } from "../../../types";
import authState from "../../../state/auth";
import { articleApi } from "../../../api/article";
import { useParams } from "react-router-dom";

const lodash = require("lodash");
var isSaving = false;
var updateQueue: any[] = [];
const ArticleEdit = () => {
  const { articleId } = useParams();
  console.log(articleId);
  const [id, setId] = useState<number>(
    articleId == null ? 0 : parseInt(articleId)
  );
  const [title, setTitle] = useState(defaultArticle.title);
  const [overview, setOverview] = useState(defaultArticle.overview);
  const [categoryId, setCategoryId] = useState(defaultArticle.categoryId);
  const [logoUrl, setLogoUrl] = useState(defaultArticle.logoUrl);
  const [tagIds, setTagIds] = useState(defaultArticle.tags ?? []);
  const [content, setContent] = useState(defaultArticle.content);

  const [allTags, setAllTags] = useState<TagModel[]>([]);
  const [visible, setVisible] = useState(false);

  const [spinning, setSpinning] = useState(true);

  const handleOvrried = (e: any) => {
    setOverview(e.target.value.trim().substr(0, 100));
  };

  useEffect(() => {
    tagApi.getTags().then((data) => {
      if (data.length > 0) {
        setAllTags(data);
      }
    });
  }, []);
  useEffect(() => {
    articleApi.getArticleById(id).then((data: any) => {
      if (data != null) {
        setTitle(data.title);
        setOverview(data.overview);
        setCategoryId(data.categoryId);
        setLogoUrl(data.logoUrl);
        setTagIds(data.tags ?? []);
        setContent(data.content);
        setSpinning(false);
      }
    });
  }, [id]);

  const saveArticle = (isPublish = false) => {
    updateQueue.push({
      id,
      title,
      logoUrl,
      overview,
      content,
      categoryId,
      tags: tagIds,
      status: isPublish ? 1 : 0,
    });
    if (isSaving) {
      return;
    }
    isSaving = true;
    setInterval(() => {
      if (updateQueue.length === 0) {
        return;
      }
      console.log(updateQueue);
      const article: any = updateQueue.pop();
      console.log(article);
      updateQueue = [];
      articleApi
        .addArticle(article)
        .then((data) => {
          setSpinning(false);
          if (data != null) {
            setId(data.id);
            console.log(setId);
            message.success("保存成功");
          }
        })
        .catch(() => {});
    }, 5000);
  };
  const showTags = useMemo(() => {
    return (
      <Card style={{ textAlign: "center" }}>
        {allTags
          .filter((tag) => tagIds.indexOf(tag.id) > -1)
          .map((tag) => {
            return (
              <Tag
                key={tag.id}
                closable
                onClose={() => {
                  setTagIds(lodash.pull(tagIds, tag.id));
                }}
              >
                {tag.name}
              </Tag>
            );
          })}
      </Card>
    );
  }, [allTags, tagIds]);
  const showSelectCategoyComp = useMemo(() => {
    return (
      <SelectCategoryComp
        defaultValue={categoryId}
        onSelectCategory={(id) => setCategoryId(id)}
      />
    );
  }, [categoryId]);

  const showLogoUrl = useMemo(() => {
    return (
      <Card bordered={true}>
        <Upload
          name="file"
          listType="picture"
          showUploadList={false}
          headers={{
            Authorization: authState.authModel?.token ?? "",
          }}
          action={`${process.env.REACT_APP_Base_Url}/upload`}
          onChange={(info) => {
            const { file } = info;
            if (file.status === "done") {
              if (file.response.code === 1000) {
                setLogoUrl(
                  `${process.env.REACT_APP_Base_Url}${file.response.data}`
                );
              }
            }
          }}
        >
          {logoUrl.length === 0 && (
            <label>
              <FileImageFilled />
              上传文章Logo
            </label>
          )}
          {logoUrl.length > 0 && (
            <img src={logoUrl} alt="avatar" style={{ width: "100%" }} />
          )}
        </Upload>
      </Card>
    );
  }, [logoUrl]);
  return (
    <Spin
      style={{ width: "100%", height: "100%" }}
      size="large"
      tip="加载中..."
      spinning={spinning}
    >
      <div className="wrap">
        <div className="left">
          <div className="head">
            <Input.Group size="large">
              <Input
                placeholder="文章标题"
                style={{ width: "70%", paddingRight: "8px" }}
                addonAfter={`${30 - title.length}`}
                value={title}
                onChange={(e) => {
                  setTitle(e.target.value.trim().substr(0, 30));
                }}
              />
              {showSelectCategoyComp}
            </Input.Group>
          </div>
          <div className="body">
            <MarkDwonEditComp
              content={content}
              setContent={(value) => {
                setContent(value);
                saveArticle();
              }}
            ></MarkDwonEditComp>
          </div>
        </div>
        <div className="right">
          <Space size="large">
            <Tooltip title="文章内容有变动时,5S自动保存">
              <Switch
                checkedChildren="开"
                unCheckedChildren="关"
                defaultChecked
              />
            </Tooltip>
            <Button
              size="middle"
              type="primary"
              onClick={() => {
                setSpinning(true);
                saveArticle(false);
              }}
            >
              仅保存
            </Button>
            <Button
              size="middle"
              type="primary"
              onClick={() => {
                setSpinning(true);
                saveArticle(true);
              }}
            >
              保存并发布
            </Button>
          </Space>
          <Divider className="spliter" dashed={true} orientation="center">
            设置logo
          </Divider>
          {showLogoUrl}
          <Divider className="spliter" dashed={true} orientation="center">
            文章概述
          </Divider>
          <Card bordered={true} bodyStyle={{ padding: "0px" }}>
            <Alert
              message={`你还可以输入${100 - overview.length}个字符`}
              type="info"
            />
            <Input.TextArea
              placeholder="用100个字符描述你的文章"
              autoSize={{ minRows: 6, maxRows: 6 }}
              style={{ resize: "none", padding: "5px" }}
              value={overview}
              onChange={handleOvrried}
            ></Input.TextArea>
          </Card>
          <Divider className="spliter" dashed={true} orientation="center">
            <Button
              type="link"
              icon={<PlusOutlined />}
              onClick={() => {
                setVisible(!visible);
              }}
            >
              添加标签
            </Button>
          </Divider>
          {showTags}
          <Modal
            title="选择文章标签"
            visible={visible}
            destroyOnClose={true}
            footer={null}
            onCancel={() => {
              setVisible(false);
            }}
          >
            <SelectTagsComp
              selectTags={tagIds}
              setSelectTags={(ids) => {
                console.log(ids);
                setTagIds(ids);
              }}
            />
          </Modal>
        </div>
      </div>
    </Spin>
  );
};

export default ArticleEdit;
