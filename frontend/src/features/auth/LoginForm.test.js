import { render, screen } from "@testing-library/react";
import { Route, Routes } from "react-router";
import { BrowserRouter } from "react-router-dom";
import MatchMediaMock from "jest-matchmedia-mock";
import LoginForm from "./LoginForm";

describe("LoginForm", () => {
  let matchMedia;

  beforeAll(() => {
    matchMedia = new MatchMediaMock();
  });

  afterAll(() => {
    matchMedia.clear();
  });

  test("render LoginForm", () => {
    render(
      <BrowserRouter>
        <Routes>
          <Route index element={<LoginForm />} />
        </Routes>
      </BrowserRouter>
    );

    const usernameElem = screen.getByPlaceholderText("Username");
    expect(usernameElem).toBeInTheDocument();
    const passwordElem = screen.getByPlaceholderText("Password");
    expect(passwordElem).toBeInTheDocument();

    const buttonElem = screen.getByRole("button");
    expect(buttonElem).toBeInTheDocument();

    const registerLinkElem = screen.getByRole("link");
    expect(registerLinkElem).toBeInTheDocument();
  });
});
