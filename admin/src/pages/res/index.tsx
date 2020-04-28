import React, { useEffect, useState } from "react";
import { resApi } from "../../api/res";
import { ResModel } from "../../types";
import { Card, Upload, Button, message } from "antd";
import copy from "copy-to-clipboard";

import "./index.scss";
import authState from "../../state/auth";
import { PictureOutlined, PlusOutlined } from "@ant-design/icons";

const ResPage = () => {
  const [res, setRes] = useState<ResModel[]>([]);

  useEffect(() => {
    resApi.getRes().then((data) => {
      setRes(data);
    });
  }, []);

  const getUrl = (res: ResModel) => {
    return `${process.env.REACT_APP_Base_Url}/res/${res.fileHash}${res.ext}`;
  };

  const copyUrl = (e: any) => {
    copy(e.target.src);
    message.success(`复制Url:${e.target.src}`);
  };

  return (
    <div className="res">
      <div className="item" title="上传图片" style={{}}>
        <Upload
          name="avatar"
          listType="picture-card"
          showUploadList={false}
          action=""
        >
          <div>
            <PlusOutlined />
            <div className="ant-upload-text">上传图片</div>
          </div>
        </Upload>
      </div>
      {res.map((item) => (
        <div className="item">
          <img alt={item.name} src={getUrl(item)} onClick={copyUrl} />
        </div>
      ))}
    </div>
  );
};

export default ResPage;
