import { configureStore } from "@reduxjs/toolkit";
import messagesReducer from "../features/messages/messagesSlice";
import loginReducer from "../features/auth/loginSlice";

export default configureStore({
  reducer: {
    messages: messagesReducer,
    login: loginReducer,
  },
});
