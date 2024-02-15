import { createSlice } from "@reduxjs/toolkit";

export const messagesSlice = createSlice({
  name: "messages",
  initialState: {
    value: [],
  },
  reducers: {
    syncMessages: (state, action) => {
      state.value = action.payload;
    },
    addMessage: (state, action) => {
      state.value.push(action.payload);
    },
  },
});

// Action creators are generated for each case reducer function
export const { syncMessages, addMessage } = messagesSlice.actions;

export default messagesSlice.reducer;
