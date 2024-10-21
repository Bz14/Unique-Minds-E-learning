"use client";
import { useState, useEffect } from "react";
import { MdVisibility, MdVisibilityOff } from "react-icons/md";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import { FieldErrors, useForm } from "react-hook-form";

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
  const [passwordVisible, setPasswordVisible] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const { control, handleSubmit, register, setValue, formState, reset } = form;
  const { errors, isDirty, isValid, isSubmitting, isSubmitSuccessful } =
    formState;

  useEffect(() => {
    if (isSubmitSuccessful) {
      reset();
    }
  }, [isSubmitSuccessful, reset]);

  const onError = (errors: FieldErrors<LoginForm>) => {
    console.log("Errors", errors);
  };

  const onSubmit = async (data: SignUpForm) => {
    setLoading(true);
    try {
      const response = await fetch(`${apiUrl}/api/auth/signup`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        setError("Signup failed");
        const data = await response.json();
        throw new Error(
          data.message || "An error occurred while creating your account."
        );
      }
      reset();
      router.push(`/verify?email=${encodeURIComponent(data.email)}`);
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
                placeholder="Create a password"
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
              >
                Login
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
