import React, { useState, useEffect } from "react";
import CheckableTag from "antd/lib/tag/CheckableTag";
import { PlusOutlined } from "@ant-design/icons";
import { Tag, Input, message } from "antd";
import { TagModel } from "../../../../types";
import { tagApi } from "../../../../api/tag";

const lodash = require("lodash");

const SelectTagsComp = (props: {
  selectTags: number[];
  setSelectTags: (arg0: number[]) => void;
}) => {
  const { selectTags, setSelectTags } = props;
  const [selectIds, setSelectIds] = useState(selectTags);
  const [inputValue, setInputValue] = useState("");
  const [inputVisible, setInputVisible] = useState(false);
  const [allTags, setAllTags] = useState<TagModel[]>([]);

  const handleInputConfirm = (e: any) => {
    if (e.target.value.trim().length === 0) {
      return;
    }
    tagApi
      .addTag(e.target.value.trim())
      .then((data) => {
        if (data != null) {
          getTags();
          setInputVisible(false);
          setInputValue("");
          message.success("添加标签成功");
          return;
        }
        message.error("添加标签失败");
      })
      .catch(() => {
        message.success("添加标签失败");
      });
  };

  useEffect(() => {
    getTags();
  }, []);

  const getTags = () => {
    tagApi.getTags().then((data: TagModel[]) => {
      setAllTags(data);
    });
  };

  const handleChange = (checked: boolean, tagId: number) => {
    if (checked) {
      if (selectIds.length === 5) {
        message.warning("最多选择5个标签");
        return;
      }
      selectIds.push(tagId);
    } else {
      lodash.pull(selectIds, tagId);
    }
    const ids = lodash.clone(selectIds);
    setSelectIds(ids);
    setSelectTags(ids);
  };

  const childs = allTags.map((item) => (
    <CheckableTag
      key={item.id}
      checked={selectIds.indexOf(item.id) >= 0}
      onChange={(checked) => handleChange(checked, item.id)}
    >
      {item.name}
    </CheckableTag>
  ));
  return (
    <div>
      {inputVisible && (
        <Input
          type="text"
          size="small"
          value={inputValue}
          onChange={(e) => {
            setInputValue(e.target.value);
          }}
          onBlur={handleInputConfirm}
          onPressEnter={handleInputConfirm}
        />
      )}
      {!inputVisible && (
        <Tag className="site-tag-plus" onClick={() => setInputVisible(true)}>
          <PlusOutlined /> 新标签
        </Tag>
      )}
      {childs}
    </div>
  );
};

export default SelectTagsComp;
