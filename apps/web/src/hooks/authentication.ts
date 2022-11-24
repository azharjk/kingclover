import axios from "axios";
import { useEffect, useState } from "react";
import { ApiUrl } from "../api-url";
import { UserInfoResponse } from "../interfaces/user-info";

export const useAuthentication = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const checkIfLoggedIn = async () => {
    try {
      const res = await axios.get<UserInfoResponse>(ApiUrl.UserInfoEndpoint, {
        withCredentials: true,
      });

      if (res.data.success) {
        setIsLoggedIn(true);
      }
    } catch (e) {}
  };

  useEffect(() => {
    checkIfLoggedIn();
  }, []);

  return {
    isLoggedIn,
  };
};
