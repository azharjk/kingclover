import { SuccessType } from "../success-type";

export interface UserInfo {
  id: string;
  email: string;
}

export interface UserInfoResponse {
  data: UserInfo;
  success: SuccessType;
}
