import { Layout } from "antd";
import { useEffect, useState } from "react";
import { Outlet } from "react-router";

const { Header, Content, Footer } = Layout;

const MainLayout = () => {
  const [loggedIn, setLoggedIn] = useState(false);

  useEffect(() => {
    setLoggedIn(!!localStorage.getItem("token"));
  }, [loggedIn]);

  return (
    <Layout className="layout">
      <Header>
        <h2 style={{ color: "white", marginTop: 0, display: "inline-block" }}>
          My Message App
        </h2>
        <span style={{ float: "right" }}>
          {loggedIn && (
            <a style={{ color: "white" }} href="/logout">
              Logout
            </a>
          )}
        </span>
      </Header>
      <Content>
        <Outlet />
      </Content>
      <Footer>My Message App ©2021 Created by Me</Footer>
    </Layout>
  );
};

export default MainLayout;
