import React, { useState, useMemo } from "react";
import "./index.scss";
import "highlight.js/scss/github.scss";
import TextArea from "antd/lib/input/TextArea";
import { FullscreenOutlined, FileSearchOutlined } from "@ant-design/icons";
import { Button } from "antd";

const marked = require("marked");
const options = {
  renderer: new marked.Renderer(),
  pedantic: false,
  gfm: true,
  breaks: true,
  sanitize: false,
  smartLists: true,
  smartypants: false,
  xhtml: false,
  highlight: function (code: string, language: string) {
    const hljs = require("highlight.js");
    return hljs.highlightAuto(code).value;
  },
};
marked.setOptions(options);

const MarkDwonEditComp = (props: {
  content: string;
  setContent: (arg0: string) => void;
}) => {
  const { content, setContent } = props;
  const [fullScreen, setFullScreen] = useState(false);
  const [isPreview, setIsPreview] = useState(false);

  const inputChange = (e: any) => {
    setContent(e.target.value);
  };
  const scroll = (e: any) => {
    const mdEdit = document.getElementById("md-edit");
    const mdPreview = document.getElementById("md-preview");
    mdEdit?.scrollTo(e.target.scrollLeft, e.target.scrollTop);
    mdPreview?.scrollTo(e.target.scrollLeft, e.target.scrollTop);
  };
  const showPreview = useMemo(() => {
    return (
      <div
        id="md-preview"
        className="preview"
        style={{ display: isPreview ? "block" : "none" }}
        dangerouslySetInnerHTML={{ __html: marked(content) }}
        onScroll={scroll}
        onScrollCapture={scroll}
      />
    );
  }, [content, isPreview]);
  return (
    <div className={fullScreen ? "markdown fullScreen" : "markdown"}>
      <div className="control">
        <div className="title">文章内容</div>
        <div className="oparea">
          <Button
            icon={<FileSearchOutlined />}
            onClick={() => {
              setIsPreview(!isPreview);
            }}
          />
        </div>
        <div className="btn">
          <Button
            icon={<FullscreenOutlined />}
            onClick={() => {
              setFullScreen(!fullScreen);
            }}
          />
        </div>
      </div>
      <TextArea
        id="md-edit"
        className="edit"
        value={content}
        onInput={inputChange}
        wrap="off"
        onScroll={scroll}
      />
      {showPreview}
    </div>
  );
};

export default MarkDwonEditComp;
