import React, { useEffect, useState } from "react";
import { Button, Form, Input, message, Table } from "antd";
import { useDispatch, useSelector } from "react-redux";
import { addMessage, syncMessages } from "./messagesSlice";

import { getMessages, postMessage } from "../../services/api";

const MessageApp = () => {
  const [messageApi, contextHolder] = message.useMessage();
  const messages = useSelector((state) => state.messages.value);
  const dispatch = useDispatch();

  const token = localStorage.getItem("token");

  useEffect(() => {
    getMessages()
      .then((data) => {
        data = data.map((m) => (m = { ...m, key: m.id }));
        dispatch(syncMessages(data));
      })
      .catch((error) => {
        console.error(error);
        const errMsg =
          error.response?.data.message || error.message || "An error occurred";
        messageApi.open({
          type: "error",
          content: errMsg,
        });
      });
  }, [dispatch, messageApi, token]);

  const columns = [
    {
      title: "ID",
      dataIndex: "id",
      key: "id",
    },
    {
      title: "Message",
      dataIndex: "content",
      key: "content",
    },
    {
      title: "Posted By",
      dataIndex: "created_by",
      key: "created_by",
    },
    {
      title: "Created At",
      dataIndex: "created_at",
      key: "created_at",
      render: (text) => new Date(text).toLocaleString(),
    },
  ];

  const [form] = Form.useForm();
  const [formLayout, setFormLayout] = useState("horizontal");
  const onFormLayoutChange = ({ layout }) => {
    setFormLayout(layout);
  };

  const onFinish = async (values) => {
    try {
      const newMessage = await postMessage(values.message);
      newMessage.key = newMessage.id;
      dispatch(addMessage(newMessage));
    } catch (error) {
      console.error(error);
      const errMsg =
        error.response?.data.message || error.message || "An error occurred";
      messageApi.open({
        type: "error",
        content: errMsg,
      });
    }
  };

  return (
    <div>
      {contextHolder}
      <Form
        layout={formLayout}
        form={form}
        initialValues={{
          layout: formLayout,
        }}
        onValuesChange={onFormLayoutChange}
        onFinish={onFinish}
        style={{
          marginTop: "30px",
          marginBottom: "20px",
          marginLeft: "30px",
          marginRight: "30px",
        }}
      >
        <Form.Item
          label="Message"
          name="message"
          rules={[
            {
              required: true,
              message: "Please input your message!",
            },
          ]}
        >
          <Input placeholder="Write the message content here" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>

      <Table columns={columns} dataSource={messages} />
    </div>
  );
};

export default MessageApp;
