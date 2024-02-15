import React from "react";
import MessageApp from "./features/messages/MessageApp";
import { Layout } from "antd";

import "./App.css";

const { Header, Content, Footer } = Layout;

function App() {
  return (
    <div className="App">
      <Header>
        <h1 style={{ color: "white" }}>My Message App</h1>
      </Header>
      <Content>
        <MessageApp />
      </Content>
      <Footer>My Message App ©2021 Created by Me</Footer>
    </div>
  );
}

export default App;
