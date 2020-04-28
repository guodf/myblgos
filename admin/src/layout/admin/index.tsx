import React, { useState, useEffect } from "react";
import { Layout, Menu } from "antd";
import { Link, Switch, useLocation, Route, Redirect } from "react-router-dom";
import {
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  LogoutOutlined,
} from "@ant-design/icons";

import "./index.scss";
import MenuItem from "antd/lib/menu/MenuItem";
import { RouterProps } from "../../types";
import authState from "../../state/auth";
import ArticlesPage from "../../pages/articles";
import ArticleEdit from "../../pages/articles/edit";
import ResPage from "../../pages/res";
import SettingsPage from "../../pages/settings";

const { Header, Sider, Content } = Layout;

function AdminLayout(props: any) {
  const { routers } = props;
  const [showSider, setShowSider] = useState(true);
  const [selectedKeys, setSelectedKeys] = useState([""]);
  const location = useLocation();

  let subMenudKeys = [""];
  if (location.pathname.startsWith("/pack/")) {
    subMenudKeys = ["/pack"];
  }

  useEffect(() => {
    setSelectedKeys([location.pathname]);
  }, [location]);

  const selectMenuItem = (item: { key: string }) => {
    setSelectedKeys([item.key]);
  };
  return (
    <>
      <Layout style={{ height: "100%" }}>
        <Sider
          style={{ height: "100vh" }}
          trigger={null}
          collapsible
          collapsed={!showSider}
          className="sider"
        >
          <div className="logo">
            <Link to="/">
              <h1>在路上</h1>
            </Link>
          </div>
          <Menu
            theme="dark"
            mode="inline"
            defaultOpenKeys={subMenudKeys}
            selectedKeys={selectedKeys}
            onClick={selectMenuItem}
          >
            {routers
              ?.filter((item: any) => item.link === undefined || item.link)
              .map((router: any, index: any) => (
                <MenuItem key={router.path}>
                  <Link
                    to={{
                      pathname: router.path,
                    }}
                  >
                    <router.icon />
                    <span style={{ marginLeft: 10 }}>{router.name}</span>
                  </Link>
                </MenuItem>
              ))}
          </Menu>
        </Sider>
        <Layout className="site-layout main">
          <Header
            className="site-layout-background"
            style={{
              padding: 0,
            }}
          >
            <div className="header">
              {React.createElement(
                showSider ? MenuUnfoldOutlined : MenuFoldOutlined,
                {
                  className: "trigger",
                  onClick: () => {
                    setShowSider(!showSider);
                  },
                }
              )}
              <div style={{ flex: "1 1 0%" }}></div>
              <div className="header-right">
                <span style={{ verticalAlign: "middle" }}>
                  {authState?.authModel?.name}
                </span>
              </div>
              <LogoutOutlined
                className="trigger"
                size={48}
                onClick={() => {
                  authState.logout();
                  props.history.replace("/login");
                }}
              />
            </div>
          </Header>

          <Content
            className="site-layout-background"
            style={{
              margin: "24px 16px",
            }}
          >
            <Switch>
              <Redirect exact from="/" to="/articles" />
              <Route path="/articles" component={ArticlesPage} />
              <Route path="/article/edit/:articleId" component={ArticleEdit} />
              <Route
                exact={true}
                path="/article/edit/:articleId"
                component={ArticleEdit}
              />
              <Route path="/article/edit" component={ArticleEdit} />
              <Route path="/res" component={ResPage} />

              <Route path="/settings" component={SettingsPage} />
            </Switch>
          </Content>
        </Layout>
      </Layout>
    </>
  );
}

export default AdminLayout;
