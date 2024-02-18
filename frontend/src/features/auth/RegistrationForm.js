import React, { useState } from "react";
import { Button, Checkbox, Col, Form, Input, message, Row } from "antd";
import { register } from "../../services/api";

const formItemLayout = {
  labelCol: {
    xs: {
      span: 24,
    },
    sm: {
      span: 8,
    },
  },
  wrapperCol: {
    xs: {
      span: 24,
    },
    sm: {
      span: 16,
    },
  },
};
const tailFormItemLayout = {
  wrapperCol: {
    xs: {
      span: 24,
      offset: 0,
    },
    sm: {
      span: 16,
      offset: 8,
    },
  },
};

const RegistrationForm = () => {
  const [form] = Form.useForm();
  const [captcha, setCaptcha] = useState("");

  const [messageApi, contextHolder] = message.useMessage();

  const onFinish = async (values) => {
    try {
      const data = await register(values);

      console.log(data);

      messageApi.open({
        type: "success",
        content: "Registration successful",
      });
    } catch (error) {
      console.error(error);
      if (error.response?.status === 400) {
        messageApi.open({
          type: "error",
          content: "Invalid input",
        });
        return;
      }
      messageApi.open({
        type: "error",
        content: error.message,
      });
    }
  };

  const onGetCaptcha = () => {
    const newCaptcha = Math.random().toString(36).slice(2, 8);
    setCaptcha(newCaptcha);
  };

  return (
    <Form
      {...formItemLayout}
      form={form}
      name="register"
      onFinish={onFinish}
      style={{
        maxWidth: 600,
        marginTop: "30px",
        marginBottom: "30px",
      }}
      scrollToFirstError
    >
      {contextHolder}
      <Form.Item
        name="email"
        label="E-mail"
        rules={[
          {
            type: "email",
            message: "The input is not valid E-mail!",
          },
          {
            required: true,
            message: "Please input your E-mail!",
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name="password"
        label="Password"
        rules={[
          {
            required: true,
            message: "Please input your password!",
          },
        ]}
        hasFeedback
      >
        <Input.Password />
      </Form.Item>
      <Form.Item
        name="confirm"
        label="Confirm Password"
        dependencies={["password"]}
        hasFeedback
        rules={[
          {
            required: true,
            message: "Please confirm your password!",
          },
          ({ getFieldValue }) => ({
            validator(_, value) {
              if (!value || getFieldValue("password") === value) {
                return Promise.resolve();
              }
              return Promise.reject(
                new Error("The new password that you entered do not match!")
              );
            },
          }),
        ]}
      >
        <Input.Password />
      </Form.Item>
      <Form.Item
        name="full_name"
        label="Full Name"
        tooltip="Your full name."
        rules={[
          {
            required: true,
            message: "Please input your full name!",
            whitespace: true,
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Captcha"
        extra="We must make sure that your are a human."
      >
        <Row gutter={8}>
          <Col span={12}>
            <Form.Item
              name="captcha"
              noStyle
              rules={[
                {
                  required: true,
                  message: "Please input the captcha you got!",
                },
              ]}
            >
              <Input value={captcha} />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Button onClick={onGetCaptcha}>Get captcha</Button>
          </Col>
        </Row>
      </Form.Item>
      <Form.Item
        name="agreement"
        valuePropName="checked"
        rules={[
          {
            validator: (_, value) =>
              value
                ? Promise.resolve()
                : Promise.reject(new Error("Should accept agreement")),
          },
        ]}
        {...tailFormItemLayout}
      >
        <Checkbox>
          I have read the <a href="/terms.txt">agreement</a>
        </Checkbox>
      </Form.Item>

      <Form.Item {...tailFormItemLayout}>
        <Button type="primary" htmlType="submit">
          Register
        </Button>
        Or <a href="/login">Go back to login!</a>
      </Form.Item>
    </Form>
  );
};
export default RegistrationForm;
