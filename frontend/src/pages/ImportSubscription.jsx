import React, { useState } from "react";
import { Form, Input, Button, message } from "antd";
import { importSubscription } from "../api/api";

const ImportSubscription = () => {
  const [loading, setLoading] = useState(false);

  const handleImport = async (values) => {
    setLoading(true);
    try {
      const response = await importSubscription(values.url);
      message.success(response.message);
    } catch (error) {
      message.error(error.response.data.error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <h2>Import Subscription</h2>
      <Form onFinish={handleImport}>
        <Form.Item
          name="url"
          label="Subscription URL"
          rules={[{ required: true, message: "Please input the subscription URL!" }]}
        >
          <Input />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" loading={loading}>
            Import
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default ImportSubscription;
