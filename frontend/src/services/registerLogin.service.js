import { authHeader } from '../helpers/auth-header.js';
import axios from 'axios';

export const authService = {
  login,
  register,
  logout,
};

async function login(username, password) {
  try {
    let respone = await axios.post(
      `${process.env.VUE_APP_API_BACKEND}/auth/login/`,
      {
        username: username,
        password: password,
      },
      { 'content-type': 'application/x-www-form-urlencoded' }
    );
    let user = respone.data;
    if (respone.data.success) {
      localStorage.setItem('user', JSON.stringify(user));
      return user;
    }
    return null;
  } catch (error) {
    throw error;
  }
}

async function register(username, password, email) {
  let response = await axios.post(
    `${process.env.VUE_APP_API_BACKEND}/auth/register/`,
    {
      username: username,
      email: email,
      password: password,
    },
    { 'content-type': 'application/x-www-form-urlencoded' }
  );
  return handleResponse(response);
}

async function logout() {
  localStorage.removeItem('user');
}

async function handleResponse(response) {
  return new Promise((resolve, reject) => {
    let data = response.data;
    if (response.status === 200) {
      if (data.success) {
        resolve(data);
      } else {
        reject(data.msg);
      }
      if (errors) {
        reject(errors);
      }
    } else if (response.status === 401) {
      // logout();
      location.reload(true);
    } else {
      const error = (data && data.msg) || response.statusText;
      reject(error);
    }
  });
}
