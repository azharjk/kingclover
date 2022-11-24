import axios from "axios";
import { useCallback, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ApiUrl } from "../api-url";
import { ErrorDetail } from "../interfaces/error-detail";
import { UserInfo, UserInfoResponse } from "../interfaces/user-info";
import { SuccessType } from "../success-type";

interface LogoutResponse {
  success: SuccessType;
}

const Root = () => {
  const [userInfo, setUserInfo] = useState<UserInfo | null>(null);
  const navigate = useNavigate();

  const refreshToken = async () => {
    const res = await axios.post(ApiUrl.TokenEndpoint, {}, { withCredentials: true });
    console.log(res);
  };

  const getUserInfo = useCallback(async () => {
    try {
      const res = await axios.get<UserInfoResponse>(ApiUrl.UserInfoEndpoint, {
        withCredentials: true,
      });

      if (!res.data.success) {
        console.log("i think u are unauthorize");
        throw new Error("Request did not succeded");
      }

      setUserInfo(res.data.data);
    } catch (e) {
      if (axios.isAxiosError(e)) {
        const data = e.response?.data as ErrorDetail;
        if (
          data.success !== SuccessType.True &&
          data.error.endsWith("token expires")
        ) {
          await refreshToken();
          navigate("/");
          return;
        }
      }
      navigate("/register");
    }
  }, [navigate]);

  useEffect(() => {
    getUserInfo();
  }, [getUserInfo]);

  const onLogoutClick = async () => {
    const res = await axios.post<LogoutResponse>(
      ApiUrl.LogoutEndpoint,
      {},
      {
        withCredentials: true,
      }
    );

    if (!res.data.success) {
      console.error("logout was not successful");
      return;
    }

    navigate("/register");
  };

  return (
    <div>
      {userInfo ? (
        <div>
          <p>ID: {userInfo.id}</p>
          <p>EMAIL: {userInfo.email}</p>
          <button type="button" onClick={onLogoutClick}>
            logout
          </button>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};

export default Root;
