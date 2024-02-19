import { useEffect } from "react";
import { logout } from "../../services/api";
import { message } from "antd";

const Logout = () => {
  const [messageApi, contextHolder] = message.useMessage();

  useEffect(() => {
    logout()
      .catch((error) => {
        messageApi.error(error.message);
      })
      .finally(() => {
        localStorage.removeItem("token");
      });
  }, [messageApi]);

  return (
    <div>
      <h1>Bye!</h1>

      <a href="/login">Login again!</a>

      {contextHolder}
    </div>
  );
};

export default Logout;
