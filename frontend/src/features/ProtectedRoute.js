import { useNavigate } from "react-router";
import MainLayout from "./layout/Layout";
import { useEffect } from "react";

const ProtectedRoute = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    const isAuthenticated = token != null;

    if (!isAuthenticated) {
      navigate("/login");
    }
  }, [navigate]);

  return <MainLayout />;
};

export default ProtectedRoute;
