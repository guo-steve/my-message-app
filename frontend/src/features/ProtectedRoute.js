import { useNavigate } from "react-router";
import MainLayout from "./layout/Layout";

const ProtectedRoute = () => {
  const navigate = useNavigate();

  const isAuthenticated = false;

  if (!isAuthenticated) {
    navigate("/login");
  }

  return MainLayout;
};

export default ProtectedRoute;
