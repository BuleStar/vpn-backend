import React, { useState, useEffect } from "react";
import { Table, Button, Modal, Form, Input } from "antd";
import axios from "axios";

const NodeManagement = () => {
  const [nodes, setNodes] = useState([]);
  const [visible, setVisible] = useState(false);
  const [form] = Form.useForm();

  useEffect(() => {
    fetchNodes();
  }, []);

  const fetchNodes = async () => {
    const res = await axios.get("/api/nodes");
    setNodes(res.data.nodes);
  };

  const handleAddNode = async (values) => {
    await axios.post("/api/nodes", values);
    fetchNodes();
    setVisible(false);
  };

  return (
    <div>
      <Button type="primary" onClick={() => setVisible(true)}>Add Node</Button>
      <Table
        dataSource={nodes}
        columns={[
          { title: "Name", dataIndex: "name" },
          { title: "Address", dataIndex: "address" },
        ]}
      />
      <Modal
        visible={visible}
        title="Add Node"
        onCancel={() => setVisible(false)}
        onOk={() => form.submit()}
      >
        <Form form={form} onFinish={handleAddNode}>
          <Form.Item name="name" label="Name">
            <Input />
          </Form.Item>
          <Form.Item name="address" label="Address">
            <Input />
          </Form.Item>
          <Form.Item name="port" label="Port">
            <Input />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};

export default NodeManagement;
