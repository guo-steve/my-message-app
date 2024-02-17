// import { render, screen, waitFor, fireEvent } from "@testing-library/react";
// import { Provider } from "react-redux";
// import axios from "axios";
// import MockAdapter from "axios-mock-adapter";
// import configureStore from "redux-mock-store";
// import thunk from "redux-thunk";
// import MessageApp from "./MessageApp";
//
// const mockAxios = new MockAdapter(axios);
// const mockStore = configureStore([thunk]);
//
// describe("MessageApp", () => {
//   let store;
//
//   beforeEach(() => {
//     store = mockStore({
//       messages: {
//         value: [
//           { id: 1, content: "Hello", created_at: "2022-01-01T00:00:00Z" },
//           { id: 2, content: "World", created_at: "2022-01-02T00:00:00Z" },
//         ],
//       },
//     });
//   });
//
//   afterEach(() => {
//     mockAxios.reset();
//   });
//
//   test("renders message form and table", async () => {
//     render(
//       <Provider store={store}>
//         <MessageApp />
//       </Provider>
//     );
//
//     // Check if the message form is rendered
//     const messageInput = screen.getByLabelText("Message");
//     expect(messageInput).toBeInTheDocument();
//
//     // Check if the table is rendered with the existing messages
//     const messageRows = screen.getAllByRole("row");
//     expect(messageRows).toHaveLength(3); // Including the table header row
//
//     // Check if the existing messages are displayed correctly
//     const messageCells = screen.getAllByRole("cell");
//     expect(messageCells[0]).toHaveTextContent("1");
//     expect(messageCells[1]).toHaveTextContent("Hello");
//     expect(messageCells[2]).toHaveTextContent("2022-01-01 00:00:00");
//     expect(messageCells[3]).toHaveTextContent("2");
//     expect(messageCells[4]).toHaveTextContent("World");
//     expect(messageCells[5]).toHaveTextContent("2022-01-02 00:00:00");
//   });
//
//   test("submits a new message", async () => {
//     mockAxios.onGet("/v1/messages").reply(200, []);
//     mockAxios.onPost("/v1/messages").reply(200, {
//       id: 3,
//       content: "New Message",
//       created_at: "2022-01-03T00:00:00Z",
//     });
//
//     render(
//       <Provider store={store}>
//         <MessageApp />
//       </Provider>
//     );
//
//     // Enter a new message in the input field
//     const messageInput = screen.getByLabelText("Message");
//     fireEvent.change(messageInput, { target: { value: "New Message" } });
//
//     // Submit the form
//     const submitButton = screen.getByText("Submit");
//     fireEvent.click(submitButton);
//
//     // Wait for the form submission and message retrieval to complete
//     await waitFor(() => {
//       expect(mockAxios.history.post.length).toBe(1);
//     });
//
//     await waitFor(() => {
//       expect(mockAxios.history.get.length).toBe(1);
//     });
//
//     // Check if the new message is added to the table
//     const messageCells = screen.getAllByRole("cell");
//     expect(messageCells[6]).toHaveTextContent("3");
//     expect(messageCells[7]).toHaveTextContent("New Message");
//     expect(messageCells[8]).toHaveTextContent("2022-01-03 00:00:00");
//   });
// });
