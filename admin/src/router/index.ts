import AdminLayout from "../layout/admin";
import Login from "../pages/login";
import NotFound from "../pages/nofoud";
import { AndroidFilled, PictureOutlined, SettingOutlined } from "@ant-design/icons";
import ArticlesPage from "../pages/articles";
import ArticleEdit from "../pages/articles/edit";
import ResPage from "../pages/res";
import SettingsPage from "../pages/settings";
const routers = [
  {
    path: "/login",
    component: Login,
    auth: false,
  },
  {
    path: "/",
    component: AdminLayout,
    auth: true,
    routers: [
      {
        path: "/articles",
        name: "文章列表",
        exact: true,
        icon: AndroidFilled,
        component: ArticlesPage,
        routers: [{
          path: "/article/edid",
          name: "编辑文章",
          link: false,
          component: ArticleEdit,
        }]
      },
      {
        path: "/res",
        name: "素材库",
        exact: "true",
        icon: PictureOutlined,
        component: ResPage,
      }, {
        path: "/settings",
        name: "设置",
        exact: "true",
        icon: SettingOutlined,
        component: SettingsPage
      }
    ]
  },
  {
    path: "*",
    component: NotFound,
    auth: false,
  },
];

export default routers;
