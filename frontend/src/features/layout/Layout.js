import { Layout } from "antd";
import { Outlet } from "react-router";

const { Header, Content, Footer } = Layout;

const MainLayout = () => {
  return (
    <>
      <Header>
        <h1 style={{ color: "white" }}>My Message App</h1>
      </Header>
      <Content>
        <Outlet />
      </Content>
      <Footer>My Message App ©2021 Created by Me</Footer>
    </>
  );
};

export default MainLayout;
