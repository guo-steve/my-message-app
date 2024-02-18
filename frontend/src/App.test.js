import { render, screen } from "@testing-library/react";
import MatchMediaMock from "jest-matchmedia-mock";
import App from "./App";

describe("App", () => {
  let matchMedia;

  beforeAll(() => {
    matchMedia = new MatchMediaMock();
  });

  afterAll(() => {
    matchMedia.clear();
  });

  test("renders my-message-app", () => {
    render(<App />);
    const titleElem = screen.getByText("My Message App");
    expect(titleElem).toBeInTheDocument();
  });
});
