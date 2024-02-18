import React from "react";
import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Col, Form, Input, Row, message } from "antd";
import { useNavigate } from "react-router";
import { login } from "../../services/api";

const LoginForm = () => {
  const [messageApi, contextHolder] = message.useMessage();

  const navigate = useNavigate();

  const onFinish = async (values) => {
    try {
      const { token } = await login({
        email: values.username,
        password: values.password,
      });

      localStorage.setItem("token", token);

      navigate("/message");
    } catch (error) {
      console.error(error);
      if (error.response?.status === 401) {
        messageApi.open({
          type: "error",
          content: "Invalid username or password",
        });
        return;
      }
      messageApi.open({
        type: "error",
        content: error.message,
      });
    }
  };
  return (
    <Row justify="center">
      <Col xs={20} sm={20} md={12} lg={8} xl={6} xxl={4}>
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
            marginLeft: "auto",
            marginRight: "auto",
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
            Or <a href="/register">register now!</a>
          </Form.Item>
        </Form>
      </Col>
    </Row>
  );
};

export default LoginForm;
