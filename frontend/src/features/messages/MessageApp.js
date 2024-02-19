import React, { useEffect, useState } from "react";
import { Button, Form, Input, message, Popconfirm, Table } from "antd";
import { useDispatch, useSelector } from "react-redux";
import { addMessage, syncMessages } from "./messagesSlice";
import {
  deleteMessage,
  getMessages,
  postMessage,
  updateMessage,
} from "../../services/api";
import { EditableCell, EditableRow } from "./EditableCell";

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

  const handleDelete = async (row) => {
    try {
      await deleteMessage(row.id);
      const newData = messages.filter((item) => item.key !== row.key);
      dispatch(syncMessages(newData));
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

  const handleSave = async (row) => {
    try {
      await updateMessage(row.id, row.content);

      const newData = [...messages];
      const index = newData.findIndex((item) => row.key === item.key);
      const item = newData[index];
      // this updates messages
      newData.splice(index, 1, {
        ...item,
        ...row,
      });
      dispatch(syncMessages(newData));
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

  const columns = [
    {
      title: "ID",
      dataIndex: "id",
    },
    {
      title: "Message",
      dataIndex: "content",
      editable: true,
      onCell: (record) => ({
        record,
        editable: true,
        dataIndex: "content",
        title: "Message",
        handleSave,
      }),
    },
    {
      title: "Posted By",
      dataIndex: "created_by",
    },
    {
      title: "Created At",
      dataIndex: "created_at",
      render: (text) => new Date(text).toLocaleString(),
    },
    {
      title: "Action",
      dataIndex: "action",
      render: (_, record) =>
        messages.length >= 1 ? (
          <Popconfirm
            title="Sure to delete?"
            onConfirm={() => handleDelete(record)}
          >
            <Button>Delete</Button>
          </Popconfirm>
        ) : null,
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

  const components = {
    body: {
      row: EditableRow,
      cell: EditableCell,
    },
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

      <Table components={components} columns={columns} dataSource={messages} />
    </div>
  );
};

export default MessageApp;
