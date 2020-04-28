import React, { useState } from "react";
import { Spin, Card, Input, Button, message } from "antd";
import { UserOutlined, KeyOutlined } from "@ant-design/icons";
import "./index.scss";
import { useApi } from "../../api/login";

const Login = (props: any) => {
  const [isLoading, setLoading] = useState(false);
  const [name, setName] = useState("");
  const [pwd, setPwd] = useState("");

  const login = async () => {
    if (name.length > 0 && pwd.length > 0) {
      setLoading(true);
      const ok = await useApi.login(name, pwd);
      setLoading(false);
      if (ok) {
        message.success("登录成功", 1, () => {
          props.history.replace("/");
        });
        return;
      }
      message.error("登录失败");
      return;
    }
    message.error("输入用户名/密码");
  };

  return (
    <div className="login">
      <Spin tip="loading..." spinning={isLoading}>
        <Card title="登录" bordered={true} style={{ width: 400 }}>
          <Input
            id="name"
            size="large"
            placeholder="输入用户名"
            value={name}
            onChange={(e) => setName(e.target.value)}
            prefix={<UserOutlined />}
          />
          <br />
          <br />
          <Input.Password
            prefix={<KeyOutlined />}
            value={pwd}
            onChange={(e) => setPwd(e.target.value)}
            placeholder="密码"
            size="large"
            id="pwd"
            onPressEnter={login}
          ></Input.Password>
          <br />
          <br />
          <Button type="primary" size="large" block onClick={login}>
            登录
          </Button>
        </Card>
      </Spin>
    </div>
  );
};
export default Login;
