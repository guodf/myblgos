import React from "react";
import "./App.css";
import { HashRouter, Switch, Redirect, Route } from "react-router-dom";
import authState from "./state/auth";
import routers from "./router";

function App() {
  return (
    <HashRouter>
      <Switch>
        {routers.map((router) => (
          <Route
            key={router.path}
            path={router.path}
            render={(props) => {
              if (router.path !== "/login") {
                if (router.auth && !authState.isLogin()) {
                  return <Redirect to="/login" />;
                }
                return <router.component {...props} routers={router.routers} />;
              }
              if (authState.isLogin()) {
                return <Redirect to="/" />;
              }
              return <router.component {...props} routers={router.routers} />;
            }}
          />
        ))}
      </Switch>
    </HashRouter>
  );
}

export default App;
