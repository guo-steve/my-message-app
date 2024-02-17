import React from "react";

import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { login } from "../login/loginSlice";

import { Button, Card, Form, Input, message } from "antd";
import axios from "axios";
import { useDispatch } from "react-redux";

const cardStyle = {
  width: "100%",
  boxShadow: "none",
};

const baseUrl = process.env.REACT_APP_BACKEND_URL;

const LoginForm = () => {
  const [messageApi, contextHolder] = message.useMessage();

  const dispatch = useDispatch();

  const onFinish = async (values) => {
    try {
      const response = await axios.post(`${baseUrl}/v1/login`, values);

      const user = response.data;

      dispatch(login(user));
    } catch (error) {
      console.error(error);
      const errorMessage = error.response?.data.message || error.message;
      messageApi.open({
        type: "error",
        content: errorMessage,
      });
    }
  };
  return (
    <Card
      bordered={false}
      style={cardStyle}
      styles={{
        body: {
          paddingLG: "300",
        },
      }}
    >
      {contextHolder}
      <Form
        name="normal_login"
        className="login-form"
        initialValues={{
          remember: true,
        }}
        onFinish={onFinish}
        style={{
          marginTop: "30px",
          marginBottom: "30px",
          marginLeft: "30px",
          marginRight: "30px",
        }}
      >
        <Form.Item
          name="username"
          rules={[
            {
              required: true,
              message: "Please input your Username!",
            },
          ]}
          hasFeedback
        >
          <Input
            prefix={<UserOutlined className="site-form-item-icon" />}
            placeholder="Username"
          />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[
            {
              required: true,
              message: "Please input your Password!",
            },
          ]}
          hasFeedback
        >
          <Input.Password
            prefix={<LockOutlined className="site-form-item-icon" />}
            type="password"
            placeholder="Password"
          />
        </Form.Item>

        <Form.Item>
          <Button
            type="primary"
            htmlType="submit"
            className="login-form-button"
          >
            Log in
          </Button>
        </Form.Item>
      </Form>
    </Card>
  );
};

export default LoginForm;
