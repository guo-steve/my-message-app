import React from "react";

import { BrowserRouter, Routes, Route } from "react-router-dom";

import LoginForm from "./features/login/LoginForm";
import MessageApp from "./features/messages/MessageApp";

import "./App.css";
import ProtectedRoute from "./features/ProtectedRoute";
import MainLayout from "./features/layout/Layout";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route element={<MainLayout />}>
            <Route index element={<LoginForm />} />
            <Route path="/login" element={<LoginForm />} />
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
