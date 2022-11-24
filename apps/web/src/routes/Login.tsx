import axios from "axios";
import { FormEvent, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { z } from "zod";
import { ApiUrl } from "../api-url";
import { useAuthentication } from "../hooks/authentication";
import { SuccessType } from "../success-type";

interface LoginForm {
  email: string;
  pwd: string;
}

interface LoginResponse {
  access_token: string;
  success: SuccessType;
}

const Login = () => {
  const [email, setEmail] = useState("");
  const [pwd, setPwd] = useState("");
  const navigate = useNavigate();

  const { isLoggedIn } = useAuthentication();

  useEffect(() => {
    if (isLoggedIn) navigate("/");
  }, [navigate, isLoggedIn]);

  const validateForm = (form: LoginForm): boolean => {
    const schema = z.object({
      email: z.string().email(),
      pwd: z.string(),
    });

    const { success } = schema.safeParse(form);

    return success;
  };

  const onLoginSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const valid = validateForm({
      email,
      pwd,
    });

    if (!valid) {
      console.error("login form is not valid");
      return;
    }

    try {
      const res = await axios.post<LoginResponse>(
        ApiUrl.LoginEndpoint,
        {
          email,
          password: pwd,
        },
        {
          withCredentials: true,
        }
      );

      if (res.data.success !== SuccessType.True) {
        console.error("login is not successful");
        return;
      }

      navigate("/");
    } catch (e) {
      if (axios.isAxiosError(e))
        console.log("perhaps email nor pwd is invalid");
    }
  };

  return (
    <form onSubmit={onLoginSubmit}>
      <div>
        <label htmlFor="email">email</label>
        <input
          type="text"
          id="email"
          onChange={(e) => setEmail(e.target.value)}
          value={email}
        />
      </div>
      <div>
        <label htmlFor="pwd">pwd</label>
        <input
          type="text"
          id="pwd"
          onChange={(e) => setPwd(e.target.value)}
          value={pwd}
        />
      </div>
      <button>login</button>
    </form>
  );
};

export default Login;
