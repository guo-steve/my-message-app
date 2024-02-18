import { render, screen } from "@testing-library/react";
import { Route, Routes } from "react-router";
import { BrowserRouter } from "react-router-dom";
import MatchMediaMock from "jest-matchmedia-mock";
import RegistrationForm from "./RegistrationForm";

describe("RegistrationForm", () => {
  let matchMedia;

  beforeAll(() => {
    matchMedia = new MatchMediaMock();
  });

  afterAll(() => {
    matchMedia.clear();
  });

  test("render RegistrationForm", () => {
    render(
      <BrowserRouter>
        <Routes>
          <Route index element={<RegistrationForm />} />
        </Routes>
      </BrowserRouter>
    );

    const emailLabel = screen.getByText("E-mail");
    expect(emailLabel).toBeInTheDocument();

    const confirmPasswordLabel = screen.getByText("Confirm Password");
    expect(confirmPasswordLabel).toBeInTheDocument();
  });
});
