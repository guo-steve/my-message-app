import axios from "axios";

const baseUrl = process.env.REACT_APP_BACKEND_URL;

export const getMessages = async () => {
  const token = localStorage.getItem("token");

  try {
    const response = await axios.get(`${baseUrl}/v1/messages`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  } catch (error) {
    console.error(error);
    if (error.response?.status === 401) {
      localStorage.removeItem("token");
    }
    throw error;
  }
};

export const updateMessage = async (id, message) => {
  const token = localStorage.getItem("token");

  try {
    const response = await axios.patch(
      `${baseUrl}/v1/messages/${id}`,
      {
        content: message,
      },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    return response.data;
  } catch (error) {
    console.error(error);
    if (error.response?.status === 401) {
      localStorage.removeItem("token");
    }
    throw error;
  }
};

export const deleteMessage = async (id) => {
  const token = localStorage.getItem("token");

  try {
    const response = await axios.delete(`${baseUrl}/v1/messages/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  } catch (error) {
    console.error(error);
    if (error.response?.status === 401) {
      localStorage.removeItem("token");
    }
    throw error;
  }
};

export const postMessage = async (message) => {
  const token = localStorage.getItem("token");

  try {
    const response = await axios.post(
      `${baseUrl}/v1/messages`,
      {
        content: message,
      },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    return response.data;
  } catch (error) {
    console.error(error);
    if (error.response?.status === 401) {
      localStorage.removeItem("token");
    }
    throw error;
  }
};

export const register = async (values) => {
  try {
    const response = await axios.post(`${baseUrl}/v1/register`, values);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

export const login = async (values) => {
  try {
    const response = await axios.post(`${baseUrl}/v1/login`, values);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

export const logout = async () => {
  const token = localStorage.getItem("token");

  try {
    const response = await axios.post(`${baseUrl}/v1/logout`, null, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};
