import React, { useState, useEffect } from "react";
import { Select, Input, Divider, Button, message } from "antd";
import { CategoryModel } from "../../../../types";
import { categoryApi } from "../../../../api/category";
const { Option } = Select;

const SelectCategoryComp = (props: {
  defaultValue: number;
  onSelectCategory: (id: number) => void;
}) => {
  const { defaultValue, onSelectCategory } = props;
  const [categoryName, setCategoryName] = useState("");
  const [categories, setCategories] = useState<CategoryModel[]>([]);
  const addCategory = () => {
    categoryApi
      .addCategory(categoryName)
      .then((data) => {
        if (data != null) {
          message.success("添加文章分类成功");
          setCategoryName("");
          getCategories();
          return;
        }
        message.error("添加文字分类失败");
      })
      .catch(() => {
        message.error("添加文字分类失败");
      });
  };

  useEffect(() => {
    getCategories();
  }, []);

  const getCategories = () => {
    categoryApi.getCategories().then((data) => {
      setCategories(data);
    });
  };
  return (
    <Select
      style={{ width: "30%" }}
      placeholder="文章类型"
      value={defaultValue === 0 ? "" : defaultValue}
      onSelect={(e: any) => {
        onSelectCategory(e);
      }}
      dropdownRender={(menu) => (
        <div>
          <div style={{ display: "flex", flexWrap: "nowrap", padding: 8 }}>
            <Input
              style={{ flex: "auto" }}
              value={categoryName}
              onChange={(e: any) => {
                setCategoryName(e.target.value);
              }}
            />
            <Button type="primary" onClick={addCategory}>
              添加
            </Button>
          </div>
          <Divider style={{ margin: "4px 0" }} />
          {menu}
        </div>
      )}
    >
      {categories.map((category) => (
        <Option key={category.id} value={category.id}>
          {category.name}
        </Option>
      ))}
    </Select>
  );
};
export default SelectCategoryComp;
