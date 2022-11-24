import axios from "axios";
import { FormEvent, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { z } from "zod";
import { ApiUrl } from "../api-url";
import { useAuthentication } from "../hooks/authentication";
import { SuccessType } from "../success-type";

interface RegisterForm {
  email: string;
  pwd: string;
}

interface RegisterResponse {
  access_token: string;
  success: number;
}

const Register = () => {
  const [email, setEmail] = useState("");
  const [pwd, setPwd] = useState("");
  const navigate = useNavigate();
  const { isLoggedIn } = useAuthentication();

  useEffect(() => {
    if (isLoggedIn) navigate("/");
  }, [navigate, isLoggedIn]);

  const validateForm = (form: RegisterForm): boolean => {
    const schema = z.object({
      email: z.string().email(),
      pwd: z.string().min(4),
    });

    const { success } = schema.safeParse(form);

    return success;
  };

  const onRegisterSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const valid = validateForm({
      email,
      pwd,
    });

    if (!valid) {
      console.error("register form is not valid");
      return;
    }

    const res = await axios.post<RegisterResponse>(
      ApiUrl.RegisterEndpoint,
      {
        email,
        password: pwd,
      },
      {
        withCredentials: true,
      }
    );

    if (res.data.success !== SuccessType.True) {
      console.error("registration is not successful");
      return;
    }

    navigate("/");
  };

  return (
    <form onSubmit={onRegisterSubmit}>
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
      <button>register</button>
    </form>
  );
};

export default Register;
