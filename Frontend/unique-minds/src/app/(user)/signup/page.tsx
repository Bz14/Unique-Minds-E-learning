"use client";
import Link from "next/link";
import { FcGoogle } from "react-icons/fc";
import { MdVisibility, MdVisibilityOff } from "react-icons/md";
import { FieldErrors, useForm } from "react-hook-form";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import Spinner from "@/app/components/spinner/page";

type SignUpForm = {
  fullName: string;
  email: string;
  password: string;
  confirmPassword: string;
  userType: string;
};

const schema = yup.object({
  fullName: yup.string().required("Full name is required."),
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
  confirmPassword: yup
    .string()
    .oneOf([yup.ref("password"), undefined], "Passwords must match.")
    .required("Please confirm your password."),
  userType: yup.string().required("User type is required."),
});
const SignUp = () => {
  const apiUrl = process.env.NEXT_PUBLIC_API;
  const form = useForm<SignUpForm>({
    defaultValues: {
      fullName: "",
      email: "",
      password: "",
      confirmPassword: "",
      userType: "student",
    },
    mode: "all",
    resolver: yupResolver(schema),
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const { register, handleSubmit, setValue, formState, reset } = form;
  const { errors, isDirty, isValid, isSubmitting, isSubmitSuccessful } =
    formState;
  const [passwordVisible, setPasswordVisible] = useState(true);
  const [confirmPasswordVisible, setConfirmPasswordVisible] = useState(true);
  const [userType, setUserType] = useState("student");
  const router = useRouter();

  useEffect(() => {
    if (isSubmitSuccessful) {
      reset();
    }
  }, [isSubmitSuccessful, reset]);

  const onError = (errors: FieldErrors<SignUpForm>) => {
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

  const handleSignUpWithGoogle = async () => {
    try {
      const response = await fetch(`${apiUrl}/api/auth/google`);
      if (!response.ok) {
        setError("Google Signup failed");
        const data = await response.json();
        throw new Error(
          data.message || "An error occurred while creating your account."
        );
      }
    } catch (error) {
      setError("Something went wrong");
      console.log("Error", error);
      throw error;
    }
  };

  const handleUserType = (type: string) => {
    setUserType(type);
    setValue("userType", type);
  };
  return (
    <div className="min-h-screen flex items-center justify-center bg-white">
      <div className="shadow-2xl rounded-lg overflow-hidden max-w-lg w-full bg-white mt-10 mx-auto">
        <div className="p-10 md:p-16">
          <h1 className="text-3xl font-bold mb-8 text-gray-800 text-center">
            Create Your Account
          </h1>
          <form
            className="space-y-6"
            onSubmit={handleSubmit(onSubmit, onError)}
            noValidate
          >
            <div>
              <label
                htmlFor="full-name"
                className="block text-sm font-semibold text-gray-600"
              >
                Full Name
              </label>
              <input
                type="text"
                placeholder="Enter your full name"
                {...register("fullName")}
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
              <p style={{ color: "red", fontSize: "12px" }}>
                {errors.fullName?.message}
              </p>
            </div>
            <div>
              <label
                htmlFor="email"
                className="block text-sm font-semibold text-gray-600"
              >
                Email Address
              </label>
              <input
                {...register("email", {
                  validate: {
                    emailAvailable: async (fieldValue) => {
                      const response = await fetch(
                        `${apiUrl}/api/auth/email?email=${fieldValue}`
                      );
                      const data = await response.json();
                      return data.length == 0 || "Email already exists.";
                    },
                  },
                })}
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
            <div className="relative">
              <label
                htmlFor="confirm-password"
                className="block text-sm font-semibold text-gray-600"
              >
                Confirm Password
              </label>
              <input
                {...register("confirmPassword")}
                type={confirmPasswordVisible ? "password" : "text"}
                placeholder="Confirm your password"
                className="mt-2 block w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 text-gray-700"
              />
              <div
                className="absolute inset-y-0 right-3 flex items-center cursor-pointer"
                onClick={() =>
                  setConfirmPasswordVisible(!confirmPasswordVisible)
                }
              >
                <span className="mt-6">
                  {confirmPasswordVisible ? (
                    <MdVisibilityOff size={24} className="text-gray-500" />
                  ) : (
                    <MdVisibility size={24} className="text-gray-500" />
                  )}
                </span>
              </div>
              <p style={{ color: "red", fontSize: "12px" }}>
                {errors.confirmPassword?.message}
              </p>
            </div>
            <div className="flex space-x-4 justify-center">
              <button
                onClick={() => handleUserType("student")}
                type="button"
                className={`w-full py-3 px-6 ${
                  userType == "student"
                    ? "bg-customBlue text-white"
                    : "bg-white text-gray-800"
                } rounded-lg shadow-md hover:bg-gray-500 transition ease-in-out duration-300`}
              >
                Student
              </button>
              <button
                onClick={() => handleUserType("teacher")}
                type="button"
                className={`w-full py-3 px-6 ${
                  userType == "teacher"
                    ? "bg-customBlue text-white"
                    : "bg-white text-gray-800"
                } rounded-lg shadow-md hover:bg-gray-500 transition ease-in-out duration-300`}
              >
                Teacher
              </button>
            </div>
            <div>{error && <p style={{ color: "red" }}>{error}</p>}</div>
            <div>
              <button
                type="submit"
                disabled={(!isDirty && !isValid) || isSubmitting}
                className="w-full py-3 px-6 bg-customBlue text-white rounded-lg shadow-md hover:bg-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition ease-in-out duration-300"
              >
                {loading ? <Spinner /> : "Sign Up"}
              </button>
            </div>
            <div className="text-center mt-6">
              <h3 className="text-sm">
                Already have an account?{" "}
                <Link href="/login">
                  <span className="text-indigo-600 hover:text-indigo-500 font-semibold">
                    Log In
                  </span>
                </Link>
              </h3>
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
                type="button"
                className="w-full py-3 px-6 bg-white text-black rounded-lg shadow-md border border-gray-300 hover:bg-gray-200 transition ease-in-out duration-300"
              >
                <div className="flex align-center justify-center">
                  <span className="mt-1 mr-2">
                    <FcGoogle />
                  </span>
                  <button onClick={handleSignUpWithGoogle}>
                    Sign In with Google
                  </button>
                </div>
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default SignUp;
