import React from "react";

import { BrowserRouter, Routes, Route } from "react-router-dom";

import LoginForm from "./features/auth/LoginForm";
import MessageApp from "./features/messages/MessageApp";

import "./App.css";
import ProtectedRoute from "./features/ProtectedRoute";
import MainLayout from "./features/layout/Layout";
import RegistrationForm from "./features/auth/RegistrationForm";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route element={<MainLayout />}>
            <Route index element={<LoginForm />} />
            <Route path="/login" element={<LoginForm />} />
            <Route path="/register" element={<RegistrationForm />} />
            <Route path="*" element={<LoginForm />} />
          </Route>
          <Route element={<ProtectedRoute />}>
            <Route path="/message" element={<MessageApp />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
