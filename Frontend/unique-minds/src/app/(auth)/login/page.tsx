"use client";
import { useState, useEffect } from "react";
import { MdVisibility, MdVisibilityOff } from "react-icons/md";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import { FieldErrors, useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import Spinner from "@/app/components/spinner/page";
import { FcGoogle } from "react-icons/fc";

type LoginForm = {
  email: string;
  password: string;
};

const schema = yup.object({
  email: yup
    .string()
    .required("Email is required.")
    .email("Invalid email format."),
  password: yup
    .string()
    .required("Password is required.")
    .min(8, "Password must be at least 8 characters.")
    .matches(
      /(?=.*[A-Z])/,
      "Password must contain at least one uppercase letter."
    )
    .matches(
      /(?=.*[a-z])/,
      "Password must contain at least one lowercase letter."
    )
    .matches(/(?=.*[0-9])/, "Password must contain at least one number.")
    .matches(
      /(?=.*[@$!%*?&])/,
      "Password must contain at least one special character."
    ),
});

const Login = () => {
  const form = useForm<LoginForm>({
    defaultValues: {
      email: "",
      password: "",
    },
    mode: "all",
    resolver: yupResolver(schema),
  });
  const apiUrl = process.env.NEXT_PUBLIC_API;
  const [passwordVisible, setPasswordVisible] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [token, setToken] = useState("");
  const [rememberMe, setRememberMe] = useState(false);

  const { handleSubmit, register, formState, reset } = form;
  const { errors, isDirty, isValid, isSubmitting, isSubmitSuccessful } =
    formState;
  const router = useRouter();

  useEffect(() => {
    if (isSubmitSuccessful) {
      reset();
    }
  }, [isSubmitSuccessful, reset]);

  const onError = (errors: FieldErrors<LoginForm>) => {
    console.log("Errors", errors);
  };

  const handleLoginWithGoogle = async () => {
    window.location.href = `${apiUrl}/api/auth/google`;
  };

  const onSubmit = async (loginData: LoginForm) => {
    setLoading(true);
    try {
      const response = await fetch(`${apiUrl}/api/auth/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(loginData),
      });
      if (!response.ok) {
        setError("Login failed");
        const data = await response.json();
        throw new Error(
          data.message || "An error occurred while logging you in."
        );
      }
      const data = await response.json();
      setToken(data.data);
      if (rememberMe) {
        localStorage.setItem("access_token", data.data);
      }
      reset();
      router.push(`/`);
    } catch (error) {
      setError("Something went wrong");
      console.log("Error", error);
      throw error;
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-white">
      <div className="shadow-2xl rounded-lg overflow-hidden max-w-lg w-full bg-white mx-auto">
        <div className="p-10 md:p-16">
          <h1 className="text-3xl font-bold mb-8 text-gray-800 text-center">
            Welcome Back!
          </h1>
          <form
            className="space-y-6"
            noValidate
            onSubmit={handleSubmit(onSubmit, onError)}
          >
            <div>
              <label
                htmlFor="email"
                className="block text-sm font-semibold text-gray-600"
              >
                Email Address
              </label>
              <input
                {...register("email")}
                type="email"
                placeholder="Enter your email"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
              <p style={{ color: "red", fontSize: "12px" }}>
                {errors.email?.message}
              </p>
            </div>
            <div className="relative">
              <label
                htmlFor="password"
                className="block text-sm font-semibold text-gray-600"
              >
                Password
              </label>
              <input
                {...register("password")}
                type={passwordVisible ? "password" : "text"}
                placeholder="Enter your password"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
              <div
                className="absolute inset-y-0 right-3 flex items-center cursor-pointer"
                onClick={() => setPasswordVisible(!passwordVisible)}
              >
                <span className="mt-6">
                  {passwordVisible ? (
                    <MdVisibilityOff size={24} className="text-gray-500" />
                  ) : (
                    <MdVisibility size={24} className="text-gray-500" />
                  )}
                </span>
              </div>
              <p style={{ color: "red", fontSize: "12px" }}>
                {errors.password?.message}
              </p>
            </div>

            <div className="flex items-center justify-between">
              <div className="flex items-center">
                <input
                  onChange={(e) => setRememberMe(e.target.checked)}
                  type="checkbox"
                  id="remember"
                  className="rounded border-gray-300 text-customBlue shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                />
                <label
                  htmlFor="remember"
                  className="ml-2 block text-sm text-gray-900"
                >
                  Remember me
                </label>
              </div>
            </div>
            <div>
              <button
                type="submit"
                className="w-full bg-customBlue text-white p-3 rounded-lg font-semibold hover:bg-gray-400"
                disabled={(!isDirty && !isValid) || isSubmitting}
              >
                {loading ? <Spinner /> : "Login"}
              </button>
              <div>{error && <p style={{ color: "red" }}>{error}</p>}</div>
            </div>
            <div className="relative my-6">
              <div className="absolute inset-0 flex items-center">
                <div className="w-full border-t border-gray-300"></div>
              </div>
              <div className="relative flex justify-center text-sm">
                <span className="bg-white px-2 text-gray-500">Or</span>
              </div>
            </div>
            <div>
              <button
                onClick={handleLoginWithGoogle}
                type="button"
                className="w-full py-3 px-6 bg-white text-black rounded-lg shadow-md border border-gray-300 hover:bg-gray-200 transition ease-in-out duration-300"
              >
                <div className="flex align-center justify-center">
                  <span className="mt-1 mr-2">
                    <FcGoogle />
                  </span>
                  <span>LogIn with Google</span>
                </div>
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;

{
  /* <DevTool control={control} /> */
}
